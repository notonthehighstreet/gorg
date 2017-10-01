# gorg

CLI to interact with mesos, marathon, chronos and others in GO.
Useful if you need to jump around throwable QA environments.

# Commands
 - `init`     Initialise gorg configuration file
 - `console`  ssh into a service box
 - `config`   [show, add, remove] Modify current configuration of your gorg
 - `ssh`      Set provided user for ssh
 - `use`      Switch default environment
 - `kv`       [ls, get] Interact with KVs stored in consul
 - `service`  [ls, show, open] Interact with running services according to consul

# Install

Download binary files from the `download` folder.
Choose the one suitable for your machine.

You then need to run the init command that would generate a json file.
The file will store services info for each QA environment.

```
    $ gorg init --domain domainname
```

Start your first environment:

```
    $ gorg config add your_environment_name
```

All ready for you!

# WIP

 - ps through mesos
 - show logs
 - show services statuses
