#!/bin/bash


# Copyright 2018 ZTE Corporation. All rights reserved.
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# http://www.apache.org/licenses/LICENSE-2.0
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


set -e
OPERATION=$1
IFDOCOMPARE=$2
MODULENAME="knitter-agent"

echo "before GOROOT=$GOROOT    GOPATH=$GOPATH  GO=$GO  KNITTERPATH=${KNITTERPATH}"
FILEDIR=$(cd "$(dirname $0)";pwd)
source ${FILEDIR}/common.sh
echo "after GOROOT=$GOROOT    GOPATH=$GOPATH  GO=$GO   KNITTERPATH=${KNITTERPATH}"

if [ ${OPERATION} = "build" ];then
	############## build begin ##############
	echo "============== building ${MODULENAME} =============="
	cd  ${KNITTERPATH}/knitter-agent
	${GO}/go clean
	${GO}/go build -ldflags "-X ${BUILDPKG}/pkg/version.moduleName=${MODULENAME} -X ${BUILDPKG}/pkg/version.verType=${VERTYPE} -X ${BUILDPKG}/pkg/version.versionInfo=${BRANCH_VERSION} -X ${BUILDPKG}/pkg/version.gitHash=${GITHASH} -X ${BUILDPKG}/pkg/version.buildTime=${BUILDTIME}"

	if [ -f ${KNITTERPATH}/knitter-agent/knitter-agent ];then
		echo "++++ build ${MODULENAME} success"
		exit 0
	else
		echo "++++ build_agent:: error knitter-agent is not existing"
		exit 1
	fi

else
	############ other begin ##############
	echo "============== other operation =============="
	exit 1

fi
