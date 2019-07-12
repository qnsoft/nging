// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type User struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*User
	namer   func(string) string
	connID  int
	
	Id        	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Username  	string  	`db:"username" bson:"username" comment:"用户名" json:"username" xml:"username"`
	Email     	string  	`db:"email" bson:"email" comment:"邮箱" json:"email" xml:"email"`
	Mobile    	string  	`db:"mobile" bson:"mobile" comment:"手机号" json:"mobile" xml:"mobile"`
	Password  	string  	`db:"password" bson:"password" comment:"密码" json:"password" xml:"password"`
	Salt      	string  	`db:"salt" bson:"salt" comment:"盐值" json:"salt" xml:"salt"`
	SafePwd   	string  	`db:"safe_pwd" bson:"safe_pwd" comment:"安全密码" json:"safe_pwd" xml:"safe_pwd"`
	Avatar    	string  	`db:"avatar" bson:"avatar" comment:"头像" json:"avatar" xml:"avatar"`
	Gender    	string  	`db:"gender" bson:"gender" comment:"性别(male-男;female-女;secret-保密)" json:"gender" xml:"gender"`
	LastLogin 	uint    	`db:"last_login" bson:"last_login" comment:"最后登录时间" json:"last_login" xml:"last_login"`
	LastIp    	string  	`db:"last_ip" bson:"last_ip" comment:"最后登录IP" json:"last_ip" xml:"last_ip"`
	Disabled  	string  	`db:"disabled" bson:"disabled" comment:"状态" json:"disabled" xml:"disabled"`
	Online    	string  	`db:"online" bson:"online" comment:"是否在线" json:"online" xml:"online"`
	RoleIds   	string  	`db:"role_ids" bson:"role_ids" comment:"角色ID(多个用“,”分隔开)" json:"role_ids" xml:"role_ids"`
	Created   	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated   	uint    	`db:"updated" bson:"updated" comment:"更新时间" json:"updated" xml:"updated"`
}

func (this *User) Trans() *factory.Transaction {
	return this.trans
}

func (this *User) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *User) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *User) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *User) Objects() []*User {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *User) NewObjects() *[]*User {
	this.objects = []*User{}
	return &this.objects
}

func (this *User) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *User) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *User) Name_() string {
	if this.namer != nil {
		return this.namer("user")
	}
	return factory.TableNamerGet("user")(this)
}

func (this *User) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *User) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *User) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *User) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *User) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *User) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Gender) == 0 { this.Gender = "secret" }
	if len(this.Online) == 0 { this.Online = "N" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
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

func (this *User) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	this.Updated = uint(time.Now().Unix())
	if len(this.Gender) == 0 { this.Gender = "secret" }
	if len(this.Online) == 0 { this.Online = "N" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *User) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *User) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *User) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["gender"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["gender"] = "secret" } }
	if val, ok := kvset["online"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["online"] = "N" } }
	if val, ok := kvset["disabled"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["disabled"] = "N" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *User) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		this.Updated = uint(time.Now().Unix())
	if len(this.Gender) == 0 { this.Gender = "secret" }
	if len(this.Online) == 0 { this.Online = "N" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Gender) == 0 { this.Gender = "secret" }
	if len(this.Online) == 0 { this.Online = "N" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
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

func (this *User) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *User) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *User) Reset() *User {
	this.Id = 0
	this.Username = ``
	this.Email = ``
	this.Mobile = ``
	this.Password = ``
	this.Salt = ``
	this.SafePwd = ``
	this.Avatar = ``
	this.Gender = ``
	this.LastLogin = 0
	this.LastIp = ``
	this.Disabled = ``
	this.Online = ``
	this.RoleIds = ``
	this.Created = 0
	this.Updated = 0
	return this
}

func (this *User) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Username"] = this.Username
	r["Email"] = this.Email
	r["Mobile"] = this.Mobile
	r["Password"] = this.Password
	r["Salt"] = this.Salt
	r["SafePwd"] = this.SafePwd
	r["Avatar"] = this.Avatar
	r["Gender"] = this.Gender
	r["LastLogin"] = this.LastLogin
	r["LastIp"] = this.LastIp
	r["Disabled"] = this.Disabled
	r["Online"] = this.Online
	r["RoleIds"] = this.RoleIds
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	return r
}

func (this *User) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["username"] = this.Username
	r["email"] = this.Email
	r["mobile"] = this.Mobile
	r["password"] = this.Password
	r["salt"] = this.Salt
	r["safe_pwd"] = this.SafePwd
	r["avatar"] = this.Avatar
	r["gender"] = this.Gender
	r["last_login"] = this.LastLogin
	r["last_ip"] = this.LastIp
	r["disabled"] = this.Disabled
	r["online"] = this.Online
	r["role_ids"] = this.RoleIds
	r["created"] = this.Created
	r["updated"] = this.Updated
	return r
}

func (this *User) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate("user", kvset)
}

func (this *User) Validate(field string, value interface{}) error {
	return factory.Validate("user", field, value)
}

