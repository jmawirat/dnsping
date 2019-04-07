# dnsping

DNSPING is a go application that allows the user to measure the performance on how the users dns is resolving IP address to Domain name and Domain name to IP address.

# Usage

For MAC OS X:

```bash
./dnsping-darwin-amd64 facebook.com

resolved facebook.com to 157.240.15.35 resolution_seq=0 time=69.169314ms
resolved facebook.com to 157.240.15.35 resolution_seq=1 time=2.481586ms
resolved facebook.com to 157.240.15.35 resolution_seq=2 time=2.428151ms
resolved facebook.com to 157.240.15.35 resolution_seq=3 time=2.704366ms
resolved facebook.com to 157.240.15.35 resolution_seq=4 time=2.836854ms
resolved facebook.com to 157.240.15.35 resolution_seq=5 time=2.491492ms
```

For Windows

```bash
dnsping-windows-386.exe facebook.com

resolved facebook.com to 157.240.15.35 resolution_seq=0 time=69.169314ms
resolved facebook.com to 157.240.15.35 resolution_seq=1 time=2.481586ms
resolved facebook.com to 157.240.15.35 resolution_seq=2 time=2.428151ms
resolved facebook.com to 157.240.15.35 resolution_seq=3 time=2.704366ms
resolved facebook.com to 157.240.15.35 resolution_seq=4 time=2.836854ms
resolved facebook.com to 157.240.15.35 resolution_seq=5 time=2.491492ms
```

```bash
dnsping-windows-amd64.exe facebook.com

resolved facebook.com to 157.240.15.35 resolution_seq=0 time=69.169314ms
resolved facebook.com to 157.240.15.35 resolution_seq=1 time=2.481586ms
resolved facebook.com to 157.240.15.35 resolution_seq=2 time=2.428151ms
resolved facebook.com to 157.240.15.35 resolution_seq=3 time=2.704366ms
resolved facebook.com to 157.240.15.35 resolution_seq=4 time=2.836854ms
resolved facebook.com to 157.240.15.35 resolution_seq=5 time=2.491492ms
```

# Roadmap

1. Add Maximum, Minimum and Average Resolution time.
2. Add Continues resolution like ping's -t option.
