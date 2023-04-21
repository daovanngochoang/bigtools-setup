# Bigtools CLI app.
This app created to simplify bigdata tools installation process which is tested and run on debian based linux (ubuntu 20.04 or higher is recommended).

### Contents
1. [Create User and establish ssh connection](#setup-process)
2. [Install Hadoop](#install-hadoop)


## Setup process.
All machine should have the ability to connect to each other via ssh, 
therefore we have to generate key in the master and copy it to the worker machine.

All machine show have the same username (hadoop is recommended).

### Step 2: Add username.
```bash
sudo useradd -m hadoop
sudo usermod -aG sudo hadoop
sudo passwd hadoop
```

### Step 2:  Generate SSH key and test ssh
```bash
ssh-keygen -t rsa -P '' -f ~/.ssh/id_rsa
cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys
chmod 0600 ~/.ssh/authorized_keys
```
After generate key we are able to connect to our machine with ssh
```bash
ssh localhost
```
However, to establish connections with other machines we have to copy the ssh key to the target machines.

2. Copy the key to the worker machines.
```bash
scp ~/.ssh/authorized_keys worker_ip:~/.ssh/authorized_keys
```
*example:* *if I have 1 master and 3 workers*
```bash
scp ~/.ssh/authorized_keys worker-1:~/.ssh/authorized_keys
scp ~/.ssh/authorized_keys worker-2:~/.ssh/authorized_keys
scp ~/.ssh/authorized_keys worker-3:~/.ssh/authorized_keys
```
3. Test SSH and all shoud be work without password
```bash
ssh master
ssh slave-1
ssh slave-2
ssh slave-3
```
### Step 3: Download installation.
```bash
wget https://github.com/daovanngochoang/bigtools-setup/raw/main/bigtools_cli/bin/bigtools.run
```
# App overview.
*To see the usage of the app run this:*
```bash
./bigtools.run -h 
```
Output
```text
Usage:
  bigtools [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  hadoop      Hadoop installation cli-app
  help        Help about any command

Flags:
  -h, --help   help for bigtools

Use "bigtools [command] --help" for more information about a command.
```

## Install hadoop
#### NOTE: make sure you complete [Create User and establish ssh connection](#setup-process) process 

*To see more about commands support to hadoop tool run :*
```bash
./bigtools.run hadoop -h 
```
Output
```text
This CLI app help to simplify the installation of hadoop

Usage:
  bigtools hadoop [command]

Available Commands:
  add-worker  Add new worker to the existing cluster.
  del-nodes   delete nodes from the cluster
  install     Install hadoop on single and multi-nodes cluster
  uninstall   uninstall this hadoop cluster 

Flags:
  -h, --help   help for hadoop

Use "bigtools hadoop [command] --help" for more information about a command.

```

### Install single node cluster hadoop
To install a single node hadoop cluster run this command without any flag.
It will automatically install for on the current machine. 
```bash
./bigtools.run hadoop install
```

### Install a multi nodes hadoop cluster.
To see how to use this command run 
```bash
./bigtools.run hadoop install -h
```
Output
```text
This command allow you to install hadoop on multiple node from the master machine.
In case you don't provide worker's ips then it will install hadoop on single node

Usage:
  bigtools hadoop install -m [master's ip] -w [worker's ip 1] -w [worker's ip n]  [flags]

Examples:
  bigtools hadoop install -m 172.16.96.103 -w 172.16.96.104 -w 172.16.96.105 ...

Flags:
  -h, --help                     help for install
  -m, --master-ip string         public ip addresses of the master node
  -w, --worker-ips stringArray   ip addresses of the worker nodes

```
The output contains the Example for installing hadoop on multi nodes, you have to get the public
ipaddress of the master machine and put it after the -m flag and add multiple worker's ips after multiple -w flag respectively.

Run this command on the master machine, it will take care of the rest of installation process on both master & worker machines .
