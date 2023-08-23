package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
	"log"
	"time"
)

var (
	// go run main.go -addr=localhost:9092
	addr      = flag.String("addr", "localhost:9092", "-addr=localhost:9092")
	topic     = "my-topic"
	partition = 0
	conn      *kafka.Conn
)

func init() {
	flag.Parse()

	connectKafka()
}

// connectKafka 连接到 kafka
func connectKafka() {
	var err error
	conn, err = kafka.DialLeader(context.Background(), "tcp", *addr, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	_ = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
}

func sendMsgToKafka() {
	_, err := conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	log.Println("send message to kafka successfully")
}

// sendMsgToKafkaWithNum 发送 num 条消息到 kafka
func sendMsgToKafkaWithNum(num int) {
	var messages []kafka.Message
	for i := 0; i < num; i++ {
		messages = append(messages, kafka.Message{
			Value: []byte(fmt.Sprintf("message-%d", i)),
		})
	}

	res, err := conn.WriteMessages(messages...)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	log.Printf("send %d message to kafka successfully\n", res)
}

// consumeSingleMsgFromKafka 消费单条消息
func consumeSingleMsgFromKafka() {
	message, err := conn.ReadMessage(1024)
	if err != nil {
		panic(err)
	}
	log.Println(message)
}

// consumeBatchMsgFromKafka 消费批量消息
func consumeBatchMsgFromKafka() {
	// 批量读取，所有消息需要达到这个阈值，才能读取的到
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}
}

// createTopic 创建 topic
func createTopic(topic string) {
	// to create topics when auto.create.topics.enable='true'
	conn, err := kafka.DialLeader(context.Background(), "tcp", *addr, topic, 0)
	defer func() {
		err := conn.Close()
		if err != nil {
			return
		}
	}()
	if err != nil {
		panic(err.Error())
	}
}

// consumeGroupExample 消费组例子
func consumeGroupExample() {
	go func() {
		createTopic("consumeGroupExample")

		w := kafka.Writer{
			Addr:     kafka.TCP(*addr),
			Topic:    "consumeGroupExample",
			Balancer: &kafka.LeastBytes{},
		}

		defer func() {
			if err := w.Close(); err != nil {
				log.Fatal("failed to close writer:", err)
			}
		}()

		for {
			err := w.WriteMessages(context.Background(),
				kafka.Message{
					Key:   []byte("Key-A"),
					Value: []byte("Hello World!"),
				},
				kafka.Message{
					Key:   []byte("Key-B"),
					Value: []byte("One!"),
				},
				kafka.Message{
					Key:   []byte("Key-C"),
					Value: []byte("Two!"),
				},
			)

			if err != nil {
				log.Fatal("failed to write messages:", err)
			}
		}
	}()

	go func() {
		// 等待生产者就绪
		time.Sleep(time.Second * 2)
		r := kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{*addr},
			GroupID:  "consume-group-id",
			Topic:    "consumeGroupExample",
			MinBytes: 10e3, // 10KB，最小的消息单元
			MaxBytes: 10e6, // 10MB
		})

		defer func() {
			err := r.Close()
			if err != nil {
				return
			}
		}()

		for {
			message, err := r.ReadMessage(context.Background())
			if err != nil {
				log.Fatal("failed to read messages:", err)
			}

			log.Printf("topic: %s key: %s value: %s\n", message.Topic, string(message.Key), string(message.Value))
		}
	}()

	time.Sleep(time.Minute)
}

func brokerExample() {
	// 返回当前 connection 所连接的 broker
	broker := conn.Broker()
	log.Println("current broker: ", broker)

	// 返回 kafka 元数据中的 broker（集群所有 broker）
	brokers, err := conn.Brokers()
	if err != nil {
		panic(err)
	}
	for _, b := range brokers {
		log.Println("Brokers：", b)
	}
}

func compressMessageExample() {
	// gzip
	codec := compress.Compression(1).Codec()
	if codec == nil {
		log.Fatalf("not found %d codec", 1)
	}
	// write message with specified codec
	_, err := conn.WriteCompressedMessages(codec, kafka.Message{
		Key:   []byte("name"),
		Value: []byte("ray"),
	})
	if err != nil {
		log.Fatal("failed to write message: ", err.Error())
	}

	message, err := conn.ReadMessage(20)
	if err != nil {
		log.Fatal("failed to read message from kafka: ", err.Error())
	}
	log.Printf("key: %s  \tvalue：%s\n", string(message.Key), string(message.Value))
}

func main() {
	defer func() {
		_ = conn.Close()
	}()

	//sendMsgToKafka()
	//consumeSingleMsgFromKafka()

	//sendMsgToKafkaWithNum(1000)
	//consumeBatchMsgFromKafka()

	//consumeGroupExample()

	//brokerExample()
	compressMessageExample()
}
