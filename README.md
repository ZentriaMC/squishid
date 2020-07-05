# squishid-go

*Squish IDs into shorter strings*

squishid is an alternative base32 encoding format that, compared to the
standard format, makes IDs shorter by *squishing* sequentially repeating
elements.

## Motivation

Web applications often use entity IDs in URLs as raw numbers or by transforming
them into a string. As the IDs can be quite large numbers, they can make the
URL overly long so that it looks unappealing or even suspicious.

Squished IDs don't necessarily make URLs more beautiful, but in many cases they
definitely make them shorter.

```
number:   https://example.com/post/4611686018427387904/help-ids-are-too-long-what-to-do-squish-them
base64:   https://example.com/post/AAAAAAAAAEA/help-ids-are-too-long-what-to-do-squish-them
base32:   https://example.com/post/aaaaaaaaaaaea/help-ids-are-too-long-what-to-do-squish-them
squishid: https://example.com/post/nnae/help-ids-are-too-long-what-to-do-squish-them
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/kadfak/squishid-go/squishid"
)

func main() {
	fmt.Println(squishid.Squish(666))

	id, err := squishid.Restore("4x")
	if err != nil {
		panic("invalid ID")
	}
	fmt.Println(id)
}
```

```
Result:
4x
666
```

## Alphabet

```
abcdefghijkmopqrstuwxyz123456789

squish indicators:
lovn
```

## Comparison

```
number:   0
base32:   aaaaaaaaaaaaa
squishid: a

number:   4611686018427387904
base32:   aaaaaaaaaaaea
squishid: nnae

worst case:
number:   14242959524133701664
binary:   1100 01011 01010 01001 01000 00111 00110 00101 00100 00011 00010 00001 00000
base32:   eceedcrzfcu4k
squishid: abcdefghijkmo
```

## Limitations

squishid is useful only when the IDs are sequential. Random IDs don't work well
with squishid because they, in general, lack sequentially repeating elements
that are required for the squishing algorithm to make the generated string
shorter. See the worst case example above.

Currently, only 64-bit integers are supported.

## License

[MIT](LICENSE)
