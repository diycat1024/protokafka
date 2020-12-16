package logic

import (
	"github.com/golang/protobuf/proto"
	"m1/aio"
	"m1/logger"
	pb "m1/protos"
	"time"
)

type Login struct {
	Id      string    `json:"id"`       // id
	UserId  string    `json:"user_id"`  // 用户ID
	Token   string    `json:"token"`    // 用户TOKEN
	LoginAt time.Time `json:"login_at"` // 登录日期
	LoginIp string    `json:"login_ip"` // 登录IP
	Plat    string    `json:"platform"` // 登录平台
}

type LoginLogic struct {
	UserLoginInfo *Login
	Client        *aio.Client
}

func init() {
	l := NewLoginLogic()
	aio.LoginHandlerMap.Store(pb.Cmd_eMsgToLSFromGC_AskLogin, l)
}

func NewLoginLogic() LoginLogic {
	return LoginLogic{
		UserLoginInfo: new(Login),
		Client:        new(aio.Client),
	}
}

func (l *LoginLogic) HandlerCommon(msg *aio.InReqMsg) {
	m := &pb.AskLogin{}
	if err := proto.Unmarshal(msg.Data, m); err != nil {
		logger.Logger.Errorf("parse error %s", err.Error())
		l.HandlerError(msg.Cmd, "login  error")
		return
	}

	//todo 登录成功
	loginResult := &pb.LoginResult{
		Cmd:  msg.Cmd,
		Code: "0",
		Uid:  "1001",
		Data: "login success",
	}
	m1 := proto.Message(loginResult)
	out := &aio.OutResMsg{
		Code: "0",
		Cmd:  msg.Cmd,
		Data: &m1,
	}
	l.HandlerResponse(out)

}

func (l *LoginLogic) HandlerError(cmd pb.Cmd, data string) {
	loginResult := &pb.LoginResult{
		Cmd:  cmd,
		Code: "-1",
		Uid:  "1001",
		Data: data,
	}
	m1 := proto.Message(loginResult)
	out := &aio.OutResMsg{
		Code: "0",
		Cmd:  cmd,
		Data: &m1,
	}
	l.Client.PutOut(out)
}

func (l *LoginLogic) HandlerResponse(msg *aio.OutResMsg) {
	l.Client.PutOut(msg)
}
