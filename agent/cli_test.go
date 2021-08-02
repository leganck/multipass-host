package agent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListAll(t *testing.T) {
	all := listAll()
	println(all)
	assert.NotNil(t, all)
}

func TestDefaultRunCommand(t *testing.T) {
	command, err := DefaultRunCommand("ls", "-lah", "~")
	assert.Nil(t, err)
	println(command)
	assert.NotNil(t, command)
}

func TestDefaultRunCommandError(t *testing.T) {
	_, err := DefaultRunCommand("cat", "/.bashrc")
	assert.NotNil(t, err)
}
