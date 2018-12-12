# testurtle!

testurtle is Infrastructure Testing Micro Engine.

# motivation!

Infrastructure is very complicated.
Firewall, WAF, Security Group, Web Server...
Testing is necessary to ensure its reliability.
However, writing test code is a big task.
testurtle provides a minimal infrastructure test.

# Installation!

install go, dep

```
$ git clone
$ dep ensure
$ go build -o testurtle
$ mv testurtle /usr/local/bin
```

# Configuration!

```
$ vi turtleconfig.json
```

# run!

```
$ testurtle run -f turtleconfig.json
```

# Copyright
Copyright (c) 2018 Ohki Nozomu
