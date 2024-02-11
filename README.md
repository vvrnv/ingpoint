# Ingpoint

app to get all unique ingress hosts for current kubernetes context

## Install

### homebrew

requirements:

- `go` installed

### go

```sh
go install github.com/vvrnv/ingpoint@latest
```

### build from source

```sh
git clone https://github.com/vvrnv/ingpoint.git
cd ingpoint
go build -o ingpoint
./ingpoint
```

## Usage

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
