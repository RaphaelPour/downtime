![](downtime.png)

# downtime

Measures downtime by detecting the start and end of ICMP request timeouts (like `ping -D -W 0.05 -i 0.1 <target>`).

The tool needs `sudo` and the net raw capability to handle raw ICMP pakets:

```bash
sudo setcap cap_net_raw+ep <path to binary>
```

## Quick start

1. install: `go install github.com/RaphaelPour/downtime`
2. set capability: `sudo setcap cap_net_raw+ep <path to binary>`
3. run: `sudo downtime --target <target>`

example:
```
   start: 2023-08-16 10:12:47.197488336 +0200 CEST m=+20.651438246
     end: 2023-08-16 10:13:06.331099665 +0200 CEST m=+39.785050063
duration: 19.133611817s
```

## Usage

```bash
Usage of downtime:
      --duration duration   ICMP request interval (default 100ms)
      --target ip           target IP
      --timeout duration    ICMP reply timeout (default 50ms)
```
