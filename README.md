# chuf

**chuf** is a command line utility that transforms stdin to stdout by sending
predefined chunks through the specified _FILTER_.  A chunk is defined by a
_BEGIN_ and _END_ byte sequence within the stream. Anything not within a chunk
is passed along unmodified.

## Synopsis

```text
chuf BEGIN END FILTER
```

_BEGIN_ : Byte sequence designating beginning of chunk

_END_ : Byte sequence designating end of chunk

_FILTER_ : Command filter to transform chunk. A multiple parameter command
filter must be enclosed in quotes.

## Example 1 - Simple text markup using unix "tr" command.

```text
$ echo "the quick {U}brown fox{R} jumped over the lazy dog" | \\
chuf {U} {R} "tr [:lower:] [:upper:]"
the quick BROWN FOX jumped over the lazy dog
```

## Example 2 - Expand markdown within html document.

```text
$ cat example.mdhtml
<!DOCTYPE html>
<html>
    <head>
        <title>Example</title>
    </head>
    <body>
<md>
# Markdown table example

| Name   | Hobby        | Age   |
| ------ |:------------:| -----:|
| Bob    | golfing      |    18 |
| Monica | programming  |    32 |
</md>
    </body>
</html>
$ chuf '<md>' '</md>' markdown < example.mdhtml
<!DOCTYPE html>
<html>
    <head>
        <title>Example</title>
    </head>
    <body>
<h1>Markdown table example</h1>
<table>
<thead>
<tr>
<th> Name   </th>
<th style="text-align:center;"> Hobby        </th>
<th style="text-align:right;"> Age   </th>
</tr>
</thead>
<tbody>
<tr>
<td> Bob    </td>
<td style="text-align:center;"> golfing      </td>
<td style="text-align:right;">    18 </td>
</tr>
<tr>
<td> Monica </td>
<td style="text-align:center;"> programming  </td>
<td style="text-align:right;">    32 </td>
</tr>
</tbody>
</table>
    </body>
</html>
```

## Compiling from source

### Dependencies

* Go compiler (v1.12 or later).
* Go package [chunkio](https://git.lenzplace.org/lenzj/chunkio)
* Go package [testcli](https://git.lenzplace.org/lenzj/testcli) to run tests.
* [scdoc](https://git.sr.ht/~sircmpwn/scdoc/) utility to generate the man page.
  Only needed if changes to man page source (chuf.1.scd) are made.
* [pgot](https://git.lenzplace.org/lenzj/pgot) utility to process files in the templates sub
  folder.  Only needed if changes to README.md, LICENCE, Makefile etc. are
  needed.

### Installing

```text
$ make
# make install
```

## Running the tests

```text
$ make check
```

## Contributing

If you have a bugfix, update, issue or feature enhancement the best way to reach
me is by following the instructions in the link below.  Thank you!

<https://blog.lenzplace.org/contact>

## Versioning

I follow the [SemVer](http://semver.org/) strategy for versioning. The latest
version is listed in the [releases](/lenzj/chuf/releases) section. 

## License

This project is licensed under a BSD two clause license - see the
[LICENSE](LICENSE) file for details.
