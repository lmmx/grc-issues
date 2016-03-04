#!/usr/bin/env bash

SpeciesArray=(chicken human mouse zebrafish)
TopDir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && dirname "$(pwd)")
XmlParser="$TopDir/src/processGRCissues/processGRCissues"
SummaryCreator="$TopDir/src/create_summary_table.sh"

# nullglob is bash shell only
shopt -s nullglob # to loop over *.xml

for species in ${SpeciesArray[*]}; do
	IssueXmlDir="$TopDir/issues/$species"
	OutputTable="$TopDir/tables/$species.tsv"
	SummaryTable="$TopDir/tables/$species"_summary.tsv
	$XmlParser -header-only > $OutputTable
	for xmlfile in $IssueXmlDir/*.xml; do
		$XmlParser -no-header "$xmlfile"
	done >> $OutputTable
	source $SummaryCreator $OutputTable > $SummaryTable
done
