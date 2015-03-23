// Copyright 2015 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main is a simple wrapper of the real etcd entrypoint package
// (located at github.com/coreos/etcd/etcdmain) to ensure that etcd is still
// "go getable"; e.g. `go get github.com/coreos/etcd` works as expected and
// builds a binary in $GOBIN/etcd
//
// This package should NOT be extended or modified in any way; to modify the
// etcd binary, work in the `github.com/coreos/etcd/etcdmain` package.
//

package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/coreos/etcd-starter/starter"
)

const defaultInternalBinaryDir = "/usr/libexec/etcd/internal_versions/"

const usage = `usage: etcd-starter [flags]
start v0.4 or v2.0 etcd server based on the layout of data directory and the content inside wal and snapshot

etcd-starter --version
show the version of etcd-starter and exit.

etcd-starter -h | --help
show the help information about etcd

etcd-starter [etcd flags]
start v0.4 or v2.0 etcd using the given flags

Please check etcd documents for more information about etcd flag usage.
`

var (
	showVersion bool
)

func main() {
	fs := flag.NewFlagSet("etcd-starter", flag.ContinueOnError)
	fs.BoolVar(&showVersion, "version", false, "print version and exit")
	fs.Usage = func() {
		fmt.Println(usage)
		os.Exit(0)
	}

	mainArgs, etcdArgs := filterArgs()
	fs.Parse(mainArgs)

	if showVersion {
		fmt.Println("etcd-starter version", version)
		os.Exit(0)
	}

	dir := os.Getenv("ETCD_INTERNAL_BINARY_DIR")
	if dir == "" {
		dir = defaultInternalBinaryDir
	}
	starter.StartDesiredVersion(dir, etcdArgs)
}

func filterArgs() ([]string, []string) {
	mainArgs := make([]string, 0)
	etcdArgs := make([]string, 0)
	pattern := regexp.MustCompile("--version|-version|-h|--help|-help")
	for _, arg := range os.Args[1:] {
		if pattern.MatchString(arg) {
			mainArgs = append(mainArgs, arg)
			continue
		}
		etcdArgs = append(etcdArgs, arg)
	}
	return mainArgs, etcdArgs
}
