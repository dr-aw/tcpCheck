# This is a simple TCP checker.

### You can use this program, if you need to log, when your port is sleeping.

-----------------
Usage: `go run tcpCheck.go [-l latency] [-tok] [-tnot] <host> <port>`


Example: `go run tcpCheck.go 192.168.88.1 22`

or you can use some flags:

`go run tcpCheck.go -l 300 -tok 10 -tnot 1 192.168.88.1 22`



where:
* `-l` means latency (in milliseconds)
* `-tok` means time to sleep after OK connection (in seconds)
* `-tnot` means time to sleep after NOT OK connection (in seconds)
-----------------
##### Enjoy :)