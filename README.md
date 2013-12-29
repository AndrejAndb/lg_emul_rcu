lg_emul_rcu
===========

Alternative script for Control LG smartTV emulator (2012/2013)

Buld
------
```
go build rcu.go
```
Run
------
```
-> % ./rcu -url http://example.com
LG Emul RCU v0.1
> Opening "http://andrej-andb.ru"... done.
> Connecting to Remocon... done.
< ?
> Commands:
  3               - Key 3 (303)
  vol_up          - Key vol_up (401)
  .....
< down
> Key down                                                                                                                                                                                                                                   
< exit                                                                                                                                                                                                                                       
> exit 
```
