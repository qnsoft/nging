// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type DbSync struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*DbSync
	namer   func(string) string
	connID  int
	
	Id             	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	DsnSource      	string  	`db:"dsn_source" bson:"dsn_source" comment:"同步源" json:"dsn_source" xml:"dsn_source"`
	DsnDestination 	string  	`db:"dsn_destination" bson:"dsn_destination" comment:"目标数据库" json:"dsn_destination" xml:"dsn_destination"`
	Tables         	string  	`db:"tables" bson:"tables" comment:"要同步的表" json:"tables" xml:"tables"`
	SkipTables     	string  	`db:"skip_tables" bson:"skip_tables" comment:"要跳过的表" json:"skip_tables" xml:"skip_tables"`
	AlterIgnore    	string  	`db:"alter_ignore" bson:"alter_ignore" comment:"要忽略的列、索引、外键" json:"alter_ignore" xml:"alter_ignore"`
	Drop           	uint    	`db:"drop" bson:"drop" comment:"删除待同步数据库中多余的字段、索引、外键 " json:"drop" xml:"drop"`
	MailTo         	string  	`db:"mail_to" bson:"mail_to" comment:"发送邮件" json:"mail_to" xml:"mail_to"`
	Created        	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated        	int     	`db:"updated" bson:"updated" comment:"更新时间" json:"updated" xml:"updated"`
}

func (this *DbSync) Trans() *factory.Transaction {
	return this.trans
}

func (this *DbSync) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *DbSync) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *DbSync) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *DbSync) Objects() []*DbSync {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *DbSync) NewObjects() *[]*DbSync {
	this.objects = []*DbSync{}
	return &this.objects
}

func (this *DbSync) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *DbSync) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *DbSync) Name_() string {
	if this.namer != nil {
		return this.namer("db_sync")
	}
	return factory.TableNamerGet("db_sync")(this)
}

func (this *DbSync) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *DbSync) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *DbSync) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *DbSync) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *DbSync) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *DbSync) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return
}

func (this *DbSync) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	this.Updated = int(time.Now().Unix())
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *DbSync) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *DbSync) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *DbSync) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	kvset["updated"] = int(time.Now().Unix())
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *DbSync) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		this.Updated = int(time.Now().Unix())
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return 
}

func (this *DbSync) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *DbSync) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *DbSync) Reset() *DbSync {
	this.Id = 0
	this.DsnSource = ``
	this.DsnDestination = ``
	this.Tables = ``
	this.SkipTables = ``
	this.AlterIgnore = ``
	this.Drop = 0
	this.MailTo = ``
	this.Created = 0
	this.Updated = 0
	return this
}

func (this *DbSync) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["DsnSource"] = this.DsnSource
	r["DsnDestination"] = this.DsnDestination
	r["Tables"] = this.Tables
	r["SkipTables"] = this.SkipTables
	r["AlterIgnore"] = this.AlterIgnore
	r["Drop"] = this.Drop
	r["MailTo"] = this.MailTo
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	return r
}

func (this *DbSync) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["dsn_source"] = this.DsnSource
	r["dsn_destination"] = this.DsnDestination
	r["tables"] = this.Tables
	r["skip_tables"] = this.SkipTables
	r["alter_ignore"] = this.AlterIgnore
	r["drop"] = this.Drop
	r["mail_to"] = this.MailTo
	r["created"] = this.Created
	r["updated"] = this.Updated
	return r
}

func (this *DbSync) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate("db_sync", kvset)
}

func (this *DbSync) Validate(field string, value interface{}) error {
	return factory.Validate("db_sync", field, value)
}

