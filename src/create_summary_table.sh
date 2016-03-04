#!/usr/bin/env bash

IssueTable=$1

>&2 echo ": : : Writing `basename $IssueTable .tsv` summary to a separate table."

cut -d $'\t' -f 2,3,19 $IssueTable
