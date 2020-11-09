# keyp

keyp is a tool to keep public keys up to date.

## Usage

### Collect public keys

Collect public keys from [GitHub](https://github.com/) ( user: `alice`, user: `bob` and team: `myorg/administrators` ).

``` console
$ export GITHUB_TOKEN=xxXXXXxxxxXXXXXXxxxXXX
$ keyp collect -b github -u alice -u bob -t myorg/administrators
ssh-rsa C1yc2EAAAADAQABAAACAQCmnCCt6PyH9jLZbPkMijSJYIu14nhxyFVw9M9eAkgcMQ3EsKf86GWlGPDfZcqcDqI+LP7LKQk4kAlmLOQQMavthrhGEURxxdX0Yk2A6pHjG3zrcW7X30ZBMwOzX/a6EWkPXPwPH6LcP3rM9yEIg95f2JntxO3z7l/8QjzJCoPIlqyoX4I7bxHus/rZVnRNh9C0PbUejbg/iWaTTxkNglSeEYpW+ID2k/4Absisa5XyY2zEOMw+6OyfRL9AlfGYv545J0g90qWS11iRSFnMR7A6FNUea/pVESIMmmBI56Ne+S8NmwR724u3d5kNJxuAKpmtThNPdxW/vmtuc5XBZgtPX/rzdAW0TQZvpVoLnoKaqYgfIpcrrkSAoPlcVxfq/NrpVlbIi6c9rZRZR4dcqmAK2eBGuDQZBiYJSESuPbE6i08GGnM8OblD8pshVeGMStztR+NuXywIXbRpyqNF2VNjil8r4qGNW9AB8ZVUB/1s6U8oxvtlbABoxXLrdNlKj3rl2YoPIZyCLAp9QDch8P1SnmQTEZK67YY5KNQrJZ2ql7pblo84JqsRbwuOrexTz6xrbBWFZMHWorFqF8ryX0LOw9TIaHbYqleynhqJ0a8VJHZMmwndYKjw3brtJ3SfCpXU0826LOExWXcjfBqHK65gM+MQ==
ssh-rsa AAADAQABAAACAQC8IEt4MqBed/yXQyjUTCZRdZoCUNhm0bEkOV8Ef5TduQvMIPDpBYyYIvFz7jxJyShPoiTMtIUnkkA2aDF0jhujFzqKmYm9H2tS7Tpf5iNwRJgJJJWv674tGUcu+6+ZadmDBQ//dwo8XWTHxmkWfgaybxs8/o0AlwZQ4pYFcky0q+/qP4cwPAmRW0rGCo0E5BhS/5eGssoLBXu4/Hcaz/93H8AtAe1UQrlCKma0rj0HIA9A9Q9EQtunw/zJTBtTyzE/TvxKcSMNulgdVmFSRmU6l84Ftc6tZPoiaCnxcvQUyjCEeQfy4DbtCWe1tEubyKeBLBTXTnpqWA3Gs9GryQA/bR7Ivan/03FshLFeVVnbvvO11sKNvkAJ8u417Q2/G9bcB1H30Xa9PSRE+2CbQ2maafhPVL17TJVBvkDCM5trmwxfM2tdlKA7R+mTj9nIrSLN4BYrge8IZ1fesC/sKMlMwhNEOrQYQQIZMIx8hfLAS37D8wbUPRodQFJsolrK6cHlNICR/TLcijNhCeHJkD8448EuJn1BCbYKglG7eUYKLbMXcVJAoTPlFTHPU80oaHJhmpLe0vFSxrhWVf/ha81zRefXOiye7Pbn/h+sa2qsKTnAMShpS1m+RP7QmHNmFAbHlPeTlnd0oJI/bt5Mysn5HHjX4vAJdQ==
[...]
```

### Update authorized_keys

Update `~/.ssh/authorized_keys` of `ubuntu` user using public keys managed by [STNS](https://stns.jp/).

``` console
$ keyp update-authorized-keys ubuntu -b stns -u k1low -g developers
```

## Support Backend

### GitHub

``` console
$ keyp collect -b github -u alice -u bob -t myorg/administrators
```

#### Environment

| key | description |
| --- | --- |
| `GITHUB_TOKEN` | GitHub Personal Access Token (required) |
| `GITHUB_ENDPOINT` | ( GitHub Enterprise ) GraphQL API Endpoint. ex `https://git.mycompany.com/api/graphql` |

### STNS

``` console
$ keyp update-authorized-keys ubuntu -b stns -u k1low -g developers
```

`stns` backend load `/etc/stns/client/stns.conf` first.

#### Environment

| key | description |
| --- | --- |
| `STNS_API_ENDPOINT` | STNS API Endpoint. ex `https://stns.lolipop.io/v1` |
| `STNS_AUTH_TOKEN` | token authentication token |
| `STNS_USER` | basic authentication user |
| `STNS_PASSWORD` | basic authentication password |
| `STNS_SKIP_VERIFY` | skip verify certs |
| `STNS_REQUEST_TIMEOUT` | http request timeout |
| `STNS_REQUEST_RETRY` | http request of retries |

## Install

**deb:**

Use [dpkg-i-from-url](https://github.com/k1LoW/dpkg-i-from-url)

``` console
$ export KEYP_VERSION=X.X.X
$ curl -L https://git.io/dpkg-i-from-url | bash -s -- https://github.com/k1LoW/keyp/releases/download/v$KEYP_VERSION/keyp_$KEYP_VERSION-1_amd64.deb
```

**RPM:**

``` console
$ export KEYP_VERSION=X.X.X
$ yum install https://github.com/k1LoW/keyp/releases/download/v$KEYP_VERSION/keyp_$KEYP_VERSION-1_amd64.rpm
```

**homebrew tap:**

```console
$ brew install k1LoW/tap/keyp
```

**manually:**

Download binary from [releases page](https://github.com/k1LoW/keyp/releases)

**go get:**

```console
$ go get github.com/k1LoW/keyp
```
