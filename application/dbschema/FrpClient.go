// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type FrpClient struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*FrpClient
	namer   func(string) string
	connID  int
	
	Id                	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Name              	string  	`db:"name" bson:"name" comment:"名称" json:"name" xml:"name"`
	Disabled          	string  	`db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	ServerAddr        	string  	`db:"server_addr" bson:"server_addr" comment:"" json:"server_addr" xml:"server_addr"`
	ServerPort        	uint    	`db:"server_port" bson:"server_port" comment:"" json:"server_port" xml:"server_port"`
	HttpProxy         	string  	`db:"http_proxy" bson:"http_proxy" comment:"" json:"http_proxy" xml:"http_proxy"`
	PoolCount         	uint    	`db:"pool_count" bson:"pool_count" comment:"" json:"pool_count" xml:"pool_count"`
	TcpMux            	string  	`db:"tcp_mux" bson:"tcp_mux" comment:"必须跟服务端一致" json:"tcp_mux" xml:"tcp_mux"`
	User              	string  	`db:"user" bson:"user" comment:"" json:"user" xml:"user"`
	DnsServer         	string  	`db:"dns_server" bson:"dns_server" comment:"" json:"dns_server" xml:"dns_server"`
	LoginFailExit     	string  	`db:"login_fail_exit" bson:"login_fail_exit" comment:"" json:"login_fail_exit" xml:"login_fail_exit"`
	Protocol          	string  	`db:"protocol" bson:"protocol" comment:"tcp or kcp" json:"protocol" xml:"protocol"`
	HeartbeatInterval 	int64   	`db:"heartbeat_interval" bson:"heartbeat_interval" comment:"" json:"heartbeat_interval" xml:"heartbeat_interval"`
	HeartbeatTimeout  	int64   	`db:"heartbeat_timeout" bson:"heartbeat_timeout" comment:"" json:"heartbeat_timeout" xml:"heartbeat_timeout"`
	LogFile           	string  	`db:"log_file" bson:"log_file" comment:"" json:"log_file" xml:"log_file"`
	LogWay            	string  	`db:"log_way" bson:"log_way" comment:"console or file" json:"log_way" xml:"log_way"`
	LogLevel          	string  	`db:"log_level" bson:"log_level" comment:"" json:"log_level" xml:"log_level"`
	LogMaxDays        	uint    	`db:"log_max_days" bson:"log_max_days" comment:"" json:"log_max_days" xml:"log_max_days"`
	Token             	string  	`db:"token" bson:"token" comment:"" json:"token" xml:"token"`
	AdminAddr         	string  	`db:"admin_addr" bson:"admin_addr" comment:"" json:"admin_addr" xml:"admin_addr"`
	AdminPort         	uint    	`db:"admin_port" bson:"admin_port" comment:"" json:"admin_port" xml:"admin_port"`
	AdminUser         	string  	`db:"admin_user" bson:"admin_user" comment:"" json:"admin_user" xml:"admin_user"`
	AdminPwd          	string  	`db:"admin_pwd" bson:"admin_pwd" comment:"" json:"admin_pwd" xml:"admin_pwd"`
	Start             	string  	`db:"start" bson:"start" comment:"要启动的代理节点名称，留空代表全部，多个用半角逗号隔开" json:"start" xml:"start"`
	Extra             	string  	`db:"extra" bson:"extra" comment:"" json:"extra" xml:"extra"`
	Uid               	uint    	`db:"uid" bson:"uid" comment:"" json:"uid" xml:"uid"`
	GroupId           	uint    	`db:"group_id" bson:"group_id" comment:"" json:"group_id" xml:"group_id"`
	Type              	string  	`db:"type" bson:"type" comment:"" json:"type" xml:"type"`
	Created           	uint    	`db:"created" bson:"created" comment:"" json:"created" xml:"created"`
	Updated           	uint    	`db:"updated" bson:"updated" comment:"" json:"updated" xml:"updated"`
}

func (this *FrpClient) Trans() *factory.Transaction {
	return this.trans
}

func (this *FrpClient) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *FrpClient) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *FrpClient) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *FrpClient) Objects() []*FrpClient {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *FrpClient) NewObjects() *[]*FrpClient {
	this.objects = []*FrpClient{}
	return &this.objects
}

func (this *FrpClient) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *FrpClient) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *FrpClient) Name_() string {
	if this.namer != nil {
		return this.namer("frp_client")
	}
	return factory.TableNamerGet("frp_client")(this)
}

func (this *FrpClient) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *FrpClient) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *FrpClient) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *FrpClient) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *FrpClient) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *FrpClient) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.LogFile) == 0 { this.LogFile = "console" }
	if len(this.LoginFailExit) == 0 { this.LoginFailExit = "Y" }
	if len(this.Protocol) == 0 { this.Protocol = "tcp" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.TcpMux) == 0 { this.TcpMux = "Y" }
	if len(this.LogWay) == 0 { this.LogWay = "console" }
	if len(this.ServerAddr) == 0 { this.ServerAddr = "0.0.0.0" }
	if len(this.LogLevel) == 0 { this.LogLevel = "info" }
	if len(this.Type) == 0 { this.Type = "web" }
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

func (this *FrpClient) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	this.Updated = uint(time.Now().Unix())
	if len(this.LogFile) == 0 { this.LogFile = "console" }
	if len(this.LoginFailExit) == 0 { this.LoginFailExit = "Y" }
	if len(this.Protocol) == 0 { this.Protocol = "tcp" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.TcpMux) == 0 { this.TcpMux = "Y" }
	if len(this.LogWay) == 0 { this.LogWay = "console" }
	if len(this.ServerAddr) == 0 { this.ServerAddr = "0.0.0.0" }
	if len(this.LogLevel) == 0 { this.LogLevel = "info" }
	if len(this.Type) == 0 { this.Type = "web" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *FrpClient) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *FrpClient) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *FrpClient) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["log_file"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["log_file"] = "console" } }
	if val, ok := kvset["login_fail_exit"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["login_fail_exit"] = "Y" } }
	if val, ok := kvset["protocol"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["protocol"] = "tcp" } }
	if val, ok := kvset["disabled"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["disabled"] = "N" } }
	if val, ok := kvset["tcp_mux"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["tcp_mux"] = "Y" } }
	if val, ok := kvset["log_way"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["log_way"] = "console" } }
	if val, ok := kvset["server_addr"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["server_addr"] = "0.0.0.0" } }
	if val, ok := kvset["log_level"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["log_level"] = "info" } }
	if val, ok := kvset["type"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["type"] = "web" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *FrpClient) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		this.Updated = uint(time.Now().Unix())
	if len(this.LogFile) == 0 { this.LogFile = "console" }
	if len(this.LoginFailExit) == 0 { this.LoginFailExit = "Y" }
	if len(this.Protocol) == 0 { this.Protocol = "tcp" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.TcpMux) == 0 { this.TcpMux = "Y" }
	if len(this.LogWay) == 0 { this.LogWay = "console" }
	if len(this.ServerAddr) == 0 { this.ServerAddr = "0.0.0.0" }
	if len(this.LogLevel) == 0 { this.LogLevel = "info" }
	if len(this.Type) == 0 { this.Type = "web" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.LogFile) == 0 { this.LogFile = "console" }
	if len(this.LoginFailExit) == 0 { this.LoginFailExit = "Y" }
	if len(this.Protocol) == 0 { this.Protocol = "tcp" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.TcpMux) == 0 { this.TcpMux = "Y" }
	if len(this.LogWay) == 0 { this.LogWay = "console" }
	if len(this.ServerAddr) == 0 { this.ServerAddr = "0.0.0.0" }
	if len(this.LogLevel) == 0 { this.LogLevel = "info" }
	if len(this.Type) == 0 { this.Type = "web" }
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

func (this *FrpClient) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *FrpClient) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *FrpClient) Reset() *FrpClient {
	this.Id = 0
	this.Name = ``
	this.Disabled = ``
	this.ServerAddr = ``
	this.ServerPort = 0
	this.HttpProxy = ``
	this.PoolCount = 0
	this.TcpMux = ``
	this.User = ``
	this.DnsServer = ``
	this.LoginFailExit = ``
	this.Protocol = ``
	this.HeartbeatInterval = 0
	this.HeartbeatTimeout = 0
	this.LogFile = ``
	this.LogWay = ``
	this.LogLevel = ``
	this.LogMaxDays = 0
	this.Token = ``
	this.AdminAddr = ``
	this.AdminPort = 0
	this.AdminUser = ``
	this.AdminPwd = ``
	this.Start = ``
	this.Extra = ``
	this.Uid = 0
	this.GroupId = 0
	this.Type = ``
	this.Created = 0
	this.Updated = 0
	return this
}

func (this *FrpClient) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Name"] = this.Name
	r["Disabled"] = this.Disabled
	r["ServerAddr"] = this.ServerAddr
	r["ServerPort"] = this.ServerPort
	r["HttpProxy"] = this.HttpProxy
	r["PoolCount"] = this.PoolCount
	r["TcpMux"] = this.TcpMux
	r["User"] = this.User
	r["DnsServer"] = this.DnsServer
	r["LoginFailExit"] = this.LoginFailExit
	r["Protocol"] = this.Protocol
	r["HeartbeatInterval"] = this.HeartbeatInterval
	r["HeartbeatTimeout"] = this.HeartbeatTimeout
	r["LogFile"] = this.LogFile
	r["LogWay"] = this.LogWay
	r["LogLevel"] = this.LogLevel
	r["LogMaxDays"] = this.LogMaxDays
	r["Token"] = this.Token
	r["AdminAddr"] = this.AdminAddr
	r["AdminPort"] = this.AdminPort
	r["AdminUser"] = this.AdminUser
	r["AdminPwd"] = this.AdminPwd
	r["Start"] = this.Start
	r["Extra"] = this.Extra
	r["Uid"] = this.Uid
	r["GroupId"] = this.GroupId
	r["Type"] = this.Type
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	return r
}

func (this *FrpClient) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["name"] = this.Name
	r["disabled"] = this.Disabled
	r["server_addr"] = this.ServerAddr
	r["server_port"] = this.ServerPort
	r["http_proxy"] = this.HttpProxy
	r["pool_count"] = this.PoolCount
	r["tcp_mux"] = this.TcpMux
	r["user"] = this.User
	r["dns_server"] = this.DnsServer
	r["login_fail_exit"] = this.LoginFailExit
	r["protocol"] = this.Protocol
	r["heartbeat_interval"] = this.HeartbeatInterval
	r["heartbeat_timeout"] = this.HeartbeatTimeout
	r["log_file"] = this.LogFile
	r["log_way"] = this.LogWay
	r["log_level"] = this.LogLevel
	r["log_max_days"] = this.LogMaxDays
	r["token"] = this.Token
	r["admin_addr"] = this.AdminAddr
	r["admin_port"] = this.AdminPort
	r["admin_user"] = this.AdminUser
	r["admin_pwd"] = this.AdminPwd
	r["start"] = this.Start
	r["extra"] = this.Extra
	r["uid"] = this.Uid
	r["group_id"] = this.GroupId
	r["type"] = this.Type
	r["created"] = this.Created
	r["updated"] = this.Updated
	return r
}

func (this *FrpClient) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate("frp_client", kvset)
}

func (this *FrpClient) Validate(field string, value interface{}) error {
	return factory.Validate("frp_client", field, value)
}

