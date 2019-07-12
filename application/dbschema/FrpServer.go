// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type FrpServer struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*FrpServer
	namer   func(string) string
	connID  int
	
	Id                  	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Name                	string  	`db:"name" bson:"name" comment:"名称" json:"name" xml:"name"`
	Disabled            	string  	`db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	TcpMux              	string  	`db:"tcp_mux" bson:"tcp_mux" comment:"" json:"tcp_mux" xml:"tcp_mux"`
	Addr                	string  	`db:"addr" bson:"addr" comment:"" json:"addr" xml:"addr"`
	Port                	uint    	`db:"port" bson:"port" comment:"" json:"port" xml:"port"`
	UdpPort             	uint    	`db:"udp_port" bson:"udp_port" comment:"" json:"udp_port" xml:"udp_port"`
	KcpPort             	uint    	`db:"kcp_port" bson:"kcp_port" comment:"" json:"kcp_port" xml:"kcp_port"`
	ProxyAddr           	string  	`db:"proxy_addr" bson:"proxy_addr" comment:"" json:"proxy_addr" xml:"proxy_addr"`
	VhostHttpPort       	uint    	`db:"vhost_http_port" bson:"vhost_http_port" comment:"" json:"vhost_http_port" xml:"vhost_http_port"`
	VhostHttpTimeout    	uint64  	`db:"vhost_http_timeout" bson:"vhost_http_timeout" comment:"" json:"vhost_http_timeout" xml:"vhost_http_timeout"`
	VhostHttpsPort      	uint    	`db:"vhost_https_port" bson:"vhost_https_port" comment:"" json:"vhost_https_port" xml:"vhost_https_port"`
	LogFile             	string  	`db:"log_file" bson:"log_file" comment:"" json:"log_file" xml:"log_file"`
	LogWay              	string  	`db:"log_way" bson:"log_way" comment:"console or file" json:"log_way" xml:"log_way"`
	LogLevel            	string  	`db:"log_level" bson:"log_level" comment:"" json:"log_level" xml:"log_level"`
	LogMaxDays          	uint    	`db:"log_max_days" bson:"log_max_days" comment:"" json:"log_max_days" xml:"log_max_days"`
	Token               	string  	`db:"token" bson:"token" comment:"" json:"token" xml:"token"`
	AuthTimeout         	uint64  	`db:"auth_timeout" bson:"auth_timeout" comment:"" json:"auth_timeout" xml:"auth_timeout"`
	SubdomainHost       	string  	`db:"subdomain_host" bson:"subdomain_host" comment:"" json:"subdomain_host" xml:"subdomain_host"`
	MaxPortsPerClient   	int64   	`db:"max_ports_per_client" bson:"max_ports_per_client" comment:"" json:"max_ports_per_client" xml:"max_ports_per_client"`
	MaxPoolCount        	uint    	`db:"max_pool_count" bson:"max_pool_count" comment:"" json:"max_pool_count" xml:"max_pool_count"`
	HeartBeatTimeout    	uint    	`db:"heart_beat_timeout" bson:"heart_beat_timeout" comment:"" json:"heart_beat_timeout" xml:"heart_beat_timeout"`
	UserConnTimeout     	uint    	`db:"user_conn_timeout" bson:"user_conn_timeout" comment:"" json:"user_conn_timeout" xml:"user_conn_timeout"`
	DashboardAddr       	string  	`db:"dashboard_addr" bson:"dashboard_addr" comment:"" json:"dashboard_addr" xml:"dashboard_addr"`
	DashboardPort       	uint    	`db:"dashboard_port" bson:"dashboard_port" comment:"" json:"dashboard_port" xml:"dashboard_port"`
	DashboardUser       	string  	`db:"dashboard_user" bson:"dashboard_user" comment:"" json:"dashboard_user" xml:"dashboard_user"`
	DashboardPwd        	string  	`db:"dashboard_pwd" bson:"dashboard_pwd" comment:"" json:"dashboard_pwd" xml:"dashboard_pwd"`
	AllowPorts          	string  	`db:"allow_ports" bson:"allow_ports" comment:"" json:"allow_ports" xml:"allow_ports"`
	Extra               	string  	`db:"extra" bson:"extra" comment:"" json:"extra" xml:"extra"`
	Uid                 	uint    	`db:"uid" bson:"uid" comment:"" json:"uid" xml:"uid"`
	GroupId             	uint    	`db:"group_id" bson:"group_id" comment:"" json:"group_id" xml:"group_id"`
	Created             	uint    	`db:"created" bson:"created" comment:"" json:"created" xml:"created"`
	Updated             	uint    	`db:"updated" bson:"updated" comment:"" json:"updated" xml:"updated"`
}

func (this *FrpServer) Trans() *factory.Transaction {
	return this.trans
}

func (this *FrpServer) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *FrpServer) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *FrpServer) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *FrpServer) Objects() []*FrpServer {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *FrpServer) NewObjects() *[]*FrpServer {
	this.objects = []*FrpServer{}
	return &this.objects
}

func (this *FrpServer) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *FrpServer) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *FrpServer) Name_() string {
	if this.namer != nil {
		return this.namer("frp_server")
	}
	return factory.TableNamerGet("frp_server")(this)
}

func (this *FrpServer) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *FrpServer) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *FrpServer) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *FrpServer) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *FrpServer) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *FrpServer) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.TcpMux) == 0 { this.TcpMux = "Y" }
	if len(this.DashboardAddr) == 0 { this.DashboardAddr = "0.0.0.0" }
	if len(this.DashboardPwd) == 0 { this.DashboardPwd = "admin" }
	if len(this.ProxyAddr) == 0 { this.ProxyAddr = "0.0.0.0" }
	if len(this.LogFile) == 0 { this.LogFile = "console" }
	if len(this.LogLevel) == 0 { this.LogLevel = "info" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Addr) == 0 { this.Addr = "0.0.0.0" }
	if len(this.LogWay) == 0 { this.LogWay = "console" }
	if len(this.DashboardUser) == 0 { this.DashboardUser = "admin" }
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

func (this *FrpServer) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	this.Updated = uint(time.Now().Unix())
	if len(this.TcpMux) == 0 { this.TcpMux = "Y" }
	if len(this.DashboardAddr) == 0 { this.DashboardAddr = "0.0.0.0" }
	if len(this.DashboardPwd) == 0 { this.DashboardPwd = "admin" }
	if len(this.ProxyAddr) == 0 { this.ProxyAddr = "0.0.0.0" }
	if len(this.LogFile) == 0 { this.LogFile = "console" }
	if len(this.LogLevel) == 0 { this.LogLevel = "info" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Addr) == 0 { this.Addr = "0.0.0.0" }
	if len(this.LogWay) == 0 { this.LogWay = "console" }
	if len(this.DashboardUser) == 0 { this.DashboardUser = "admin" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *FrpServer) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *FrpServer) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *FrpServer) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["tcp_mux"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["tcp_mux"] = "Y" } }
	if val, ok := kvset["dashboard_addr"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["dashboard_addr"] = "0.0.0.0" } }
	if val, ok := kvset["dashboard_pwd"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["dashboard_pwd"] = "admin" } }
	if val, ok := kvset["proxy_addr"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["proxy_addr"] = "0.0.0.0" } }
	if val, ok := kvset["log_file"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["log_file"] = "console" } }
	if val, ok := kvset["log_level"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["log_level"] = "info" } }
	if val, ok := kvset["disabled"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["disabled"] = "N" } }
	if val, ok := kvset["addr"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["addr"] = "0.0.0.0" } }
	if val, ok := kvset["log_way"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["log_way"] = "console" } }
	if val, ok := kvset["dashboard_user"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["dashboard_user"] = "admin" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *FrpServer) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		this.Updated = uint(time.Now().Unix())
	if len(this.TcpMux) == 0 { this.TcpMux = "Y" }
	if len(this.DashboardAddr) == 0 { this.DashboardAddr = "0.0.0.0" }
	if len(this.DashboardPwd) == 0 { this.DashboardPwd = "admin" }
	if len(this.ProxyAddr) == 0 { this.ProxyAddr = "0.0.0.0" }
	if len(this.LogFile) == 0 { this.LogFile = "console" }
	if len(this.LogLevel) == 0 { this.LogLevel = "info" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Addr) == 0 { this.Addr = "0.0.0.0" }
	if len(this.LogWay) == 0 { this.LogWay = "console" }
	if len(this.DashboardUser) == 0 { this.DashboardUser = "admin" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.TcpMux) == 0 { this.TcpMux = "Y" }
	if len(this.DashboardAddr) == 0 { this.DashboardAddr = "0.0.0.0" }
	if len(this.DashboardPwd) == 0 { this.DashboardPwd = "admin" }
	if len(this.ProxyAddr) == 0 { this.ProxyAddr = "0.0.0.0" }
	if len(this.LogFile) == 0 { this.LogFile = "console" }
	if len(this.LogLevel) == 0 { this.LogLevel = "info" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Addr) == 0 { this.Addr = "0.0.0.0" }
	if len(this.LogWay) == 0 { this.LogWay = "console" }
	if len(this.DashboardUser) == 0 { this.DashboardUser = "admin" }
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

func (this *FrpServer) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *FrpServer) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *FrpServer) Reset() *FrpServer {
	this.Id = 0
	this.Name = ``
	this.Disabled = ``
	this.TcpMux = ``
	this.Addr = ``
	this.Port = 0
	this.UdpPort = 0
	this.KcpPort = 0
	this.ProxyAddr = ``
	this.VhostHttpPort = 0
	this.VhostHttpTimeout = 0
	this.VhostHttpsPort = 0
	this.LogFile = ``
	this.LogWay = ``
	this.LogLevel = ``
	this.LogMaxDays = 0
	this.Token = ``
	this.AuthTimeout = 0
	this.SubdomainHost = ``
	this.MaxPortsPerClient = 0
	this.MaxPoolCount = 0
	this.HeartBeatTimeout = 0
	this.UserConnTimeout = 0
	this.DashboardAddr = ``
	this.DashboardPort = 0
	this.DashboardUser = ``
	this.DashboardPwd = ``
	this.AllowPorts = ``
	this.Extra = ``
	this.Uid = 0
	this.GroupId = 0
	this.Created = 0
	this.Updated = 0
	return this
}

func (this *FrpServer) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Name"] = this.Name
	r["Disabled"] = this.Disabled
	r["TcpMux"] = this.TcpMux
	r["Addr"] = this.Addr
	r["Port"] = this.Port
	r["UdpPort"] = this.UdpPort
	r["KcpPort"] = this.KcpPort
	r["ProxyAddr"] = this.ProxyAddr
	r["VhostHttpPort"] = this.VhostHttpPort
	r["VhostHttpTimeout"] = this.VhostHttpTimeout
	r["VhostHttpsPort"] = this.VhostHttpsPort
	r["LogFile"] = this.LogFile
	r["LogWay"] = this.LogWay
	r["LogLevel"] = this.LogLevel
	r["LogMaxDays"] = this.LogMaxDays
	r["Token"] = this.Token
	r["AuthTimeout"] = this.AuthTimeout
	r["SubdomainHost"] = this.SubdomainHost
	r["MaxPortsPerClient"] = this.MaxPortsPerClient
	r["MaxPoolCount"] = this.MaxPoolCount
	r["HeartBeatTimeout"] = this.HeartBeatTimeout
	r["UserConnTimeout"] = this.UserConnTimeout
	r["DashboardAddr"] = this.DashboardAddr
	r["DashboardPort"] = this.DashboardPort
	r["DashboardUser"] = this.DashboardUser
	r["DashboardPwd"] = this.DashboardPwd
	r["AllowPorts"] = this.AllowPorts
	r["Extra"] = this.Extra
	r["Uid"] = this.Uid
	r["GroupId"] = this.GroupId
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	return r
}

func (this *FrpServer) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["name"] = this.Name
	r["disabled"] = this.Disabled
	r["tcp_mux"] = this.TcpMux
	r["addr"] = this.Addr
	r["port"] = this.Port
	r["udp_port"] = this.UdpPort
	r["kcp_port"] = this.KcpPort
	r["proxy_addr"] = this.ProxyAddr
	r["vhost_http_port"] = this.VhostHttpPort
	r["vhost_http_timeout"] = this.VhostHttpTimeout
	r["vhost_https_port"] = this.VhostHttpsPort
	r["log_file"] = this.LogFile
	r["log_way"] = this.LogWay
	r["log_level"] = this.LogLevel
	r["log_max_days"] = this.LogMaxDays
	r["token"] = this.Token
	r["auth_timeout"] = this.AuthTimeout
	r["subdomain_host"] = this.SubdomainHost
	r["max_ports_per_client"] = this.MaxPortsPerClient
	r["max_pool_count"] = this.MaxPoolCount
	r["heart_beat_timeout"] = this.HeartBeatTimeout
	r["user_conn_timeout"] = this.UserConnTimeout
	r["dashboard_addr"] = this.DashboardAddr
	r["dashboard_port"] = this.DashboardPort
	r["dashboard_user"] = this.DashboardUser
	r["dashboard_pwd"] = this.DashboardPwd
	r["allow_ports"] = this.AllowPorts
	r["extra"] = this.Extra
	r["uid"] = this.Uid
	r["group_id"] = this.GroupId
	r["created"] = this.Created
	r["updated"] = this.Updated
	return r
}

func (this *FrpServer) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate("frp_server", kvset)
}

func (this *FrpServer) Validate(field string, value interface{}) error {
	return factory.Validate("frp_server", field, value)
}

