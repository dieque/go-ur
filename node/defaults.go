// Copyright 2016 The go-ur Authors
// This file is part of the go-ur library.
//
// The go-ur library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ur library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ur library. If not, see <http://www.gnu.org/licenses/>.

package node

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

const (
	DefaultIPCSocket = "gur.ipc"  // Default (relative) name of the IPC RPC socket
	DefaultHTTPHost  = "localhost" // Default host interface for the HTTP RPC server
	DefaultHTTPPort  = 9595        // Default TCP port for the HTTP RPC server
	DefaultWSHost    = "localhost" // Default host interface for the websocket RPC server
	DefaultWSPort    = 9596        // Default TCP port for the websocket RPC server
)

// DefaultDataDir is the default data directory to use for the databases and other
// persistence requirements.
func DefaultDataDir() string {
	// Try to place the data folder in the user's home dir
	home := homeDir()
	if home != "" {
		if runtime.GOOS == "darwin" {
			return filepath.Join(home, "Library", "UR")
		} else if runtime.GOOS == "windows" {
			return filepath.Join(home, "AppData", "Roaming", "UR")
		} else {
			return filepath.Join(home, ".ur")
		}
	}
	// As we cannot guess a stable location, return empty and handle later
	return ""
}

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}
