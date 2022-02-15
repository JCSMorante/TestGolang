package kafkaUtils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Test con libreria estandar
func TestProducirMensaje(t *testing.T) {
	err := EnviarMensaje("", "algún tópico")
	if err == nil {
		t.Errorf("Se esperaba: %v, se obtuvo %v", errors.New("el mensaje está vacio"), err)
	}
	err = EnviarMensaje("algun mensaje", "")
	if err == nil {
		t.Errorf("Se esperaba: %v, se obtuvo %v", errors.New("el topico está vacio"), err)
	}
	err = EnviarMensaje("algun mensaje", "algun topico")
	if err != nil {
		t.Errorf("Se esperaba: %v, se obtuvo %v", nil, err)
	}
}

func TestConsumirMensaje(t *testing.T) {
	mensaje, err := LeerMensaje("")
	if err == nil {
		t.Errorf("Se esperaba: %v, se obtuvo %v", errors.New("el topico está vacio"), err)
	}
	if mensaje != "" {
		t.Errorf("Se esperaba: %v, se obtuvo %v", "", mensaje)
	}
	mensaje, err = LeerMensaje("algun topico")
	if err != nil {
		t.Errorf("Se esperaba: %v, se obtuvo %v", nil, err)
	}
	if mensaje == "" {
		t.Errorf("Se esperaba: %v, se obtuvo %v", "Algun mensaje", mensaje)
	}
}

//Test con stretchr/testify/assert
func TestProducirMensajeConAssert(t *testing.T) {

	err := EnviarMensaje("", "algún tópico")
	assert.Equal(t, errors.New("el mensaje está vacio"), err)

	err = EnviarMensaje("algun mensaje", "")
	assert.Equal(t, errors.New("el topico está vacio"), err)

	err = EnviarMensaje("algun mensaje", "algun topico")
	assert.Equal(t, nil, err)
}

func TestConsumirMensajeConAssert(t *testing.T) {
	mensaje, err := LeerMensaje("")
	assert.Equal(t, errors.New("el topico está vacio"), err)
	assert.Equal(t, "", mensaje)

	mensaje, err = LeerMensaje("algun topico")
	assert.Equal(t, nil, err)
	assert.Equal(t, "Algun mensaje", mensaje)
}
