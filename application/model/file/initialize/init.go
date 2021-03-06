package initialize

import (
	"io"

	"github.com/admpub/nging/application/library/common"
	modelFile "github.com/admpub/nging/application/model/file"
	"github.com/admpub/nging/application/registry/upload"
	uploadClient "github.com/webx-top/client/upload"
	"github.com/webx-top/db"
	"github.com/admpub/nging/application/dbschema"
)

func init() {
	common.OnUpdateOwnerFilePath = modelFile.OnUpdateOwnerFilePath
	common.OnRemoveOwnerFile = modelFile.OnRemoveOwnerFile
	upload.DefaultDBSaver = func(fileM *modelFile.File, result *uploadClient.Result, reader io.Reader) error {
		return fileM.Add(reader)
	}
	upload.DBSaverRegister(`user-avatar`, func(fileM *modelFile.File, result *uploadClient.Result, reader io.Reader) error {
		if fileM.OwnerId <= 0 {
			return fileM.Add(reader)
		}
		fileM.UsedTimes = 1
		m := &dbschema.File{}
		err := m.Get(nil, db.And(
			db.Cond{`table_name`: `user-avatar`},
			db.Cond{`table_id`: fileM.OwnerId},
		))
		if err != nil {
			if err != db.ErrNoMoreRows {
				return err
			}
			// 不存在
			return fileM.Add(reader)
		}
		// 已存在
		fileM.Id = m.Id
		fileM.Created = m.Created
		fileM.FillData(reader, true)
		return fileM.Edit(nil, db.Cond{`id`: m.Id})
	})
}
