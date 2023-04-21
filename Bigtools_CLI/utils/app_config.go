package utils

import (
	"fmt"
	"os"
)

type AppConfig struct {
	AppHome      string
	ConfigFolder string
	Version      string
}

type Config struct {
	WorkerFilePath string
	MasterFilePAth string
	ToolHome       string
	Version        string
	AppConfig      AppConfig
	Cluster        *ClusterInformation
}

func GetAppConfig() AppConfig {
	appFolderName := "bigtools"
	appHome := fmt.Sprintf("%s/%s", os.Getenv("HOME"), appFolderName)
	configFolder := fmt.Sprintf("%s/.%s", os.Getenv("HOME"), appFolderName)

	return AppConfig{
		AppHome:      appHome,
		ConfigFolder: configFolder,
		Version:      "1.0.0",
	}
}

type HostInfo struct {
	Ip   string
	Host string
}

type ClusterInformation struct {
	Master HostInfo
	Slave  []HostInfo
}

func ReadClusterInformation(masterIp string, workerIps []string) *ClusterInformation {

	// Unmarshal JSON content into Config struct
	var config ClusterInformation

	// make sure host name are lowercase.
	if masterIp != "" {
		config.Master.Host, _ = os.Hostname()
		config.Master.Ip = masterIp
	} else {
		config.Master.Host = "localhost"
		config.Master.Ip = "127.0.0.1"
	}

	for _, ip := range workerIps {
		config.Slave = append(config.Slave, HostInfo{
			Ip:   ip,
			Host: GetRemoteHostname(ip),
		})
	}

	return &config
}
