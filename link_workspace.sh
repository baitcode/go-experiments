#!/usr/bin/env bash

WORKSPACE=vchain
PROJECT_PATH=github.com/vchain-dev/go-experiments
VG_ENV=~/.virtualgo/$WORKSPACE/src/$PROJECT_PATH

mkdir -p $VG_ENV
bindfs . $VG_ENV

cd $VG_ENV