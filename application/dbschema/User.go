// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo/param"
	
	"time"
)

// User 用户
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
	FileSize  	uint64  	`db:"file_size" bson:"file_size" comment:"上传文件总大小" json:"file_size" xml:"file_size"`
	FileNum   	uint64  	`db:"file_num" bson:"file_num" comment:"上传文件数量" json:"file_num" xml:"file_num"`
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

func (this *User) Short_() string {
	return "user"
}

func (this *User) Struct_() string {
	return "User"
}

func (this *User) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
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

func (this *User) GroupBy(keyField string, inputRows ...[]*User) map[string][]*User {
	var rows []*User
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*User{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*User{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *User) KeyBy(keyField string, inputRows ...[]*User) map[string]*User {
	var rows []*User
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*User{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *User) AsKV(keyField string, valueField string, inputRows ...[]*User) map[string]interface{} {
	var rows []*User
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]interface{}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
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
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Gender) == 0 { this.Gender = "secret" }
	if len(this.Online) == 0 { this.Online = "N" }
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
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Gender) == 0 { this.Gender = "secret" }
	if len(this.Online) == 0 { this.Online = "N" }
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
	
	if val, ok := kvset["disabled"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["disabled"] = "N" } }
	if val, ok := kvset["gender"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["gender"] = "secret" } }
	if val, ok := kvset["online"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["online"] = "N" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *User) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		this.Updated = uint(time.Now().Unix())
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Gender) == 0 { this.Gender = "secret" }
	if len(this.Online) == 0 { this.Online = "N" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Gender) == 0 { this.Gender = "secret" }
	if len(this.Online) == 0 { this.Online = "N" }
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
	this.FileSize = 0
	this.FileNum = 0
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
	r["FileSize"] = this.FileSize
	r["FileNum"] = this.FileNum
	return r
}

func (this *User) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
		case map[string]interface{}:
			for kk, vv := range k {
				this.Set(kk, vv)
			}
		default:
			var (
				kk string
				vv interface{}
			)
			if k, y := key.(string); y {
				kk = k
			} else {
				kk = fmt.Sprint(key)
			}
			if len(value) > 0 {
				vv = value[0]
			}
			switch kk {
				case "Id": this.Id = param.AsUint(vv)
				case "Username": this.Username = param.AsString(vv)
				case "Email": this.Email = param.AsString(vv)
				case "Mobile": this.Mobile = param.AsString(vv)
				case "Password": this.Password = param.AsString(vv)
				case "Salt": this.Salt = param.AsString(vv)
				case "SafePwd": this.SafePwd = param.AsString(vv)
				case "Avatar": this.Avatar = param.AsString(vv)
				case "Gender": this.Gender = param.AsString(vv)
				case "LastLogin": this.LastLogin = param.AsUint(vv)
				case "LastIp": this.LastIp = param.AsString(vv)
				case "Disabled": this.Disabled = param.AsString(vv)
				case "Online": this.Online = param.AsString(vv)
				case "RoleIds": this.RoleIds = param.AsString(vv)
				case "Created": this.Created = param.AsUint(vv)
				case "Updated": this.Updated = param.AsUint(vv)
				case "FileSize": this.FileSize = param.AsUint64(vv)
				case "FileNum": this.FileNum = param.AsUint64(vv)
			}
	}
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
	r["file_size"] = this.FileSize
	r["file_num"] = this.FileNum
	return r
}

func (this *User) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *User) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

