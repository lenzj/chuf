chuf(1)
; This is a scdoc source file for generating a man page
; See https://git.sr.ht/~sircmpwn/scdoc

# NAME

chuf - transform a stream by sending specific chunks through a filter

# SYNOPSIS

*chuf* _BEGIN_ _END_ _FILTER_

# DESCRIPTION

*chuf* is a command line utility that transforms stdin to stdout by sending
predefined chunks through the specified _FILTER_.  A chunk is defined by a
_BEGIN_ and _END_ byte sequence within the stream. Anything not within a chunk
is passed along unmodified. 

_BEGIN_
	Byte sequence designating beginning of chunk

_END_
	Byte sequence designating end of chunk

_FILTER_
	Command filter to transform chunk. A multiple parameter command filter
	must be enclosed in quotes.

# EXAMPLE

$ *echo "the quick {U}brown fox{R} jumped over the lazy dog" | \\++
chuf {U} {R} "tr [:lower:] [:upper:]"*++
the quick BROWN FOX jumped over the lazy dog

# BUGS

Bug reporting instructions:++
<https://blog.lenzplace.org/about/contact.html>

# SEE ALSO

Website for chuf:++
<https://git.lenzplace.org/lenzj/chuf>
