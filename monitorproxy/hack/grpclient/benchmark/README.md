# gRPC benchmark client

## Compile gRPC benchmark client

```sh
go build
```

## Set ulimit

```sh
ulimit -n 65535
```

## Run gRPC benchmark client

Test with `10 sites` and `report period 5 secs` and `live period 1 sec` with duration time `60 seconds`.

```sh
./benchmark -ca=<path-to-ca> -cert=<path-to-client-pem> -key=<path-to-client-key> -server=<ip:port> -sites=10 -report_period=5 -live_period=1 -duration=60
```

## Monitor CPU usage of monitorproxy

a.Install `htop` or `sysstat`

```sh
apt update
apt install -y sysstat htop
```

b.Run `sar` command

It will print CPU information every 1 second for 10,000 counts.

```sh
sar 1 10000
```

c.Continuosly watch metrics.log or events.log size change every 1 sec

```sh
while true; do du -s /var/log/hostlog/metrics.log ; sleep 1; done
while true; do du -s /var/log/hostlog/events.log ; sleep 1; done
```
