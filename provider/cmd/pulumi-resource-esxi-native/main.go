// Copyright 2016-2020, Pulumi Corporation.
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

//go:generate go run ./generate.go
package main

import (
	"flag"
	"fmt"

	"github.com/petrosight/pulumi-esxi-native/provider/pkg/provider"
    "github.com/petrosight/pulumi-esxi-native/provider/pkg/version"

	_ "embed"
)

var providerName = "esxi-native"

//go:embed schema-embed.json
var pulumiSchema []byte

func main() {
	showVersion := flag.Bool("version", false, "display version number and exit")
	flag.Parse()

	if *showVersion {
		fmt.Fprintf(flag.CommandLine.Output(), "Version v%s\n", version.Version)
		return
	}
	provider.Serve(providerName, version.Version, pulumiSchema)
}
