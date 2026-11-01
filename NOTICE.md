# Inherited source and license provenance

`src/JS/BigInt.purs` and `src/JS/BigInt.js` are byte-identical to
[`purescript-js-bigints` v2.2.1](https://github.com/purescript-contrib/purescript-js-bigints/tree/36a7d8ac75a7230043ae511f3145f9ed130954a9),
commit `36a7d8ac75a7230043ae511f3145f9ed130954a9`.

The restored [LICENSE](LICENSE) is the unchanged upstream MIT text,
Copyright (c) 2022 Jan Schulte. Its SHA-256 is
`4a56dd0010e318fc1ebb942ee91a858ddb2a4de76f8085ffc5d84c58d1eadd26`.
This restoration preserves the upstream notice. The additional attribution
below applies to the Stack Overflow material retained in the JavaScript file.

Upstream metadata is inconsistent at that commit:

- [LICENSE](https://github.com/purescript-contrib/purescript-js-bigints/blob/36a7d8ac75a7230043ae511f3145f9ed130954a9/LICENSE),
  [bower.json](https://github.com/purescript-contrib/purescript-js-bigints/blob/36a7d8ac75a7230043ae511f3145f9ed130954a9/bower.json)
  and [spago.dhall](https://github.com/purescript-contrib/purescript-js-bigints/blob/36a7d8ac75a7230043ae511f3145f9ed130954a9/spago.dhall)
  specify MIT.
- [package.json](https://github.com/purescript-contrib/purescript-js-bigints/blob/36a7d8ac75a7230043ae511f3145f9ed130954a9/package.json)
  specifies ISC. The registry package's `purs.json` records `MIT AND ISC`.

This port contains none of those npm, Bower or Dhall manifests. The discrepancy
above records upstream history; it is not a local package metadata declaration.
Any future standalone package must describe the licenses of the files it
actually includes rather than copy the upstream aggregate identifier.

## JavaScript conversion adapted from Stack Overflow

The chunk-based conversion in `src/JS/BigInt.js`, `fromStringAsImpl`, is adapted
from [Nina Scholz's answer to “Base 36 to BigInt”](https://stackoverflow.com/a/55646905).
Author: [Nina Scholz](https://stackoverflow.com/users/1447675/nina-scholz).
The answer was posted on 12 April 2019 and last edited on 13 April 2019 (UTC).
It is licensed under
[Creative Commons Attribution-ShareAlike 4.0 International](https://creativecommons.org/licenses/by-sa/4.0/)
([full license](https://creativecommons.org/licenses/by-sa/4.0/legalcode.en)).
The [answer API record](https://api.stackexchange.com/2.3/answers/55646905?site=stackoverflow&filter=withbody)
identifies that license explicitly; the dates also fall under Stack Overflow's
[CC BY-SA 4.0 contribution policy](https://stackoverflow.com/help/licensing).

The inherited implementation changes the chunk size from 10 to 6, wraps the
conversion in curried PureScript FFI arguments and `Maybe` constructors, and
adds negative-number preprocessing and invalid-number error handling. The
original source comment is retained. This notice changes no source code.

The adapted Stack Overflow material is distributed under CC BY-SA 4.0,
including its attribution and share-alike terms. Keep this attribution, the
license reference and the description of adaptations with that material.
The upstream MIT notice is retained alongside it; it does not replace the
terms applicable to the Stack Overflow material. No endorsement by the
original author is implied.

## Go adaptation

The Go adaptation in `src/JS/BigInt.go` is distributed under the MIT terms in
[LICENSE](LICENSE), as selected by this port's maintainer.
For this Go file: `SPDX-License-Identifier: MIT`.
Its `FromStringAsImpl` calls `math/big.Int.SetString`. The retained upstream
copyright attribution identifies inherited material; this notice does not
attribute authorship of the Go adaptation to upstream. The explicit MIT
choice for the Go adaptation does not change the separate terms of the
JavaScript material described above.
