# Package protocol

## Services

{{range $f := .ProtoFile -}}
{{ if .Service -}}
{{range $s := .Service -}}
### {{.Name}}

File: {{ $f.Name }}

| Method Name | Request Type | Response Type |
| ----------- | ------------ | ------------- |
{{range .Method -}}
| {{.Name}} | {{.InputType}} | {{.OutputType}} |
{{end}}

{{end}}
{{end}}
{{end}}

## Messages

{{range $f := .ProtoFile -}}
{{ if .MessageType -}}
{{range $m := .MessageType -}}
### {{.Name}}

File: {{ $f.Name }}

| Name | Type | Options |
| ----------- | ------------ | ------------- |
{{range .Field -}}
| {{.Name}} | {{.Type}} | {{.Options}} |
{{end}}

{{end}}
{{end}}
{{end}}

