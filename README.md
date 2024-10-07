# Go fullstack template

## Installation & Set Up

1. Install [the Go Programming Language](https://go.dev)
2. Install [Air](https://github.com/air-verse/air) by running:

```sh
go install github.com/air-verse/air@latest
```

3. Install [Templ](https://templ.guide/) by running:

```sh
go install github.com/a-h/templ/cmd/templ@latest
```

4. Install [Node.js](https://nodejs.org/en)
5. Install Go deps with `go mod tidy`
6. Install Node deps with `npm install`

## VSCode setup

To best use this template install Go & Templ LSPs:

- [gopls](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- [templ](https://marketplace.visualstudio.com/items?itemName=a-h.templ)

## Neovim setup

You can install `gopls` and `templ` with Mason.

## Running

Just run:

```sh
npm run dev
```

This will start 3 processes:

- Tailwind watch
- Air
- Templ watch


## Building

Compile the Go binary:

```sh
go build main.go
```

And ship it along the `public` folder

