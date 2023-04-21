package hadoop

import (
	"bigtools_cli/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	delNode       []string
	DeleteNodeCmd = &cobra.Command{
		Use:   "del-nodes -h hostname1 -h hostname2 -h hostname_n ... ",
		Short: "delete nodes from the cluster",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			println("Start uninstalling process ...")

			// get hadoop config
			config := getHadoopConfig()

			//check whether hadoop is installed in this machine
			folderPath := config.ToolHome
			if _, err := os.Stat(folderPath); os.IsNotExist(err) {
				fmt.Println("Hadoop is not install in this machine!")
				return
			}

			command := uninstallCommands{
				rmJavaEnv:        "sed -i '/JAVA_HOME/d' ~/.bashrc && sed -i '/java/d' ~/.bashrc",
				rmHadoopEnv:      "sed -i '/HADOOP/d' ~/.bashrc && sed -i '/hadoop/d' ~/.bashrc",
				rmHadoop:         "rm -r " + config.ToolHome,
				rmBigtoolsFolder: fmt.Sprintf("rm %s -rf", config.AppConfig.AppHome),
			}

			for _, ip := range delNode {
				if !utils.IsIpv4(ip) {
					println("Invalid ip address : ", ip)
					return
				}
			}

			// delete node
			for _, ip := range delNode {

				// get hostname of the ip address
				hostname := utils.GetRemoteHostname(ip)

				// uninstall hadoop on remote machine
				remoteUninstall(hostname, command)

				// remove from worker list
				utils.RunScript(fmt.Sprintf("sed -i '/%s/d' %s", hostname, config.WorkerFilePath))
				utils.RunScript(fmt.Sprintf("sudo sed -i '/%s/d' /etc/hosts", hostname))

			}
		},
	}
)
