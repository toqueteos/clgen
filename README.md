# clgen

A fast changelog generator using git as source.

Generates two files: `changelog.md` and `changelog.html`.

## Installation (development)

```
$ go get github.com/toqueteos/clgen/cmd/clgen
```

## Usage

```
$ clgen --help
Usage of /home/toqueteos/code/go/bin/clgen:
  -html-out string
        HTML output file (default "changelog.html")
  -md-out string
        Markdown output file (default "changelog.md")
  -ref string
        ref to start going back, use HEAD~100 for last 100 commits (default "HEAD")
  -title string
        HTML document title (default "Changelog")
  -tpl string
        template input file (default "changelog.tpl")
```

Example usage:

```
$ cd $YOUR_GIT_PROJECT
$ clgen -tpl changelog.tpl
Done! Elapsed 78.9971ms
$ ls -lh changelog*
-rw-r--r-- 1 toq 197121  19K nov.  6 13:03 changelog.html
-rw-r--r-- 1 toq 197121 7,8K nov.  6 13:03 changelog.md
```

## Samples

**template.tpl**

```
# Changelog

{{- range $commit := .}}
{{- if ne (len $commit.Tags) 0}}

## {{index $commit.Tags 0}} ({{$commit.Date.Format "2006/01/02"}})
{{- end}}
- {{$commit.Hash | shortHash}} {{$commit.Title}}
{{- end}}
```
