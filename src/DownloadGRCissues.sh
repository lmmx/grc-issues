#!/usr/bin/env bash

FtpRootDir="ftp://ftp.ncbi.nlm.nih.gov/pub/grc/"
FtpDirSuffix="/GRC/Issue_Mapping/"
SpeciesArray=(chicken human mouse zebrafish)
IssuesTopDir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && dirname "$(pwd)")"/issues"

for species in ${SpeciesArray[*]}; do
	IssueDirURL="$FtpRootDir$species$FtpDirSuffix"
	SpeciesDir="$IssuesTopDir/$species/"

	>&2 echo "Downloading from $IssueDirURL into $SpeciesDir"

	if [ ! -d $SpeciesDir ]; then
		>&2 echo "Creating $species directory under $IssuesTopDir"
		mkdir -p "$SpeciesDir"
	fi

	# wget or curl into this directory
	# wget $IssueDirURL chr*.xml $SpeciesDir
	wget -q -N -nd -r -l1 --no-parent -A "chr*.xml" $IssueDirURL -P $SpeciesDir
done
