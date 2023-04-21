package hadoop

import (
	"github.com/spf13/cobra"
)

var (
	// Commands Hadoop Used for flags.
	Commands = &cobra.Command{
		Use:   "hadoop [command] ",
		Short: "Hadoop installation cli-app",
		Long:  `This CLI app help to simplify the installation of hadoop`,
	}
)

func init() {

	// add flag for install command
	// add flag for install command
	InstallCmd.Flags().StringVarP(&masterIp, "master-ip", "m", "", "public ip addresses of the master node")
	InstallCmd.Flags().StringArrayVarP(&workerIps, "worker-ips", "w", []string{}, "ip addresses of the worker nodes")

	// add flag for add-worker command.
	AddNodeCmd.Flags().StringVarP(&ipaddr, "worker-ip", "i", "", "ip address of the new worker")

	// add flag to the delete node cmd
	DeleteNodeCmd.Flags().StringArrayVarP(&delNode, "del-nodes", "l", []string{}, "list of to-delete ip addresses ")

	Commands.AddCommand(InstallCmd)
	Commands.AddCommand(AddNodeCmd)
	Commands.AddCommand(UninstallCmd)
	Commands.AddCommand(DeleteNodeCmd)
}
