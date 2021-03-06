#!/usr/bin/env bash
# Copyright 2017 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -e

DEP_ROOT=$(git rev-parse --show-toplevel)

mkdir -p "${DEP_ROOT}/release"

cd "${DEP_ROOT}/release/"
tar -jcvf typpes.tar.bz2 $HOME/gopath/src/github.com/rancher/types/
