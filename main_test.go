package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockKafkaManage struct {
	mock.Mock
}

func (mock *MockKafkaManage) EnviarMensaje(mensaje string, topico string) error {
	args := mock.Called(mensaje, topico)
	return args.Error(0)
}
func (mock *MockKafkaManage) LeerMensaje(topico string) (mensaje string, err error) {
	args := mock.Called(mensaje, topico)
	return args.String(0), args.Error(1)
}

func MockProducirMensajeSuccess(mensaje, topico string) error {
	return nil
}

func TestEnviarMensaje(t *testing.T) {
	EMensaje = MockProducirMensajeSuccess
	err := EnviarMensaje("Mensaje desde main_test", "Tópico desde main_test")
	assert.Equal(t, nil, err)
}

func TestProducirMensaje(t *testing.T) {
	kafkaManage := new(MockKafkaManage)
	mensaje := MiMensaje{
		mensaje:     "Algun mensaje mock",
		topico:      "algun topico mock",
		kafkaManage: kafkaManage,
	}
	kafkaManage.On("EnviarMensaje", mensaje.mensaje, mensaje.topico).Return(errors.New("algun error"))
	err := mensaje.ProducirMensaje()
	kafkaManage.AssertExpectations(t)
	assert.Equal(t, errors.New("algun error"), err)
}
