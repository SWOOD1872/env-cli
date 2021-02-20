# env-cli

Playing around with the env package by caarlos0.

Here, I use the env package to parse the HOME environment variable, then use it as a default value for a flag.

### Build
```bash
go build .
```

### Run
##### Using default value parsed by the env package
```bash
-> ./env-cli 
Directory = <HOME>
```

##### Using flag value
```bash
-> ./env-cli -dir /tmp
Directory = /tmp
```