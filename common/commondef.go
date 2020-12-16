package common

import pb "m1/protos"

type SNetHead struct {
	Size int64
	Cmd  pb.Cmd
}

const HeadLen = 8
