#!/usr/bin/env bash

SpeciesArray=(chicken human mouse zebrafish)
TopDir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && dirname "$(pwd)")
XmlParser="$TopDir/src/processGRCissues/processGRCissues"
SummaryWriter="$TopDir/src/create_summary_table.sh"
LocationWriter="$TopDir/src/create_longform_location_table.sh"

# nullglob is bash shell only
shopt -s nullglob # to loop over *.xml

for species in ${SpeciesArray[*]}; do
	IssueXmlDir="$TopDir/issues/$species"
	OutputTable="$TopDir/tables/$species.tsv"
	SummaryTable="$TopDir/tables/$species"_summary.tsv
	LocationTable="$TopDir/tables/$species"_locations.tsv
	$XmlParser -header-only > $OutputTable
	for xmlfile in $IssueXmlDir/*.xml; do
		$XmlParser -no-header "$xmlfile"
	done >> $OutputTable
	source $SummaryWriter $OutputTable > $SummaryTable
	source $LocationWriter $OutputTable > $LocationTable
done
