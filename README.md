# zlpp
A tiny glue utility that takes a stream of zerolog (or similarly structured log) output and passes it back through zerolog.ConsoleWriter, to pretty print for your console.

## Installation

`go get -u github.com/yeahphil/zlpp`

## Usage

Pipe a stream of log into zlpp:

`cat my_logs.txt | zlpp`

Reformat a file directly:

`zlpp my_logs.txt`