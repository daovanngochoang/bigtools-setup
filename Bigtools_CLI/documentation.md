# Bigtools ClI app.
This app created to simplify hadoop installation process which is tested and run on debian based linux (ubuntu 20.04 or higher is recommended).

### Setup process.
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
However, to establish connections with other machine we have to copy the ssh key to the other machine.

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
3. Test SSH and all shoud be work without
```bash
ssh master
ssh slave-1
ssh slave-2
ssh slave-3
```
### Step 3: Download installation.
```bash
```
### Step 3: Install hadoop
```bash

```