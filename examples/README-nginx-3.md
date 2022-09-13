# nginx

[![CI status](https://github.com/distroless/nginx/actions/workflows/release.yaml/badge.svg)](https://github.com/distroless/nginx/actions/workflows/release.yaml)

A minimal nginx base image rebuilt every night from source.

## Get It!

The image is available on `distroless.dev`:

```
docker pull distroless.dev/nginx:latest
```

## Supported tags

| Tag | Digest | Arch | Signature |
| --- | ------ | ---- | --------- |
| `1.22` `1.22.0` `1.22.0-r0` `stable` | `sha256:18a14206c1b29466cab0f9b31bb98f17b6a99729c412d471c4728f50c4d40717` | `amd64` `arm64` `armv7` | [View Rekor entry](https://rekor.tlog.dev/?hash=sha256:18a14206c1b29466cab0f9b31bb98f17b6a99729c412d471c4728f50c4d40717) |
| `1` `1.23` `1.23.1` `1.23.1-r0` `latest` `mainline` | `sha256:58765ffd6b8d01a5159b143b6a38efa6fb0a06f4f0e06659588aedd45a2c72cd` | `amd64` `arm64` `armv7` | [View Rekor entry](https://rekor.tlog.dev/?hash=sha256:58765ffd6b8d01a5159b143b6a38efa6fb0a06f4f0e06659588aedd45a2c72cd) |


## Usage

To try out the image, run:

```
docker run -p 8080:80 distroless.dev/nginx
```

If you navigate to `localhost:8080`, you should see the nginx welcome page.

To run an example Hello World app, navigate to the root of this repo and run:

```
docker run -v $(pwd)/examples/hello-world/site-content:/var/lib/nginx/html -p 8080:80 distroless.dev/nginx
```

If you navigate to `localhost:8080`, you should see `Hello World from Nginx Distroless!`.


## Signing

All distroless images are signed using [Sigstore](https://sigstore.dev)!

<details>
<br/>
To verify an image, download <a href="https://github.com/sigstore/cosign">cosign</a> and run:

```
COSIGN_EXPERIMENTAL=1 cosign verify distroless.dev/nginx:latest | jq
```

```
TODO
```

You can verify that the image was built in Github Actions in this repository from the `Issuer` and `Subject` fields.
</details>

## Build

This image is built with [melange](https://github.com/chainguard-dev/melange) and [apko](https://github.com/chainguard-dev/apko).

