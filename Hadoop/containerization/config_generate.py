# read file host and determine master and slaves.
# generate master and slave config xml files
# compose for master
# compost for slave




from config import BaseConfig
import io
import yaml

class ConfigGenerator:
    def __init__(self) -> None:
        super().__init__()

        hosts = open("hosts", "r")
        self.domains=[]
        for host in hosts:
            host = host.replace("\n", "")
            host = host.replace(" ", "")
            self.domains.append(host)
        
        self.master = self.domains[0]
        self.slaves = self.domains[1:]

        self.config = BaseConfig(self.master, self.domains)




    def write_docker(self, path, master=True):

        master_set_up="{}set-up.sh".format("master-" if master == True else "slave-")
        
        value = self.config.DOCKER_FILE + """\nCOPY ../{1}_config /
            \nCOPY {0} / \nRUN ["bash", "./{0}"]""".format(master_set_up, "master" if master else "slave")
        self.write_file(path, value)



     
    def write_file(self, path, value):
        file = open(path, "w")
        file.write(value)
        file.close()



    def write_config(self, master=True):
        
        self.write_docker(master)
        config_files = {
            "core-site":self.config.core_site_config, 
            "hdfs-site":self.config.hdfs_site_config, 
            "mapred-site":self.config.mapred_site_config, 
            "yarn-site":self.config.yarn_site}

        for f in config_files:
            path = "{0}_config/{1}.xml".format("master" if master else "slave", f)
            self.write_file(path, config_files[f])
    
    
    
    def write_docker_compose(self):
        path = 'docker-compose-test.yml'
        self.write_file(path, "version: '3'\n")
        # Write YAML file
        with open(path, 'a') as yaml_file:
            yaml.dump(self.config.compose, yaml_file)
        return

    def go(self):
        self.write_docker("dockerfile/{}Dockerfile".format(self.master))
        self.write_config()
        self.write_config(False)
        
        for i in self.slaves:
            path = "dockerfile/{}Dockerfile".format(i)
            self.write_docker(path, master=False)
        
        self.write_docker_compose()
    


test = ConfigGenerator()
test.go()
