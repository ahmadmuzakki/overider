# Overider
Override structured file value. For now it just support .ini file.


### Download
```bash
wget https://github.com/ahmadmuzakki/overider/raw/master/bin/overider
```
### Syntax
```bash
overider <source> <target> [-o <output>]
```
### Use Case
Consider we have `config.ini` file like this.

```ini 
#config for Service1
[Service1]
Endpoint = https://service1.com
Path = /somepath

#config for Service2
[Service2]
Endpoint = https://service2.com
Path = /anotherpath
```

and we want to change only Endpoint of Service1 and Service2 with another value. so we'd create new file let's say `mock.ini`

```ini
[Service1]
Endpoint = http://mock1.com

[Service2]
Endpoint = http://mock2.com
```

and we can override the value using `overider`
```bash
overider mock.ini config.ini
```

this will update `config.ini` into
```ini
#config for Service1
[Service1]
Endpoint = http://mock1.com
Path = /somepath

#config for Service2
[Service2]
Endpoint = http://mock2.com
Path = /anotherpath
```

if you want to output to another file you can do 

`overider mock.ini config.ini -o output.ini`
