package global

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"reflect"
)

/**
 * @ClassName rabbitmq
 * @Description TODO
 * @Author khr
 * @Date 2023/4/21 16:14
 * @Version 1.0
 */
var RabbitConn *amqp.Connection
var RabbitChannel *amqp.Channel

func Mqinit() {

	var err error
	RabbitConn, err = amqp.Dial(RabbitMQConfig.Url)
	if err != nil {
		log.Println("连接RabbitMQ失败")
		panic(err)
	}
	RabbitChannel, err = RabbitConn.Channel()
	if err != nil {
		log.Println("获取RabbitMQ channel失败")
		panic(err)
	}
	// 声明交换机
	err = RabbitChannel.ExchangeDeclare(
		"my_exchange",       // 交换机名称
		amqp.ExchangeDirect, // 交换机类型
		true,                // 是否持久化
		false,               // 是否自动删除
		false,               // 是否内部使用
		false,               // 是否等待服务器响应
		nil,                 // 其他属性
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}

	fmt.Println("rabbitmq初始化连接成功")
}

// // 0表示channel未关闭，1表示channel已关闭
func CheckRabbitClosed(ch amqp.Channel) int64 {
	d := reflect.ValueOf(ch)
	i := d.FieldByName("closed").Int()
	return i
}

func Producer(message string, producerName string) {
	// 声明队列，没有则创建
	// 队列名称、是否持久化、所有消费者与队列断开时是否自动删除队列、是否独享(不同连接的channel能否使用该队列)
	declare, err := RabbitChannel.QueueDeclare(producerName, true, false, false, false, nil)
	if err != nil {
		log.Printf("声明队列 %v 失败, error: %v", producerName, err)
		panic(err)
	}
	// 将队列绑定到交换机
	err = RabbitChannel.QueueBind(
		declare.Name,  // 队列名称
		"",            // 绑定键
		"my_exchange", // 交换机名称
		false,         // 是否等待服务器响应
		nil,           // 其他属性
	)
	if err != nil {
		log.Fatalf("Failed to bind the queue to the exchange: %v", err)
	}

	//request := model.Request{}
	marshal, _ := json.Marshal(message)
	// exchange、routing key、mandatory、immediate
	err = RabbitChannel.Publish("my_exchange", declare.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(marshal),
	})
	if err != nil {
		fmt.Println("生产者发送消息失败, error: %v", err)
	} else {
		fmt.Println("生产者发送消息成功")
	}
}
func Consumer(consumerName string) string {
	fmt.Println("传输的姓名", consumerName)
	// 声明队列，没有则创建
	// 队列名称、是否持久化、所有消费者与队列断开时是否自动删除队列、是否独享(不同连接的channel能否使用该队列)
	_, err := RabbitChannel.QueueDeclare("", true, false, false, false, nil)
	if err != nil {
		log.Printf("声明队列 %v 失败, error: %v", Consumer, err)
		panic(err)
	}

	err = RabbitChannel.Qos(
		1,     // prefetch count 服务器将在收到确认之前将那么多消息传递给消费者。
		0,     // prefetch size  服务器将尝试在收到消费者的确认之前至少将那么多字节的交付保持刷新到网络
		false, // 当 global 为 true 时，这些 Qos 设置适用于同一连接上所有通道上的所有现有和未来消费者。当为 false 时，Channel.Qos 设置将应用于此频道上的所有现有和未来消费者
	)
	if err != nil {
		log.Printf("rabbitmq设置Qos失败, error: %v", err)
	}

	// 队列名称、consumer、auto-ack、是否独享
	// deliveries是一个管道，有消息到队列，就会消费，消费者的消息只需要从deliveries这个管道获取
	deliveries, err := RabbitChannel.Consume(consumerName, "", false, false, false, false, nil)
	if err != nil {
		log.Printf("从队列 %v 获取数据失败, error: %v", Consumer, err)
	} else {
		log.Println("从消费队列获取任务成功")
	}

	// 阻塞住
	for {
		select {
		case message := <-deliveries:
			//closed := CheckRabbitClosed(*RabbitChannel)
			//if closed == 1 { // channel 已关闭，重连一下
			//	init()
			//	err = RabbitChannel.Qos(1, 0, false)
			//	if err != nil {
			//		log.Printf("rabbitmq重连后设置Qos失败, error: %v", err)
			//	}
			//} else {
			msgData := string(message.Body)
			//request := resmsg
			//err := json.Unmarshal([]byte(msgData), &request)
			//if err != nil {
			//	log.Printf("解析rabbitmq数据 %v 失败, error: %v", msgData, err)
			//} else {
			// TODO...
			// 处理逻辑
			return msgData
			// 处理完毕手动ACK
			message.Ack(true)
			//}
			//}
		}
	}
}
