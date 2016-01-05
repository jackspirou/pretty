# pretty [![Go Report Card](http://goreportcard.com/badge/jackspirou/pretty)](http://goreportcard.com/report/jackspirou/pretty)
A CLI tool that automatically identifies and formats JSON or XML.
It is written in Go, is simple, and has no third party dependencies.

- JSON formatting was tested against a 189.8 MB file, no crashing resulted.
- XML formatting was tested against a 71.1 MB file, no crashing resulted.

## Install
### go get
`$ go get -u github.com/jackspirou/pretty`

## JSON
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

## XML
Here is an example of a messy `.xml` file:
```bash
$ cat messy.xml
<config><diagnostics><path>C:\temp</path>
<proxy usedefault="true"/></diagnostics></config>
```

Let's make it pretty:
```bash
$ cat messy.xml | pretty
<config>
	<diagnostics>
		<path>C:\temp</path>
		<proxy usedefault="true"></proxy>
	</diagnostics>
</config>%
```

If you wanted to do something other than print the pretty result to the terminal, pipe it to somewhere:
```bash
$ cat messy.xml | pretty | somewhere
```
