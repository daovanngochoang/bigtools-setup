package hadoop

import (
	"bigtools_cli/utils"
	"encoding/xml"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func extractHadoop(hadoopVersion, appHome string) {

	fmt.Println(fmt.Sprintf("Extract hadoop-%s....", hadoopVersion))
	utils.RunScript(fmt.Sprintf("tar -xzf %s/hadoop-%s.tar.gz -C %s", appHome, hadoopVersion, appHome))
	utils.RunScript(fmt.Sprintf("mv %s/hadoop-%s %s/hadoop", appHome, hadoopVersion, appHome))

}

func downloadHadoop(hadoopVersion, appHome string) {

	//check whether hadoop tar is in the machine.
	folderPath := appHome + fmt.Sprintf("/hadoop-%s.tar.gz", hadoopVersion)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {

		//Download hadoop
		fmt.Println(fmt.Sprintf("Download hadoop-%s....", hadoopVersion))
		utils.RunScript("curl -O https://dist.apache.org/repos/dist/release/hadoop/common/KEYS")
		utils.RunScript(fmt.Sprintf("wget https://dlcdn.apache.org/hadoop/common/hadoop-%s/hadoop-%s.tar.gz -O %s/hadoop-%s.tar.gz", hadoopVersion, hadoopVersion, appHome, hadoopVersion))
	}

	// extract hadoop.
	extractHadoop(hadoopVersion, appHome)
}

type Property struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

type Configuration struct {
	XMLName  xml.Name   `xml:"configuration"`
	Property []Property `xml:"property"`
}

func writeConfig(configFolder, filename string, config Configuration) {

	// translate Config struct to xml output
	println("write config to xml files ...")
	configXmlOutput, err := xml.MarshalIndent(config, "", "\t")
	if err != nil {
		fmt.Println(fmt.Sprintf("Error marshaling %s XML: ", filename), err)
		return
	}

	cmd := fmt.Sprintf("echo '%s' > %s/%s", string(configXmlOutput), configFolder, filename)
	utils.RunScript(cmd)
}

func generateAndWriteHadoopConfig(masterHost, hadoopHome string) {

	////////////////////////////////////////////////
	/////////////// CONSTRUCT CONFIG ///////////////
	////////////////////////////////////////////////

	fmt.Println("Construct hadoop config ...")

	// construct coresite config
	coreSiteConfig := Configuration{}

	coreSiteConfig.Property = append(
		coreSiteConfig.Property,
		Property{Name: "fs.defaultFS", Value: fmt.Sprintf("hdfs://%s:9000", "0.0.0.0")},
	)

	// Construct hdfSite confid
	hdfsSiteConfig := Configuration{}

	// add config values.
	hdfsSiteConfig.Property = append(
		hdfsSiteConfig.Property,
		Property{Name: "dfs.namenode.name.dir", Value: fmt.Sprintf("file://%s/hadoop/data/nameNode", os.Getenv("HOME"))},
	)
	hdfsSiteConfig.Property = append(
		hdfsSiteConfig.Property,
		Property{Name: "dfs.datanode.data.dir", Value: fmt.Sprintf("file://%s/hadoop/data/dataNode", os.Getenv("HOME"))},
	)
	hdfsSiteConfig.Property = append(
		hdfsSiteConfig.Property,
		Property{Name: "dfs.replication", Value: "2"},
	)

	// construct yarnSite config
	mapredSiteConfig := Configuration{}

	mapredSiteConfig.Property = append(
		mapredSiteConfig.Property,
		Property{Name: "mapreduce.framework.name", Value: "yarn"},
	)
	mapredSiteConfig.Property = append(
		mapredSiteConfig.Property,
		Property{Name: "yarn.app.mapreduce.am.env", Value: "HADOOP_MAPRED_HOME=" + hadoopHome},
	)
	mapredSiteConfig.Property = append(
		mapredSiteConfig.Property,
		Property{Name: "mapreduce.map.env", Value: "HADOOP_MAPRED_HOME=" + hadoopHome},
	)
	mapredSiteConfig.Property = append(
		mapredSiteConfig.Property,
		Property{Name: "mapreduce.reduce.env", Value: "HADOOP_MAPRED_HOME=" + hadoopHome},
	)

	// construct yarnSite config
	yarnSiteConfig := Configuration{}

	yarnSiteConfig.Property = append(
		yarnSiteConfig.Property,
		Property{Name: "yarn.acl.enable", Value: "0"},
	)

	yarnSiteConfig.Property = append(
		yarnSiteConfig.Property,
		Property{Name: "yarn.resourcemanager.hostname", Value: masterHost},
	)

	yarnSiteConfig.Property = append(
		yarnSiteConfig.Property,
		Property{Name: "yarn.nodemanager.aux-services", Value: "mapreduce_shuffle"},
	)
	yarnSiteConfig.Property = append(
		yarnSiteConfig.Property,
		Property{Name: "yarn.nodemanager.aux-services.mapreduce.shuffle.class", Value: "org.apache.hadoop.mapred.ShuffleHandler"},
	)

	////////////////////////////////////////////////////////////////
	/////// WRITE HADOOP CONFIG TO THE hadoop_config FOLDER ////////
	////////////////////////////////////////////////////////////////

	fmt.Println("Write hadoop config to file...")
	configFolder := hadoopHome + "/etc/hadoop"

	// write core-site.xml
	writeConfig(configFolder, "core-site.xml", coreSiteConfig)

	// write hdfs-site.xml
	writeConfig(configFolder, "hdfs-site.xml", hdfsSiteConfig)

	// write mapred-site.xml
	writeConfig(configFolder, "mapred-site.xml", mapredSiteConfig)

	// write yarn-site.xml
	writeConfig(configFolder, "yarn-site.xml", yarnSiteConfig)

	utils.RunScript(fmt.Sprintf("echo 'export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64' >> %s/etc/hadoop/hadoop-env.sh", hadoopHome))

}

var (
	masterIp  string
	workerIps []string

	InstallCmd *cobra.Command = &cobra.Command{
		Use:   "install -m [master's ip] -w [worker's ip 1] -w [worker's ip n] ",
		Short: "Install hadoop on single and multi-nodes cluster",
		Long: "This command allow you to install hadoop on multiple node from the master machine." +
			"\nIn case you don't provide worker's ips then it will install hadoop on single node",
		Args:    cobra.MinimumNArgs(0),
		Example: "  bigtools hadoop install -m 172.16.96.103 -w 172.16.96.104 -w 172.16.96.105 ...",
		Run: func(cmd *cobra.Command, args []string) {

			println("m ip: ", masterIp)
			if masterIp != "" && len(workerIps) > 0 {
				inputIps := append(workerIps, masterIp)

				//check ip format!
				for _, ip := range inputIps {
					if !utils.IsIpv4(ip) {
						println(fmt.Sprintf("IP: %s is not valid!", ip))
						return
					}
				}
			}

			// read
			clusterConfig := utils.ReadClusterInformation(masterIp, workerIps)

			// if worker not add => install on single node.
			if len(clusterConfig.Slave) == 0 {
				println("No worker's ip is found, Installing on single node ....")

			} else {
				// if there is any worker -> install on multi-node cluster.
				println("Installing hadoop on multi nodes")
				println("Master IP: " + clusterConfig.Master.Ip)
				println(fmt.Sprintf("Install hadoop on %d worker", len(clusterConfig.Slave)))
			}

			// update
			utils.RunScript("sudo apt-get update")
			utils.RunScript("sudo apt-get install openjdk-8-jdk net-tools curl netcat gnupg openssh-server libsnappy-dev -y")

			config := getHadoopConfig()
			config.Cluster = clusterConfig

			// create folders
			println("Create app folder ...")
			utils.RunScript(fmt.Sprintf("mkdir -p %s", config.AppConfig.AppHome))

			// Download and extract hadoop
			downloadHadoop(config.Version, config.AppConfig.AppHome)

			// generate and write xml config to hadoop config.
			generateAndWriteHadoopConfig(clusterConfig.Master.Host, config.ToolHome)

			// write ip and host workers
			if clusterConfig.Master.Host != "localhost" {
				utils.RunScript(fmt.Sprintf("echo %s > %s", clusterConfig.Master.Host, config.ToolHome+"/etc/hadoop/masters"))
				utils.RunScript(fmt.Sprintf("echo '' > %s", config.WorkerFilePath))

				//write to /etc/hosts workers, masters file in hadoop folder
				utils.RunScript(fmt.Sprintf("sudo sed -i '/%s/d' /etc/hosts", clusterConfig.Master.Host))
				utils.RunScript(fmt.Sprintf("echo '%s        %s' | sudo tee -a /etc/hosts", clusterConfig.Master.Ip, clusterConfig.Master.Host))

				for _, slave := range clusterConfig.Slave {
					utils.RunScript(fmt.Sprintf("echo '%s        %s' | sudo tee -a /etc/hosts", slave.Ip, slave.Host))
					utils.RunScript(fmt.Sprintf("echo %s >> %s", slave.Host, config.WorkerFilePath))

				}
			}

			env := getEnv(config.ToolHome)
			utils.RunScript(fmt.Sprintf("echo '%s' >> ~/.bashrc", env))

			// send setup to workers
			setupWorkers(config, env)

		},
	}
)
