package logic

import (
	"github.com/golang/protobuf/proto"
	"m1/aio"
	"m1/logger"
	pb "m1/protos"
)

type SendGroupMsg struct {
	Client *aio.Client
}

func init() {
	aio.LoginHandlerMap.Store(pb.Cmd_eMsgToLSFromGC_SendGroupMsg, &SendGroupMsg{Client: nil})
}

func (l *SendGroupMsg) HandlerCommon(msg *aio.InReqMsg) {
	l.Client = msg.Client

	m := &pb.SendGroupMsg{}
	if err := proto.Unmarshal(msg.Data, m); err != nil {
		logger.Logger.Errorf("parse error %s", err.Error())
		l.HandlerError(msg.Cmd, "login  error")
		return
	}

	l.Client.PutOut(aio.NewOutResMsg(m.Cmd, m))

	msg.Server.Clients.Range(func(key, value interface{}) bool {
		client := value.(*aio.Client)
		logger.Logger.Debugf("接收对象：%s\n", client.Key)
		l.Client.PutOut(aio.NewOutResMsg(m.Cmd, m))
		return true
	})
}

func (l *SendGroupMsg) HandlerError(cmd pb.Cmd, data string) {
	loginResult := &pb.LoginResult{
		Cmd:  cmd,
		Code: "-1",
		Uid:  "1001",
		Data: data,
	}
	l.Client.PutOut(aio.NewOutResMsg(cmd, loginResult))
}

func (l *SendGroupMsg) HandlerResponse(msg *aio.OutResMsg) {
	l.Client.PutOut(msg)
}
