package kafkaUtils

import (
	"errors"
	"fmt"
)

type KafkaManage struct {
}

func EnviarMensaje(mensaje string, topico string) error {
	if mensaje == "" {
		return errors.New("el mensaje está vacio")
	}
	if topico == "" {
		return errors.New("el topico está vacio")
	}
	fmt.Println("Nuevo mensaje", mensaje)
	return nil
}

func LeerMensaje(topico string) (mensaje string, err error) {
	if topico == "" {
		return "", errors.New("el topico está vacio")
	}

	return "Algun mensaje", nil
}

func (KafkaManage) EnviarMensaje(mensaje string, topico string) error {
	if mensaje == "" {
		return errors.New("el mensaje está vacio")
	}
	if topico == "" {
		return errors.New("el topico está vacio")
	}
	fmt.Println("Nuevo mensaje", mensaje)
	return nil
}

func (KafkaManage) LeerMensaje(topico string) (mensaje string, err error) {
	if topico == "" {
		return "", errors.New("el topico está vacio")
	}

	return "Algun mensaje", nil
}
