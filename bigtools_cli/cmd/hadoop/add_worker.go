package hadoop

import (
	"bigtools_cli/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	ipaddr     string
	AddNodeCmd = &cobra.Command{
		Use:   "add-worker -i [worker's ipaddr]",
		Short: "Add new worker to the existing cluster.",
		Long:  "This command allow you to add a new worker to the existing cluster",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {

			// get hadoop config
			config := getHadoopConfig()

			//check whether hadoop is installed in this machine
			folderPath := config.ToolHome
			if _, err := os.Stat(folderPath); os.IsNotExist(err) {
				fmt.Println("Hadoop is not install in this machine!")
				return
			}

			// add master info to config
			config.Cluster = &utils.ClusterInformation{
				Master: utils.HostInfo{},
				Slave:  nil,
			}
			config.Cluster.Master.Ip, config.Cluster.Master.Host = utils.GetMasterInfo()

			// check whether the ip addr is valid
			if utils.IsIpv4(ipaddr) {

				//get hostname
				hostname := utils.GetRemoteHostname(ipaddr)

				// add to network
				utils.RunScript(fmt.Sprintf("echo '%s        %s' | sudo tee -a /etc/hosts", ipaddr, hostname))

				// add to hadoop workers
				utils.RunScript(fmt.Sprintf("echo %s >> %s", hostname, config.WorkerFilePath))
				setupWorker(ipaddr, config, getEnv(config.ToolHome))
			} else {
				println("invalid ip address.")
			}
		},
	}
)
