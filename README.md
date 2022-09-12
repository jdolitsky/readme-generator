# readme-generator

A README generator for Chainguard Images

## Usage

```
go run main.go \
  -name nginx \
  -location distroless.dev/nginx \
  -exclude-tags 1.20.2,1.20,1.23.0 \
  > README-GENERATED.md
```
