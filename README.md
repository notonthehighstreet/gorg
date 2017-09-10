# gorg

CLI to interact with mesos, marathon, chronos and others in GO.
Useful if you need to jump around throwable QA environments.

# Commands
 - `init`     Initialise gorg configuration file
 - `config`   [show, add, remove] Modify current configuration of your gorg
 - `ssh`      Set provided user for ssh
 - `use`      Switch default environment
 - `service`  [ls, show, open] Interact with running services according to consul

# Install

```
    $ go get github.com/notonthehighstreet/gorg
    $ make install
```

You then need to run the init command that would generate a json file.
The file will store services info for each QA environment.

```
    $ gorg init --environment-name envname --domain domainname
```

# Requirement

 - Installed Go in your machine
 - GOPATH set correctly

# WIP

 - ssh to QA boxes
 - listing KV
 - ps through mesos
 - show logs
 - show services statuses
