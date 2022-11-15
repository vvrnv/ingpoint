# Ingpoint

script for getting k8s ingress public endpoints for current context

# Install & build from source
requirements:
- `go`
- `$GOPATH/bin`

### go
```sh
go install github.com/vvoronov2/ingpoint@latest
```

### build from source
```sh
git clone https://github.com/vvoronov2/ingpoint.git
cd ingpoint
go build
```

# Usage
requirements:
- `kubectl`

```sh
ingpoint
```

TODO:
- rewrite to kuberctl plugin
