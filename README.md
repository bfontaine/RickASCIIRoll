# RickASCIIRoll

`rickasciiroll` is a TCP server that serves a terminal-based version of the
good ol' Rick Astley video.

## Build

In order to compile using `go build .` you need a Go file in the root
directory that defines `frames`, a `[]string` variable. This is a hacky way to
embed the movie in the final binary.

See under [`moviemaker`](./moviemaker) for instructions to build that Go file.
It downloads Rick Astley’s famous video off YouTube; extract its frames; render
them as text; then write them down in a `movie.go` file. The code should work
with any YouTube video but you may need to tweak a few variables for that.

The base code of the server is from [`tellmenet`][tellmenet].

[tellmenet]: https://github.com/bfontaine/tellmenet
