# git-wip

Simple tool which commit all files to WIP commit with sequence number.

## Usage

For WIP commit just run `git wip`

## Download a latest binary

### Mac

```
curl -L https://github.com/ondrejsika/git-wip-bin/raw/master/latest/git-wip-darwin-amd64 -o git-wip && chmod +x git-wip && sudo mv git-wip /usr/local/bin
```

### Linux

```
curl -L https://github.com/ondrejsika/git-wip-bin/raw/master/latest/git-wip-linux-amd64 -o git-wip && chmod +x git-wip && sudo mv git-wip /usr/local/bin
```

## Build from source (go build)

```
make build
sudo mv git-wip /usr/local/bin/git-wip
```

## Changlog

- `v1.0.0` - First minimal working version
