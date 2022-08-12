# GOUKI

![gouki throwing a big ass fireball](https://github.com/tacohole/gouki/blob/main/gouki.png?raw=true)

## WHAT?

- A re-implementation of the [Heroku CLI](https://github.com/heroku/cli) in Golang
- Uses [heroku-go](https://github.com/bgentry/heroku-go) for interacting with the Heroku API
- Uses [cobra](https://github.com/spf13/cobra) for CLI scaffolding

## WHY?

- I use the Heroku CLI every day
- There are some things I want to improve about it, namely performance and error messages
- I like writing command line tools

## Project Goals

- Support for all existing [Heroku CLI commands](https://devcenter.heroku.com/articles/heroku-cli-commands) and functionality
- Similar command syntax, arguments, and flags
- Support for MacOS, Windows, and common Linux distros
- Support for common Heroku CLI plug-ins like `pg-extras`, `builds`, and `repo`

## Project Status

### What works:

- `access`
  - `access add`, `access remove`, `access update`
- `apps`
- `logs`

### What doesn't work:

- everything else

## Pre-reqs

- Go v1.18+
- Heroku account
- Heroku API token

## Installation

- `gh repo clone tacohole/gouki`
- `cd gouki`
- `go install`
- `cp .env.sample .env`
- `HEROKU_USERNAME` set to your Heroku username
- `HEROKU_API_TOKEN` containing a valid Heroku API token

## Usage

`gouki cmd subcmd --flagName flagvalue`

### examples

invocation: `gouki apps`
output: `??`
