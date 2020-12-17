package logic

import (
	"github.com/golang/protobuf/proto"
	"m1/aio"
	"m1/logger"
	pb "m1/protos"
)

type LoginLogic struct {
	Client *aio.Client
}

func init() {
	aio.LoginHandlerMap.Store(pb.Cmd_eMsgToLSFromGC_AskLogin, NewLoginLogic())
}

func NewLoginLogic() *LoginLogic {
	return &LoginLogic{
		Client: new(aio.Client),
	}
}

func (l *LoginLogic) HandlerCommon(msg *aio.InReqMsg) {
	l.Client = msg.Client
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
		Uid:  m.Uid,
		Data: "login success",
	}
	l.Client.PutOut(aio.NewOutResMsg(msg.Cmd, loginResult))

}

func (l *LoginLogic) HandlerError(cmd pb.Cmd, data string) {
	loginResult := &pb.LoginResult{
		Cmd:  cmd,
		Code: "-1",
		Uid:  "1001",
		Data: data,
	}
	l.Client.PutOut(aio.NewOutResMsg(cmd, loginResult))
}

func (l *LoginLogic) HandlerResponse(msg *aio.OutResMsg) {
	l.Client.PutOut(msg)
}
