package hadoop

import (
	"bigtools_cli/utils"
	"fmt"
)

func getHadoopConfig() utils.Config {
	config := utils.GetAppConfig()
	hadoopHome := config.AppHome + "/hadoop"

	return utils.Config{
		WorkerFilePath: fmt.Sprintf("%s/etc/hadoop/workers", hadoopHome),
		MasterFilePAth: fmt.Sprintf("%s/etc/hadoop/masters", hadoopHome),
		ToolHome:       hadoopHome,
		Version:        "3.3.1",
		Cluster:        nil,
		AppConfig:      config,
	}
}

func setupWorker(workerIp string, config utils.Config, env string) {

	// update worker and install necessary dependencies
	utils.RunRemoteScriptSudo(workerIp, fmt.Sprintf("echo '\n\nProvide password for running sudo operations on slave machine IP: %s' ", workerIp)+
		`&& sudo apt update && sudo apt-get install openjdk-8-jdk net-tools curl netcat gnupg openssh-server libsnappy-dev -y`)

	// create appHome in the target machine.
	utils.RunRemoteScript(workerIp, fmt.Sprintf("mkdir -p %s", config.AppConfig.AppHome))

	// write env
	utils.RunRemoteScript(workerIp, fmt.Sprintf("echo '%s' >> ~/.bashrc", env))

	// add master host to slave
	utils.RunRemoteScriptSudo(workerIp, fmt.Sprintf("echo '%s        %s' | sudo tee -a /etc/hosts", config.Cluster.Master.Ip, config.Cluster.Master.Host))

	// copy setup to the target machine.
	utils.RunScript(fmt.Sprintf("scp -r -q %s %s:%s", config.ToolHome, workerIp, config.AppConfig.AppHome))
	// delete worker file in the target machine.
	utils.RunRemoteScript(workerIp, fmt.Sprintf("echo 'localhost' > %s", config.WorkerFilePath))
}

func setupWorkers(config utils.Config, env string) {
	for _, info := range config.Cluster.Slave {
		if utils.IsIpv4(info.Ip) {
			setupWorker(info.Ip, config, env)
		} else {
			println(info.Ip + " is invalid!")
		}
	}
}

// get ENV variable
func getEnv(hadoopHome string) string {
	javaEnv := "#config java home paths.\nexport JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64\nexport PATH=$PATH:$JAVA_HOME/bin\n\n"
	hadoopEnv := fmt.Sprintf("#config for hadoop paths.\nexport HADOOP_HOME=%s\nexport HADOOP_INSTALL=$HADOOP_HOME\nexport HADOOP_MAPRED_HOME=$HADOOP_HOME\nexport HADOOP_COMMON_HOME=$HADOOP_HOME\nexport HADOOP_HDFS_HOME=$HADOOP_HOME\nexport YARN_HOME=$HADOOP_HOME\nexport HADOOP_COMMON_LIB_NATIVE_DIR=$HADOOP_HOME/lib/native\nexport PATH=$PATH:$HADOOP_HOME/sbin:$HADOOP_HOME/bin\nexport HADOOP_OPTS=\"-Djava.library.path=$HADOOP_HOME/lib/native\"", hadoopHome)
	return javaEnv + hadoopEnv
}

// remote hadoop on remote machine
func remoteUninstall(ip string, command uninstallCommands) {
	println("Remove hadoop on worker with the ipaddr: " + ip)
	utils.RunRemoteScript(ip, command.rmJavaEnv)
	utils.RunRemoteScript(ip, command.rmHadoopEnv)
	utils.RunRemoteScript(ip, command.rmHadoop)
	utils.RunRemoteScript(ip, command.rmBigtoolsFolder)
}
