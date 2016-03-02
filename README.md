# grc-issues

Tabulated mappings of human genome assembly issues to their location (and all other associated information from NCBI).

See also: [Ensembl assembly exceptions](https://github.com/lmmx/ensembl-assembly-exceptions), tabulated mappings of _H. sapiens_ and _M. musculus_ genome 'assembly exceptions' to their locations in core Ensembl database releases.

## Requirements

Should not require anything other than a `bash` shell and internet connection to download updates. Processing and download should not take long. 

The XML parser is written in Go 1.6, and provided as an executable binary (i.e. does not require you to install Go). To build from source, run `go build` in the `src/processGRCissues` directory to regenerate the executable.

## Usage

Run the helper script to execute downloads and parse the received XML into tab-separated output files under the `tables` directory. The timestamping flag to `wget` (`-N`) ensures it only redownloads the file if the server's version is newer than that on disk.

```sh
./update_tables.sh
```

* To download the Genome Reference Consortium issues XML from the main directory, run `./src/DownloadGRCissues.sh`
  * FTP area for all species: ftp://ftp.ncbi.nlm.nih.gov/pub/grc/
* To generate tables for chicken, human, mouse, and zebrafish from the main directory run `./src/ParseGRCissues.sh`.

## Table summary

As of 1st March 2016:

| Species   | No. issues |
|-----------|------------|
| chicken   |         6  |
| human     |      2230  |
| mouse     |      1050  |
| zebrafish |      4401  |

## Modifications made

Tables contain data 'as-is' from NCBI, except:
* chromosome names have the leading 'chr' prefix removed where present. chrNA therefore became NA
* empty fields relating to patch metadata were filled with 'na' (in-keeping with the other NA values), since it was a nillable field in the NCBI XML schema.
* Free text fields (`<summary>`, `<description>`, `<status_text>`, and `<resolution_text>` tags) were parsed to Unicode, making them readable, but left quoted, with newlines unescaped (`like\nthis`). Other columns are not quoted (and do not contain HTML-encoded characters).
