# toolbox - config

It's a lightweight toolbox to interact with configuration files such as Yaml, JSON, TOML... For example, the toolbox can :

- Read a file 
- Read a path at the file (user.personal.name) 
- Convert from a format to another one (JSON -> TOML, TOML -> YAML ...) 
- Remove comments
- write a file  (tb -w deployment.yaml -p "user.personal.name=user" -p "user.personal.last=last" -p "user.secret=true")
- Override parameters at the file (tb -o deployment.yaml -p "metadata.app.version=NewVersionY")
- Override parrameters merging files (tb -o deployment.yaml -m newParameter.yaml)

## It's a tool to help me handle configuration files without using regex replacement everywhere and for everything.

## Notes :
Code:: Golang 1.18
License:: GNU GPLv3
Author:: Davi Mello

Feel free to use and improve.
