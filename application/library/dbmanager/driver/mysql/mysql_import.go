/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package mysql

import (
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/admpub/errors"
	"github.com/admpub/nging/application/handler"
	"github.com/admpub/nging/application/library/dbmanager/driver/mysql/utils"
	"github.com/admpub/nging/application/library/notice"
	"github.com/webx-top/echo"
)

func responseDropzone(err error, ctx echo.Context) error {
	if err != nil {
		user := handler.User(ctx)
		if user != nil {
			notice.OpenMessage(user.Username, `upload`)
			notice.Send(user.Username, notice.NewMessageWithValue(`upload`, ctx.T(`文件上传出错`), err.Error()))
		}
		return ctx.JSON(echo.H{`error`: err.Error()}, 500)
	}
	return ctx.String(`OK`)
}

func (m *mySQL) importing() error {
	return m.bgExecManage(utils.OpImport)
}

func (m *mySQL) Import() error {
	process := m.Queryx(`process`).Bool()
	if process {
		return m.importing()
	}
	var err error
	if m.IsPost() {
		if len(m.dbName) == 0 {
			m.fail(m.T(`请选择数据库`))
			return m.returnTo(m.GenURL(`listDb`))
		}
		async := m.Formx(`async`).Bool()
		var sqlFiles []string
		saveDir := TempDir(`import`)
		err = m.SaveUploadedFiles(`file`, func(fdr *multipart.FileHeader) (string, error) {
			extension := filepath.Ext(fdr.Filename)
			switch strings.ToLower(extension) {
			case `.sql`:
			case `.zip`:
			default:
				return ``, errors.New(`只能上传扩展名为“.sql”和“.zip”的文件`)
			}
			sqlFile := filepath.Join(saveDir, fdr.Filename)
			sqlFiles = append(sqlFiles, sqlFile)
			return sqlFile, nil
		})
		if err != nil {
			return responseDropzone(err, m.Context)
		}
		cfg := *m.DbAuth
		cfg.Db = m.dbName

		imports := utils.Exec{}
		bgExec := utils.NewGBExec(nil, echo.H{
			`database`: m.dbName,
			`sqlFiles`: sqlFiles,
			`async`:    async,
		})
		for _, sqlFile := range sqlFiles {
			fi, _ := os.Stat(sqlFile)
			bgExec.AddFileInfo(&utils.FileInfo{
				Start: time.Now(),
				Path:  sqlFile,
				Size:  fi.Size(),
			})
		}
		cacheKey := bgExec.Started.Format(`20060102150405`)
		imports.Add(utils.OpImport, cacheKey, bgExec)
		done := make(chan struct{})
		clientGone := m.Response().StdResponseWriter().(http.CloseNotifier).CloseNotify()
		go func() {
			for {
				select {
				case <-clientGone:
					bgExec.Cancel()()
					return
				case <-done:
					return
				}
			}
		}()
		err = utils.Import(bgExec.Context(), &cfg, TempDir(`import`), sqlFiles, async)
		done <- struct{}{}

		return responseDropzone(err, m.Context)
	}

	return m.Render(`db/mysql/import`, m.checkErr(err))
}
