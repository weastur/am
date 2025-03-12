#!/bin/sh

set -e

rm -rf build/completions
mkdir build/completions

for sh in bash zsh fish; do
	./bin/am completion "$sh" >"build/completions/am.$sh"
done
