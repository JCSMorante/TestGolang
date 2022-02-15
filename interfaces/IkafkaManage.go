package interfaces

type IkafkaManage interface {
	EnviarMensaje(mensaje string, topico string) error
	LeerMensaje(topico string) (mensaje string, err error)
}
