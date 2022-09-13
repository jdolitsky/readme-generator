# nginx

<!---
Note: Do NOT edit directly, this file was generated using https://github.com/distroless/readme-generator
-->

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
| `1.22` `1.22.0` `1.22.0-r0` `latest` `stable` | `sha256:2406be6ffd046c13bfd4452120be8149e9eb0610306a84ab3ce39edc8c50b4e4` | `amd64` `arm64` `armv7` | [View Rekor entry](https://rekor.tlog.dev/?hash=sha256:2406be6ffd046c13bfd4452120be8149e9eb0610306a84ab3ce39edc8c50b4e4) |
| `1` `1.23` `1.23.1` `1.23.1-r0` `mainline` | `sha256:8a05a673cfbeda4f9c16c243ad8a3a0309065c082bba8774df9bbb6da90119de` | `amd64` `arm64` `armv7` | [View Rekor entry](https://rekor.tlog.dev/?hash=sha256:8a05a673cfbeda4f9c16c243ad8a3a0309065c082bba8774df9bbb6da90119de) |


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

Output:
```
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
        "docker-manifest-digest": "sha256:2406be6ffd046c13bfd4452120be8149e9eb0610306a84ab3ce39edc8c50b4e4"
      },
      "type": "cosign container image signature"
    },
    "optional": {
      "1.3.6.1.4.1.57264.1.2": "push",
      "1.3.6.1.4.1.57264.1.3": "ca4381349d1ae0b5128db6d21339a8cb44576feb",
      "1.3.6.1.4.1.57264.1.4": "Create Release",
      "1.3.6.1.4.1.57264.1.5": "distroless/nginx",
      "1.3.6.1.4.1.57264.1.6": "refs/heads/main",
      "Bundle": {
        "SignedEntryTimestamp": "MEUCIA7+r+WeCfd0PQcerGcafaFV9Oi5EzF7Zn2q0d2UW54uAiEAy/W3OwJoRZczlGHYjUTai6XxjDSWwcgysh2HwaxBvu0=",
        "Payload": {
          "body": "eyJhcGlWZXJzaW9uIjoiMC4wLjEiLCJraW5kIjoiaGFzaGVkcmVrb3JkIiwic3BlYyI6eyJkYXRhIjp7Imhhc2giOnsiYWxnb3JpdGhtIjoic2hhMjU2IiwidmFsdWUiOiIwMTcyMTFmZTE4ZjBjOWUxMTk3NDc1YTJlMDVlZTRiN2I0ZWNkOWMwNzZhMTQ3MTRmNzkzNDUxMzg3Zjg1ZGY1In19LCJzaWduYXR1cmUiOnsiY29udGVudCI6Ik1FVUNJUUN2bkVtemR1akd5RTNjOTlJS3drM2djNGsyWndPYVpLY29aU01BVDdsNnFRSWdJMTgvZTRQYWlIalZXQTRKUThVRmpMR0QvYm1zYnpEdG1IVEJubjAwU1pvPSIsInB1YmxpY0tleSI6eyJjb250ZW50IjoiTFMwdExTMUNSVWRKVGlCRFJWSlVTVVpKUTBGVVJTMHRMUzB0Q2sxSlNVUnRWRU5EUVhnMlowRjNTVUpCWjBsVlVHMVZhRzVvWmxKUFdUSXJVbFJFZWtwV2JrNTFUV0Z2ZDJNNGQwTm5XVWxMYjFwSmVtb3dSVUYzVFhjS1RucEZWazFDVFVkQk1WVkZRMmhOVFdNeWJHNWpNMUoyWTIxVmRWcEhWakpOVWpSM1NFRlpSRlpSVVVSRmVGWjZZVmRrZW1SSE9YbGFVekZ3WW01U2JBcGpiVEZzV2tkc2FHUkhWWGRJYUdOT1RXcEpkMDlVUlhwTmFrVjZUWHBCTkZkb1kwNU5ha2wzVDFSRmVrMXFSVEJOZWtFMFYycEJRVTFHYTNkRmQxbElDa3R2V2tsNmFqQkRRVkZaU1V0dldrbDZhakJFUVZGalJGRm5RVVZ2ZVU5VlYyeHJUaXRZYTNwYU1UQXhhbkJpV2xCQlVYQXdlV1J0YkVSSGJGbHBUMHNLVkRrMVdrbHFVbVJZZDJ3NVdEWkdRMDlaUTFOSGVYWnRhbnB5VFcxNWNVRk5NRWRDZURaQkszaDFhakZKY2pCYWN6WlBRMEZxTUhkblowazFUVUUwUndwQk1WVmtSSGRGUWk5M1VVVkJkMGxJWjBSQlZFSm5UbFpJVTFWRlJFUkJTMEpuWjNKQ1owVkdRbEZqUkVGNlFXUkNaMDVXU0ZFMFJVWm5VVlV5Y0ZkekNtMXpUM0UyZEZRNWRFUlFVbU00YjBOUVZYWm5lbEZqZDBoM1dVUldVakJxUWtKbmQwWnZRVlV6T1ZCd2VqRlphMFZhWWpWeFRtcHdTMFpYYVhocE5Ga0tXa1E0ZDFsQldVUldVakJTUVZGSUwwSkdXWGRXU1ZwVFlVaFNNR05JVFRaTWVUbHVZVmhTYjJSWFNYVlpNamwwVERKU2NHTXpVbmxpTW5oc1l6Tk5kZ3BpYldSd1ltNW5ka3h0WkhCa1IyZ3hXV2s1TTJJelNuSmFiWGgyWkROTmRtTnRWbk5hVjBaNldsTTFOVmxYTVhOUlNFcHNXbTVOZG1GSFZtaGFTRTEyQ21KWFJuQmlha0UxUW1kdmNrSm5SVVZCV1U4dlRVRkZRa0pEZEc5a1NGSjNZM3B2ZGt3elVuWmhNbFoxVEcxR2FtUkhiSFppYmsxMVdqSnNNR0ZJVm1rS1pGaE9iR050VG5aaWJsSnNZbTVSZFZreU9YUk5Ra2xIUTJselIwRlJVVUpuTnpoM1FWRkpSVUpJUWpGak1tZDNUbWRaUzB0M1dVSkNRVWRFZG5wQlFncEJkMUZ2V1RKRk1FMTZaM2hOZWxFMVdrUkdhRnBVUW1sT1ZFVjVUMGRTYVU1dFVYbE5WRTE2VDFkRk5Ga3lTVEJPUkZVelRtMWFiRmxxUVdOQ1oyOXlDa0puUlVWQldVOHZUVUZGUlVKQk5VUmpiVlpvWkVkVloxVnRWbk5hVjBaNldsUkJaVUpuYjNKQ1owVkZRVmxQTDAxQlJVWkNRa0pyWVZoT01HTnRPWE1LV2xoT2Vrd3lOVzVoVnpVMFRVSXdSME5wYzBkQlVWRkNaemM0ZDBGUldVVkVNMHBzV201TmRtRkhWbWhhU0UxMllsZEdjR0pxUTBKcGQxbExTM2RaUWdwQ1FVaFhaVkZKUlVGblVqbENTSE5CWlZGQ00wRkJhR2RyZGtGdlZYWTViMUprU0ZKaGVXVkZia1ZXYmtkTGQxZFFZMDAwTUcwemJYWkRTVWRPYlRsNUNrRkJRVUpuZW1wR1MyUnpRVUZCVVVSQlJXZDNVbWRKYUVGS1ZWa3djazFSVlVWRmVVTlNUbGRRWldKalRXbHZVV1pVWVdkMGNYTk5VeXROTlcxclFXUUtkREJuYjBGcFJVRnJVV1ZPWkdoalJGWlFWSGh1YkRKS1MzaHBaRFYxY0ZGallXSXpTVGhhZFhsd1dsaGphVmhHYzJSWmQwTm5XVWxMYjFwSmVtb3dSUXBCZDAxRVlWRkJkMXBuU1hoQlRESkRObWwyZDFSYVUwazJVMFJKY3poNGMyWXZjM1F5V1VVek5qSlpUVEE1UW14NE9EWkhaU3ROUmxoRGJuVTBkbWRRQ2tJeU9WRlVTSHBZUm5CV2NEUlJTWGhCVUdGTFFrWk9TMloyZFRJMGJIZG9jV3BrY1dRdk5uVXdXRFYyVFRKeVRWTnJaalI2WW5wUFFtWnNaWEpKTnpRS1NFSlNNVkpJVWtWT1FrYzBSVzVIUlU1UlBUMEtMUzB0TFMxRlRrUWdRMFZTVkVsR1NVTkJWRVV0TFMwdExRbz0ifX19fQ==",
          "integratedTime": 1663104796,
          "logIndex": 3487994,
          "logID": "c0d23d6ad406973f9559f3ba2d1ca01f84147d8ffc5b8445c224f98b9591801d"
        }
      },
      "Issuer": "https://token.actions.githubusercontent.com",
      "Subject": "https://github.com/distroless/nginx/.github/workflows/release.yaml@refs/heads/main",
      "run_attempt": "1",
      "run_id": "3048410944",
      "sha": "ca4381349d1ae0b5128db6d21339a8cb44576feb"
    }
  }
]
```

You can verify that the image was built in Github Actions in this repository from the `Issuer` and `Subject` fields.
</details>

## Build

This image is built with [melange](https://github.com/chainguard-dev/melange) and [apko](https://github.com/chainguard-dev/apko).

