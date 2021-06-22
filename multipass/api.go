package multipass

import (
	"encoding/json"
	"fmt"
)

type InstanceInfo struct {
	Ipv4    []string
	Name    string
	Release string
	State   string
}
type InstanceInfoList struct {
	List []*InstanceInfo
}

func GetAllInfo() []*InstanceInfo {
	all := listAll()
	var infos InstanceInfoList
	err := json.Unmarshal([]byte(all), &infos)
	if err != nil {
		panic(fmt.Sprintf("GetAll Info error: %v", err))
	}
	return infos.List
}

func RunningInstance() []*InstanceInfo {
	all := GetAllInfo()
	var runningInstances []*InstanceInfo
	for _, info := range all {
		if (*info).State == "Running" {
			runningInstances = append(runningInstances, info)
		}
	}
	return runningInstances
}

func IsRunning(name string) bool {
	instances := RunningInstance()
	for _, instance := range instances {
		if instance.Name == name {
			return true
		}
	}
	return false
}

func GetInstance(name string) *InstanceInfo {
	for _, info := range GetAllInfo() {
		if info.Name == name {
			return info
		}
	}
	panic("No instance found")
}

func GetIP(name string) (string, error) {
	running := IsRunning(name)
	if !running {
		return "", fmt.Errorf("GetIP failed, distro '%s' is not running", name)
	}
	return GetInstance(name).Ipv4[0],nil
}
