# Changelog

{{- range $commit := .}}
{{- if ne (len $commit.Tags) 0}}

## {{index $commit.Tags 0}} ({{$commit.Date.Format "2006/01/02"}})
{{- end}}
- `{{$commit.Hash | shortHash}}` {{$commit.Title}}
{{- end}}
