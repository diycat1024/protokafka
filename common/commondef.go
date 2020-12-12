package common

import pb "m1/protos"

type SNetHead struct {
	Size int64
	Cmd  pb.MsgID
}

const HeadLen = 8
