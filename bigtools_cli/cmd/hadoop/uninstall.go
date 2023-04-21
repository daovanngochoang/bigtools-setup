package hadoop

import (
	"bigtools_cli/utils"
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type uninstallCommands struct {
	rmJavaEnv        string
	rmHadoopEnv      string
	rmHadoop         string
	rmBigtoolsFolder string
}

var UninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Short:   "uninstall this hadoop cluster ",
	Example: "   bigtools hadoop uninstall",
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		config := getHadoopConfig()

		folderPath := config.ToolHome
		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			fmt.Println("Hadoop is not install in this machine!")
			return
		}

		println("Start uninstalling process ...")

		command := uninstallCommands{
			rmJavaEnv:        "sed -i '/JAVA_HOME/d' ~/.bashrc && sed -i '/java/d' ~/.bashrc",
			rmHadoopEnv:      "sed -i '/HADOOP/d' ~/.bashrc && sed -i '/hadoop/d' ~/.bashrc",
			rmHadoop:         "rm -r " + config.ToolHome,
			rmBigtoolsFolder: fmt.Sprintf("rm  %s -rf", config.AppConfig.AppHome),
		}

		// delete on worker nodes.
		file, err := os.Open(config.WorkerFilePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		// Create a scanner
		scanner := bufio.NewScanner(file)

		// read the worker workerHost and uninstall
		workerHost := ""
		for scanner.Scan() {
			workerHost = scanner.Text()
			println("\n\nIP: ", workerHost, "\n\n")
			if len(workerHost) > 0 {
				remoteUninstall(workerHost, command)
				utils.RunScript(fmt.Sprintf("sudo sed -i /%s/d /etc/hosts", workerHost))
			}
		}

		// if the master workerHost isn't the current machine but on other machine => uninstall remotely else do uninstallation locally
		// remove hadoop on master
		println(" remove hadoop on master ....")
		utils.RunScript(command.rmJavaEnv)
		utils.RunScript(command.rmHadoopEnv)
		utils.RunScript(command.rmHadoop)
		utils.RunScript(command.rmBigtoolsFolder)

		println("Done!")
	},
}
