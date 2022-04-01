# arweaveutil
An Arweave util, can transfer AR/PST, check balance, generate Arweave keyfile.

# Usage Example
## Transfer AR
Transfer 1 AR to address dRFuVE-s6-TgmykU4Zqn246AR2PIsf3HhBhZ0t5-WXE:
```shell
$ arweaveutil --wallet keyfile.json transfer dRFuVE-s6-TgmykU4Zqn246AR2PIsf3HhBhZ0t5-WXE 1
```

## Transfer PST
Transfer 1 ARDRIVE to address dRFuVE-s6-TgmykU4Zqn246AR2PIsf3HhBhZ0t5-WXE:
```shell
$ arweaveutil --wallet keyfile.json transfer-pst ARDRIVE dRFuVE-s6-TgmykU4Zqn246AR2PIsf3HhBhZ0t5-WXE 1
```

## Check Balance
Check balance of an address:
```shell
$ arweaveutil balance dRFuVE-s6-TgmykU4Zqn246AR2PIsf3HhBhZ0t5-WXE
addr dRFuVE-s6-TgmykU4Zqn246AR2PIsf3HhBhZ0t5-WXE, balance 10200042.49 AR
```

## Generate Arweave keyfile
Generate Arweave keyfile:
```shell
$ arweaveutil gen-wallet
keyfile arweave-keyfile-nWUrW0o2wfN8XeuqiSnHGIbOzl0iSjv2SbweGCsEg1I.json saved
```

# Install
```shell
GO111MODULE=on go install github.com/10gic/arweaveutil@latest
```
