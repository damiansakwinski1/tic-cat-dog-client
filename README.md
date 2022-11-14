# Tic Cat Dog

**WORK IN PROGRESS!**

Disclaimer: This repository contains code that is considered unfinished.

+ Some parts of it may not be refactored according to the rules of clean code or the principles of effective go.
+ Some parts of it may not work properly.

## Description

Tic-tac-toe but it's cats and dogs...

## Running (Linux)

## Using Docker

1. Make sure you have Docker and Make installed on your system.
2. Make sure you're in the project's root directory.
3. Run `make run-game`.

### Locally

1. Make sure your system meets the dependencies of [Ebitenengine](https://ebitengine.org/en/documents/install.html).
2. Run `export ASSETS_ABS_PATH=<path to your assets path>`. If your assets folder is
   in `/home/user/Projects/tic-cat-dog/assets`, then the command will look as
   follows: `export ASSETS_ABS_PATH=/home/user/Projects/tic-cat-dog/assets`.
3. Make sure you're in the project's root directory.
4. Run `go build`.
5. Run `./tic-cat-dog-client`

## Running tests

1. Make sure you're in the projects root directory
2. Run `make run-tests`.