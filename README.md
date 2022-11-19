# Ingpoint

script for getting ingress public endpoints for current kubernetes context

# Install & build from source
requirements:
- `go` installed

### go
```sh
go install github.com/vvoronov2/ingpoint@latest
```

### build from source
```sh
git clone https://github.com/vvoronov2/ingpoint.git
cd ingpoint
go build ingpoint.go
./ingpoint
```

# Usage
requirements:
- `kubectl`

```sh
➜  ~ kubectl config use-context dev
Switched to context "dev".

➜  ~ ingpoint
"jira-666.example.domain.com"
"test.domain.com"
"*.example.domain.com"
```

TODO:
- rewrite to kubectl plugin
