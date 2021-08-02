package agent

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllInfo(t *testing.T) {
	infos := GetAllInfo()
	assert.NotNil(t, infos)
	for _, info := range infos {
		fmt.Printf("%v\n", *info)
	}
	assert.Len(t, infos, 1)
}

func TestRunningInstance(t *testing.T) {
	runnings := RunningInstance()
	for _, running := range runnings {
		fmt.Printf("%v", *running)
	}
	assert.NotNil(t, runnings)

}

func TestIsRunning(t *testing.T) {
	running := IsRunning("primary")
	assert.True(t, running)

}

func TestGetInstance(t *testing.T) {
	instance := GetInstance("primary")
	fmt.Printf("%v\n", *instance)
	assert.NotNil(t, *instance)
}

func TestGetIP(t *testing.T) {
	ip, err := GetIP("primary")
	fmt.Printf("%s\n", ip)
	assert.Nil(t, err)
	assert.NotNil(t, ip)
}
