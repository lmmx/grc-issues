#!/usr/bin/env bash

TopDir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
ScriptDir=$TopDir/src

source $ScriptDir/DownloadGRCissues.sh
source $ScriptDir/ParseGRCissues.sh

echo "Tables written into $TopDir/tables"
