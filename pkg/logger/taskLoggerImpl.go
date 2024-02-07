package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const qName = "logger"

func (t *LoggerImpl) Write(msg interface{}) {

	logMsg := &LoggerMsg{
		Log: fmt.Sprintf("%v", msg),
	}

	body, err := json.Marshal(logMsg)
	if err != nil {
		log.Printf("%s: %s\n", "Error encoding JSON:", err)
	}
	log.Println("body: ", string(body))

	err = t.amqpChannel.Publish("", qName, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})

	if err != nil {
		log.Fatalf("Error publishing message: %s\n", err)
		//return err
	}
	log.Printf("publishing message: %s\n", msg)
}

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

func (t *LoggerImpl) Init() error {

	time.Sleep(time.Second * 15)

	log.Printf("dialing %q\n", t.config.MQ.Host)

	conn, err := amqp.Dial(t.config.MQ.Host) // подключение к RabbitMQ
	if err != nil {
		log.Printf("%s: %s\n", "ошибка подключения AMQP:", err)
		return err
	}
	//defer conn.Close()

	t.amqpChannel, err = conn.Channel() //установка канала RabbitMQ
	if err != nil {
		log.Printf("%s: %s\n", "ошибка создания amqpChannel:", err)
		return err
	}
	//defer t.amqpChannel.Close()

	err = t.amqpChannel.ExchangeDeclare(
		qName,    // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Printf("%s: %s\n", "ошибка ExchangeDeclare:", err)
		return err
	}

	_, err = t.amqpChannel.QueueDeclare( // объявляет очередь для хранения сообщений и их доставки потребителям.
		qName, //  имя очереди
		true,  //   Сохранять ли
		false, // Удаляется ли оно автоматически
		false, // Это эксклюзив
		false, // Следует ли блокировать
		nil,   // Дополнительные атрибуты
	)

	if err != nil {
		log.Printf("%s: %s\n", "ошибка объявления очереди:", err)
		return err
	}

	return nil
}
