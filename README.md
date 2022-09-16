# readme-generator

A README generator for Chainguard Images

## Usage

### Github Actions

```
- uses: distroless/readme-generator@main
  with:
    repo: https://github.com/distroless/nginx
    name: nginx
    location: distroless.dev/nginx
    description: "A minimal nginx base image rebuilt every night from source."
    exclude-tags: 1.20.2-r0,1.20.2,1.20,1.23.0
    output-path: README-GENERATED.md
    push-to-repo: true
    push-to-repo-message: "Regenerate README.md"
```

Note: if `push-to-repo: true` is set, this action will attempt to push back to the repo.
This will require setting the following permissions on the job context:

```
    permissions:
      id-token: write
      contents: write
```

### Locally

```
go run main.go \
  -repo https://github.com/distroless/nginx \
  -name nginx \
  -location distroless.dev/nginx \
  -description "A minimal nginx base image rebuilt every night from source." \
  -exclude-tags 1.20.2-r0,1.20.2,1.20,1.23.0 \
  > README-GENERATED.md
```

## Caveats

This tool assumes the following:

- Repos and images are publically accessible
- The image has a `latest` tag
- The image tags all point to a multiarch OCI index
- The repo is hosted on GitHub and the default branch is `main`
- `cosign` and `jq` are installed (shells out to these)
- A GitHub actions workflow file exists called `release.yaml`
- The repo contains a file called `USAGE.md` to populate the "Usage" section
- The repo contains a file called `melange.yaml` to determine if it uses melange
