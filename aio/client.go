package aio

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/google/uuid"
	comm "m1/common"
	"m1/logger"
	"net"
)

type Client struct {
	//连接信息
	Key     string //客户端连接的唯一标志
	Conn    net.Conn
	scanner *bufio.Scanner
	writer  *bufio.Writer
	Input   chan InReqMsg  //输入消息
	OutPut  chan OutResMsg //输出消息
	Quit    chan *Client   //推出消息
	//登录信息
}

func NewClient(key uuid.UUID, c net.Conn) *Client {
	client := &Client{
		Key:     key.String(),
		Conn:    c,
		scanner: bufio.NewScanner(c),
		writer:  bufio.NewWriter(c),
		Input:   make(chan InReqMsg),
		OutPut:  make(chan OutResMsg),
		Quit:    make(chan *Client),
	}
	client.listen()
	return client
}
func (c *Client) listen() {
	go c.read()
	go c.write()
}

func (c *Client) Quiting() {
	c.Quit <- c
}

func (c *Client) Close() {
	c.Conn.Close()
}

func (c *Client) PutOut(resp *OutResMsg) {
	c.OutPut <- *resp
}

func (c *Client) read() {
	for {
		c.scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			if !atEOF && len(data) > comm.HeadLen { //4字节数据包长度  4字节指令
				length := int32(0)
				binary.Read(bytes.NewReader(data[0:4]), binary.BigEndian, &length)
				if length <= 0 {
					return 0, nil, fmt.Errorf("length is 0")
				}
				logger.Logger.Info("len_data %d; length: %d\n", len(data), length)
				if int(length)+comm.HeadLen <= len(data) {
					return int(length) + comm.HeadLen, data[:int(length)+comm.HeadLen], nil
				}
			}
			return 1, nil, fmt.Errorf("data's length lt 8")
		})

		for c.scanner.Scan() {
			req, err := DecodeInReqMsg(c.scanner.Bytes())
			if err != nil {
				logger.Logger.Error("cmd err!!")
			} else {
				req.Client = c
				c.Input <- *req
			}
		}

		if err := c.scanner.Err(); err != nil {
			logger.Logger.Error("无数据包！")
			c.Quiting()
			return
		}
	}
}

func (c *Client) write() {
	for resp := range c.OutPut {
		if _, err := c.writer.Write(resp.Decode()); err != nil {
			logger.Logger.Errorf("write error %s\n", err.Error())
			c.Quiting()
			return
		}
		if err := c.writer.Flush(); err != nil {
			logger.Logger.Errorf("Write error: %s\n", err)
			c.Quiting()
			return
		}
	}

}
