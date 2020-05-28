#!/usr/bin/env bash

# Copyright 2020 MongoDB Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


set -Eeou pipefail

VERSION=$(git describe --abbrev=0 | cut -d "v" -f 2)
FILENAME=mongocli_"${VERSION}"_linux_x86_64
if [[ "${unstable-}" == "-unstable" ]]; then
    FILENAME="mongocli_next_linux_x86_64"
fi

cd dist

mkdir yum apt

# we could generate a similar name with goreleaser but we want to keep the vars evg compatible to use later
cp "$FILENAME.deb" apt/
mv "apt/$FILENAME.deb" "apt/mongodb-cli${unstable-}_${VERSION}${latest_deb-}_amd64.deb"
cp "$FILENAME.rpm" yum/
mv "yum/$FILENAME.rpm" "yum/mongodb-cli${unstable-}-${VERSION}${latest_rpm-}.x86_64.rpm"
