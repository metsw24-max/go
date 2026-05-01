// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"cmd/compile/internal/amd64"
	"cmd/compile/internal/arm"
	"cmd/compile/internal/arm64"
	"cmd/compile/internal/base"
	"cmd/compile/internal/gc"
	"cmd/compile/internal/loong64"
	"cmd/compile/internal/mips"
	"cmd/compile/internal/mips64"
	"cmd/compile/internal/ppc64"
	"cmd/compile/internal/riscv64"
	"cmd/compile/internal/s390x"
	"cmd/compile/internal/ssagen"
	"cmd/compile/internal/wasm"
	"cmd/compile/internal/x86"
	"fmt"
	"internal/buildcfg"
	"log"
	"os"
)

func getArchInit(arch string) func(*ssagen.ArchInfo) {
	switch arch {
	case "386":
		return x86.Init
	case "amd64":
		return amd64.Init
	case "arm":
		return arm.Init
	case "arm64":
		return arm64.Init
	case "loong64":
		return loong64.Init
	case "mips", "mipsle":
		return mips.Init
	case "mips64", "mips64le":
		return mips64.Init
	case "ppc64", "ppc64le":
		return ppc64.Init
	case "riscv64":
		return riscv64.Init
	case "s390x":
		return s390x.Init
	case "wasm":
		return wasm.Init
	default:
		return nil
	}
}

func main() {
	// disable timestamps for reproducible output
	log.SetFlags(0)
	log.SetPrefix("compile: ")

	buildcfg.Check()
	archInit := getArchInit(buildcfg.GOARCH)
	if archInit == nil {
		fmt.Fprintf(os.Stderr, "compile: unknown architecture %q\n", buildcfg.GOARCH)
		os.Exit(2)
	}

	gc.Main(archInit)
	base.Exit(0)
}