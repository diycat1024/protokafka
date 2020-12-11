package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

var (
	producer sarama.SyncProducer
	brokers = []string{"127.0.0.1:9092"}
	topic   = "mykafka"
)

func init() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	brokers := brokers
	var err error
	producer, err = sarama.NewSyncProducer(brokers, config)
	if err != nil {
		fmt.Printf("init producer failed -> %v \n", err)
		panic(err)
	}
	fmt.Println("producer init success")
}

func produceMsg(msg string) {
	msgX := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}
	fmt.Printf("SendMsg -> %v\n", dumpString(msgX))

	partition, offset, err := producer.SendMessage(msgX)
	if err != nil {
		fmt.Printf("send msg error:%s \n", err)
	}
	fmt.Printf("msg send success, message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
}

func startProduce() {
	tick := time.Tick(1 * time.Millisecond)
	for {
		select {
		case <-tick:
			t := time.Now().Unix() * 1000
			msg := fmt.Sprintf("{\"timestamp\":%d}", t)
			produceMsg(msg)
			produceMsg(msg)
		}
	}
}

//解析为json字符串
func dumpString(v interface{}) (str string) {

	bs, err := json.Marshal(v)
	b := bytes.Buffer{}
	if err != nil {
		b.WriteString("{err:\"json format error.")
		b.WriteString(err.Error())
		b.WriteString("\"}")
	} else {
		b.Write(bs)
	}
	str = b.String()
	return str
}

func test(v interface{})  {
	switch v.(type) {
	case bool:
		fmt.Printf("this is bool %v\n", v)
	case int:
		fmt.Printf("this is int %v\n", v)
	default:
		fmt.Printf("nothing \n")
	}
}

func main() {
	startProduce()
}