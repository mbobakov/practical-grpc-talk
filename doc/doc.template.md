# {{.Package}} Package protocol

{{- define "Comments" -}}
	{{- $l := location . -}}
	{{- if $l -}}
		{{- if or $l.LeadingComments $l.TrailingComments -}}
			{{- with $l.LeadingComments -}}{{- . -}}{{end}}
			{{- with $l.TrailingComments -}}{{- . -}}{{end}}
		{{- end -}}
	{{- end -}}
{{- end -}}



## Services
{{ if .Service -}}
{{ range $s := .Service -}}
### {{.Name}}

Info:  {{ template "Comments" . }}

| Method Name | Request Type | Response Type | Comments |
| ----------- | ------------ | ------------- | ------- |
{{range .Method }}
| {{.Name}} | {{.InputType}} | {{.OutputType}} | {{- template "Comments" . -}}
{{end}}
{{end}}
{{end}}

## Messages

{{ if .MessageType -}}
{{range $m := .MessageType }}
### {{.Name}}
Info: {{ template "Comments" . }}

| Name | Type | Comments|
| ----------- | ------------ | ---------- |
{{range .Field -}}
| {{.Name}} | {{- fieldType . -}} | {{- template "Comments" . -}} |
{{end}}
{{end}}
{{end}}
