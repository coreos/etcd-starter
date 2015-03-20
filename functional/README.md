
etcd starter functional tests
=====

This functional test suite deploys a etcd cluster using processes, and asserts etcd-starter is functioning properly.

Usage
-----

Set environment variables point to the respective binaries that are used to drive the actual tests:

```
$ export ETCD_V1_BIN=/path/to/v1_etcd
$ export ETCD_V2_BIN=/path/to/v2_etcd
$ export ETCDCTL_BIN=/path/to/etcdctl
$ export ETCD_STARTER_BIN=/path/to/etcd-starter
```

Then the tests can be run:

```
$ go test github.com/coreos/etcd/migrate/functional
```
