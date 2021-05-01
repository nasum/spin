package entity

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCreateUser(t *testing.T) {
	user := CreateUser("test")
	assert.Equal(t, user.Name, "test")
}
