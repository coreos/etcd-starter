
etcd-starter
=====

etcd-starter is the starter code to faciliate etcd upgrade in CoreOS image. The main functionality is:

1. do binary selection between etcd v0.4 and etcd v2.0 based on the layout of data directory and the content inside wal and snapshot
2. modify flag for standby-mode v0.4 etcd, which helps it run on proxy mode

Installation
------------

```
mkdir $GOPATH/src/github.com/coreos
cd $GOPATH/src/github.com/coreos
git clone git@github.com:coreos/etcd-starter.git
cd etcd-starter
go build .
```

Usage
-----

Set `ETCD_INTERNAL_BINARY_DIR` to the internal binary directory, which should
be organized in this way:

```
ETCD_INTERNAL_BINARY_DIR
├── 1
    ├── etcd (points to v0.4 etcd binary)
├── 2
    ├── etcd (points to v2.0 etcd binary)
```

The default value for `ETCD_INTERNAL_BINARY_DIR` is "/usr/libexec/etcd/internal_versions/".

Start etcd-starter just like etcd. Here is an example:

```
$ etcd-starter --name default --data-dir default.etcd
```

Upgrade Example
---------------

Assume that you have a data directory `default.etcd` which is used by v0.4 etcd.

You can start it using v0.4 etcd as before:

```
$ etcd-starter --name default --data-dir default.etcd
```

When you want to upgrade it to v2.0, trigger it:

```
$ etcdctl upgrade --peer-urls http://127.0.0.1:7001
```

etcd will exit after 10 seconds if everything goes well. Restart it using the same command line:

```
$ etcd-starter --name default --data-dir default.etcd
```

It will run using v2.0 etcd now.
