// Copyright 2016 The go-ur Authors
// This file is part of go-ur.
//
// go-ur is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ur is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ur. If not, see <http://www.gnu.org/licenses/>.

// Command bzzhash computes a swarm tree hash.
package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/ur-technology/go-ur/swarm/storage"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if len(os.Args) < 2 {
		fmt.Println("Usage: bzzhash <file name>")
		os.Exit(0)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file " + os.Args[1])
		os.Exit(1)
	}

	stat, _ := f.Stat()
	chunker := storage.NewTreeChunker(storage.NewChunkerParams())
	key, err := chunker.Split(f, stat.Size(), nil, nil, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	} else {
		fmt.Printf("%v\n", key)
	}
}
