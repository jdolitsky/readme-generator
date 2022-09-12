# {{.Name}}

## Image location

```
{{.Location}}
```

## Supported tags

| Tag | Digest | Signature |
| --- | ------ | --------- |
{{- range $tag := .Tags}}
| {{formatTagAliases $tag.Aliases}} | `{{$tag.Digest}}` | [View Rekor entry]({{$tag.RekorURL}}) |
{{- end -}}
