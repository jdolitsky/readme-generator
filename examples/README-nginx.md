# nginx

[![go](https://github.com/chainguard-dev/hello-melange-apko/actions/workflows/go.yml/badge.svg)](https://github.com/chainguard-dev/hello-melange-apko/actions/workflows/go.yml)

A minimal nginx base image rebuilt every night from source.

## Get It!

The image is available on `cgr.dev`:

```
docker pull cgr.dev/chainguard/nginx:latest
```

## Supported tags

| Tag | Digest | Arch | Signature |
| --- | ------ | -- | --------- |
| `1` `1.23` `1.23.1` `1.23.1-r0` `latest` `mainline` | `sha256:d08d864569e20105bed1d4f58b852ea3d810e7d26ec0280011dcae1135421f3f` | `amd64` `arm64` `armv7` | [View Rekor entry](https://rekor.tlog.dev/?hash=sha256:d08d864569e20105bed1d4f58b852ea3d810e7d26ec0280011dcae1135421f3f) |
| `1.22` `1.22.0` `1.22.0-r0` `stable` | `sha256:2b428426605629d3110dddc75de095ac6068c138b08a430baac9b8637633afb8` | `amd64` `arm64` `armv7` | [View Rekor entry](https://rekor.tlog.dev/?hash=sha256:2b428426605629d3110dddc75de095ac6068c138b08a430baac9b8637633afb8) |

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
To verify an image, download [cosign](https://github.com/sigstore/cosign) and run:

```
COSIGN_EXPERIMENTAL=1 cosign verify distroless.dev/nginx | jq

Verification for distroless.dev/nginx:latest --
The following checks were performed on each of these signatures:
  - The cosign claims were validated
  - Existence of the claims in the transparency log was verified offline
  - Any certificates were verified against the Fulcio roots.
[
  {
    "critical": {
      "identity": {
        "docker-reference": "ghcr.io/distroless/nginx"
      },
      "image": {
        "docker-manifest-digest": "sha256:3b28db71687f52741598f4f68d2e4bea8ee86db57d7394337118316d1f4c8b9f"
      },
      "type": "cosign container image signature"
    },
    "optional": {
      "Issuer": "https://token.actions.githubusercontent.com",
      "Subject": "https://github.com/distroless/nginx/.github/workflows/release.yaml@refs/heads/main",
      "run_attempt": "1",
      "run_id": "2626578822"
      ...
    }
  }
]
```

You can verify that the image was built in Github Actions in this repository from the `Issuer` and `Subject` fields.
</details>


## Build

This image is built with [apko](https://github.com/chainguard-dev/apko) and
[melange](https://github.com/chainguard-dev/melange) tooling.
