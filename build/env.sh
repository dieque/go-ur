#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
urdir="$workspace/src/github.com/ur-technology"
if [ ! -L "$urdir/go-ur" ]; then
    mkdir -p "$urdir"
    cd "$urdir"
    ln -s ../../../../../. go-ur
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
GO15VENDOREXPERIMENT=1
export GOPATH GO15VENDOREXPERIMENT

# Run the command inside the workspace.
cd "$urdir/go-ur"
PWD="$urdir/go-ur"

# Launch the arguments with the configured environment.
exec "$@"
