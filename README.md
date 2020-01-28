# Number quizz

## Installation

Run `git clone https://github.com/CaptainQuirk/numberquizz.git`

## Configuration

### Configure the GOBIN variable

- Create a [globally ignored][globally_ignored_rcd] `.rc.d` directory with a `go.zsh`
file in it with the following information

```zsh
#!/usr/bin/zsh

go env -w GOBIN=<path-to-project>/numberquizz/bin
```

> This will reference the bin directory inside the repository as the location
> where build binary should go

### Add the GOBIN directory to the path

- Create a [globally ignored][globally_ignore_envrc] `.envrc` file to work with
[direnv][direnv]

```sh
PATH_add /home/leonard/Projects/CaptainQuirk/LearnGo/Code/numberquizz/bin
```

- Run `direnv allow`

## Build

- Run `make build` or `go install`

## Run

- Run `make run` or `numberquizz`

[direnv]: https://direnv.net/
[globally_ignored_rcd]: https://github.com/CaptainQuirk/.gitglobal/blob/master/gitignore#L30
[globally_ignore_envrc]: https://github.com/CaptainQuirk/.gitglobal/blob/master/gitignore#L56
