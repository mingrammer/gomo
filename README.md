[![Go Report Card](https://goreportcard.com/badge/github.com/mingrammer/gomo)](https://goreportcard.com/report/github.com/mingrammer/gomo) [![Travis Build](https://travis-ci.org/mingrammer/gomo.svg?branch=master)](https://travis-ci.org/mingrammer/gomo.svg?branch=master)

# gomo
Gomo is terminal based memo application written in Go.

## Who is it for?
This was created just for me. I often use terminal on my mac so it's annoying to swtich the windows between terminal and to-do application.

This is why I make this simple program.

So, this is for someone like me or for one want to use memo on terminal simply.

## Introduction
It provides minimal functions for memo or todo. not much functions.
**gomo** has just 4 commands now. `init`, `new`, `list`, `delete` are all!

Very simple to use.

## Installation
Just `go get github.com/mingrammer/gomo`

> `go` is required to use it now. I'll register it to `brew` and other package managers of many linux dist.

## Usage
* `gomo init` : init memo file on your machine
* `gomo new 'content'` : create new memo
* `gomo list` : list the all stored memos
* `gomo delete [number]` : delete a memo

What a so simple. 

This is example of use.

![Gomo Example](image/gomo-gif.gif)
