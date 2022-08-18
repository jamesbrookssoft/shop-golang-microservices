package consume

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func HandleConsumeCreateProduct(queue string, msg amqp.Delivery, err error) {
	if err != nil {
		log.Error(err)
	}
	log.Infof("Message received on '%s' queue: %s", queue, string(msg.Body))
}
