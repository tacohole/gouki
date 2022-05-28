# GOUKI

![gouki throwing a big ass fireball](https://github.com/tacohole/gouki/blob/main/gouki.png?raw=true)

## WHAT?

- A re-implementation of the [Heroku CLI](https://github.com/heroku/cli) in Golang
- Uses [heroku-go](https://github.com/bgentry/heroku-go) for interacting with the Heroku API
- Uses [cobra](https://github.com/spf13/cobra) for command

## WHY?

- I use the CLI every day
- There are some things I want to improve about it, namely performance and error messages
- I like writing command line tools

## Project Goals

- Support for all existing [Heroku CLI commands](https://devcenter.heroku.com/articles/heroku-cli-commands) and functionality
- Identical command syntax, arguments, and flags
- Support for MacOS, Windows, and common Linux distros
- Support for common Heroku CLI plug-ins like `pg-extras`, `builds`, and `repo`

## Project Status

- We are at v-0.04, it doesn't do anything yet, please don't perceive me
