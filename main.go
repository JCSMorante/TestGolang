package main

import (
	"test/interfaces"
	"test/kafkaUtils"
)

var (
	EMensaje = kafkaUtils.EnviarMensaje
	// print           = fmt.Println
)

type MiMensaje struct {
	mensaje, topico string
	kafkaManage     interfaces.IkafkaManage
}

func main() {
	//1. Mock cuando no es posible inyectar el objeto
	EnviarMensaje("Mensaje desde main", "Tópico desde main")
	//2. Mock cuando podemos usar un objeto inyectado
	mensaje := MiMensaje{
		mensaje:     "Un mensaje inyectado",
		topico:      "Un topico inyectado",
		kafkaManage: new(kafkaUtils.KafkaManage),
	}
	mensaje.ProducirMensaje()
}

func EnviarMensaje(mensaje, topico string) error {
	err := EMensaje(mensaje, topico)
	if err != nil {
		return err
	}
	return nil
}

func (miMensaje *MiMensaje) ProducirMensaje() error {
	err := miMensaje.kafkaManage.EnviarMensaje(miMensaje.mensaje, miMensaje.topico)
	if err != nil {
		return err
	}
	return nil
}

//Ejemplo de hacer un mock con variables
// func EnviarMensaje2(mensaje, topico string) error {
// 	_, err := print(mensaje, topico)
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }
