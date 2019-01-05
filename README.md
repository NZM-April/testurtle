# testurtle!

testurtle is Infrastructure Testing Micro Engine.

# motivation!

Infrastructure is very complicated.<br>
DNS, Firewall, Security Group, Web Server...<br>
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
        {"cmd": "echo 'oknum is $oknum'"},
        {"cmd": "echo 'ngnum is $ngnum'"},
        {"cmd": "echo 'msg is $msg'"}
    ]
}
```

# run!

```
$ testurtle run -f turtleconfig.json
```

![image](https://user-images.githubusercontent.com/38576286/49904495-b38d9d80-fead-11e8-8ccb-e7d849e1317a.png)

# Copyright!
Copyright (c) 2019 NZM-April(Ohki Nozomu)
