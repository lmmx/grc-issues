#!/usr/bin/env bash
cut -d $'\t' -f 2,3,19 $1 | grep -vP '\tna$|\tNone$'
