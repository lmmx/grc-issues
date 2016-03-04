#!/usr/bin/env bash

TopDir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && dirname "$(pwd)")
LocationParser="$TopDir/src/parse_shortform_locations.r"
IssueTable=$1

>&2 echo ": : : Writing `basename $IssueTable .tsv` longform issue location details to a separate table."

cut -d $'\t' -f 2,3,22 $IssueTable | $LocationParser
