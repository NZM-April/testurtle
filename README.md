# testurtle!

testurtle is Infrastructure Testing Micro Engine.

# motivation!

Infrastructure is very complicated.<br>
Firewall, WAF, Security Group, Web Server...<br>
Testing is necessary to ensure its reliability.<br>
However, writing test code is a big task.<br>
testurtle provides a minimal infrastructure test.<br>

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

```
{
    "items": [
        {"URL": "https://www.hashicorp.com/", "contain": "Cloud"},
        {"URL": "https://github.com/", "title": "GitHub"},
        {"URL": "https://www.google.com/", "status": 200},
        {"URL": "https://www.cncf.io", "contain": "Open", "title": "Native", "status": 200}
    ],
    "notifications": [
        {"cmd": "echo 'ok'"}
    ]
}
```

# run!

```
$ testurtle run -f turtleconfig.json
```

# Copyright!
Copyright (c) 2018 Ohki Nozomu
