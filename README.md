# pretty
A CLI tool that automatically identifies and formats JSON or XML.
It is written in Go, is simple, and has no third party dependencies.

## Install
`go get -u github.com/jackspirou/pretty`

## Usage

Here is an example of a messy `.json` file:
```bash
$ cat messy.json
{"type":"text","$t":"day: Monday, menu: Paninis, baguettes, spaghetti / Chicken curry"}
```

Let's make it pretty:
```bash
$ cat messy.json | pretty
{
	"type": "text",
	"$t": "day: Monday, menu: Paninis, baguettes, spaghetti / Chicken curry"
}
```

If you wanted to do something other than print the pretty result to the terminal, pipe it to somewhere:
```bash
$ cat messy.json | pretty | somewhere
```
