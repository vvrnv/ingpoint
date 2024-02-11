# Ingpoint

app to get all unique ingress hosts for current kubernetes context

## Install

### homebrew

```sh
brew install vvrnv/tap/ingpoint
```

### go or build from source code

requirements:

- `go` installed

```sh
go install github.com/vvrnv/ingpoint@latest
```

or

```sh
git clone https://github.com/vvrnv/ingpoint.git
cd ingpoint
go build -o ingpoint
./ingpoint
```

### download a biany from [release page](https://github.com/vvrnv/ingpoint/releases)

## Usage

requirements:

- `kubectl`

```sh
➜  ~ kubectl config use-context dev
Switched to context "dev".

➜  ~ ingpoint
jira-666.example.domain.com
test.domain.com
*.example.domain.com
```
