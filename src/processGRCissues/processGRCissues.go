package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var l *log.Logger = log.New(os.Stderr, "", 0)
var xml_arg string

type Listing struct {
	IssueList []Issue `xml:"issue"`
}

type Issue struct {
	IssueType      string         `xml:"type"`
	IssueID        string         `xml:"key"`
	Chromosome     string         `xml:"assignedChr"`
	Accession1     string         `xml:"accession1"`
	Accession2     string         `xml:"accession2"`
	ReportType     string         `xml:"reportType"`
	Summary        string         `xml:"summary"`
	Status         string         `xml:"status"`
	StatusText     string         `xml:"status_text"` // &#xa; line feeds need escaping
	Description    string         `xml:"description"` // &#xa; line feeds need escaping
	ExperimentType string         `xml:"experiment_type"`
	Update         string         `xml:"update"`
	ExtInfoType    string         `xml:"external_info_type"`
	Resolution     string         `xml:"resolution"`
	ResolutionText string         `xml:"resolution_text"` // &#xa; line feeds need escaping
	AltPatchType   string         `xml:"AltPatchType"`
	FixedInVer     string         `xml:"fixVersion"`
	AffectedVer    string         `xml:"affectVersion"`
	PatchInfo      PatchInfo      `xml:"patchInfo"`
	Location       []PositionInfo `xml:"location>position"`
}

type PatchInfo struct {
	// Complex type from patchInfo
	GenBankID  string `xml:"gb_acc,attr"`
	RefSeqID   string `xml:"ref_acc,attr"`
	RegionName string `xml:"region_name,attr"`
}

type PositionInfo struct {
	SuccessfullyMapped string         `xml:"mapStatus"`
	MappedSeqInfo      MappedSeqInfo  `xml:"mapSequence"`
	ChrStart           string         `xml:"start"`
	ChrEnd             string         `xml:"stop"`
	MappingQuality     MappingQuality `xml:"quality"`
	AssemblyName       string         `xml:"name,attr"`
	GenBankAssemblyAcc string         `xml:"gb_asm_acc,attr"`
	RefSeqAssemblyAcc  string         `xml:"ref_asm_acc,attr"`
	AssemblyStatus     string         `xml:"asm_status,attr"`
}

type MappedSeqInfo struct {
	GenBankID    string `xml:"gb_acc,attr"`
	RefSeqID     string `xml:"ref_acc,atr"`
	SequenceType string `xml:"type,attr"`
}

type MappingQuality struct {
	MappedVersion    []MappedVersion `xml:"version_mapped"`
	Accession1Method string          `xml:"method_acc1"`
	Accession2Method string          `xml:"method_acc2"`
}

type MappedVersion struct {
	MappedVersionAccession string `xml:"acc,attr"`
	MappedVersionNumber    string `xml:",chardata"`
}

/*
func (e Episode) String() string {
	return fmt.Sprintf("S%02dE%02d - %s - %s", e.SeasonNumber, e.EpisodeNumber, e.EpisodeName, e.FirstAired)
}
*/

func PrintHeader() {
	fmt.Println(strings.Join([]string{
		"Issue_Type",
		"Issue_ID",
		"Chromosome",
		"Accession_1",
		"Accession_2",
		"Report_Type",
		"Summary",
		"Status",
		"Status_Text",
		"Description",
		"Experiment_Type",
		"Update",
		"External_Information_Type",
		"Resolution_Summary",
		"Resolution_Information",
		"Patch_Type",
		"Fixed_in_Version",
		"Affected_Version",
		"Patch_GenBank_Accession",
		"Patch_RefSeq_Accession",
		"Patch_Region_Name"},
		"\t"))
}

func main() {
	headerOnlyPtr := flag.Bool("header-only", false, "Print table header line and exit")
	noHeaderPtr := flag.Bool("no-header", false, "Will not print table header line with the table if used")

	flag.Parse()

	if *headerOnlyPtr && *noHeaderPtr {
		l.Println("The '-header-only' and '-no-header' flags cannot be supplied together")
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		xml_arg = os.Args[2]
	} else {
		xml_arg = os.Args[1]
	}

	if *headerOnlyPtr {
		PrintHeader()
		return
	}

	// Conditionally provide N/A values for nillable fields, or to trim "chr"
	var patch_type, fix_ver, aff_ver, gb_acc, rs_acc, region_name, chromosome string

	// Do this for all chromosomes and automate the output of 2 tables...
	xmlFile, err := os.Open(xml_arg)
	if err != nil {
		l.Println("Error reading XML: ", err)
		return
	}
	defer xmlFile.Close()

	// b := []byte(xmlstr)
	b, _ := ioutil.ReadAll(xmlFile)

	var l Listing
	xml.Unmarshal(b, &l)

	if !*noHeaderPtr {
		PrintHeader()
	}

	for _, issue := range l.IssueList {

		chromosome = strings.TrimPrefix(issue.Chromosome, "chr")

		fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t",
			issue.IssueType,
			issue.IssueID,
			chromosome,
			issue.Accession1,
			issue.Accession2,
			issue.ReportType,
			issue.Summary,
			issue.Status)

		fmt.Printf("%#v\t%#v\t",
			issue.StatusText,
			issue.Description) // escape and quote

		fmt.Printf("%s\t%s\t%s\t%s\t",
			issue.ExperimentType,
			issue.Update,
			issue.ExtInfoType,
			issue.Resolution)

		fmt.Printf("%#v\t", issue.ResolutionText) // escape and quote

		if issue.AltPatchType != "" {
			patch_type = issue.AltPatchType
		} else {
			patch_type = "na"
		}
		if issue.FixedInVer != "" {
			fix_ver = issue.FixedInVer
		} else {
			fix_ver = "na"
		}
		if issue.AffectedVer != "" {
			aff_ver = issue.AffectedVer
		} else {
			aff_ver = "na"
		}
		if issue.PatchInfo.GenBankID != "" {
			gb_acc = issue.PatchInfo.GenBankID
		} else {
			gb_acc = "na"
		}

		if issue.PatchInfo.RefSeqID != "" {
			rs_acc = issue.PatchInfo.RefSeqID
		} else {
			rs_acc = "na"
		}
		if issue.PatchInfo.RegionName != "" {
			region_name = issue.PatchInfo.RegionName
		} else {
			region_name = "na"
		}

		fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\n",
			patch_type,
			fix_ver,
			aff_ver,
			gb_acc,
			rs_acc,
			region_name)
	}
}
