package logic

import (
	"github.com/golang/protobuf/proto"
	"m1/aio"
	"m1/logger"
	pb "m1/protos"
)

type SendMsg struct {
	Client *aio.Client
}

func init() {
	aio.LoginHandlerMap.Store(pb.Cmd_eMsgToLSFromGC_SendMsg, &SendMsg{Client: nil})
}

func (l *SendMsg) HandlerCommon(msg *aio.InReqMsg) {
	l.Client = msg.Client

	m := &pb.SendMsg{}
	if err := proto.Unmarshal(msg.Data, m); err != nil {
		logger.Logger.Errorf("parse error %s", err.Error())
		l.HandlerError(msg.Cmd, "login  error")
		return
	}
	l.Client.PutOut(aio.NewOutResMsg(m.Cmd, m))

}

func (l *SendMsg) HandlerError(cmd pb.Cmd, data string) {
	loginResult := &pb.LoginResult{
		Cmd:  cmd,
		Code: "-1",
		Uid:  "1001",
		Data: data,
	}
	l.Client.PutOut(aio.NewOutResMsg(cmd, loginResult))
}

func (l *SendMsg) HandlerResponse(msg *aio.OutResMsg) {
	l.Client.PutOut(msg)
}
