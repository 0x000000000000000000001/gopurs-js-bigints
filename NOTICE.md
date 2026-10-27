# Inherited source and license provenance

`src/JS/BigInt.purs` and `src/JS/BigInt.js` are byte-identical to
[`purescript-js-bigints` v2.2.1](https://github.com/purescript-contrib/purescript-js-bigints/tree/36a7d8ac75a7230043ae511f3145f9ed130954a9),
commit `36a7d8ac75a7230043ae511f3145f9ed130954a9`.

The restored [LICENSE](LICENSE) is the unchanged upstream MIT text,
Copyright (c) 2022 Jan Schulte. Its SHA-256 is
`4a56dd0010e318fc1ebb942ee91a858ddb2a4de76f8085ffc5d84c58d1eadd26`.
This restoration preserves the notice for inherited material.

Upstream metadata is inconsistent at that commit:

- [LICENSE](https://github.com/purescript-contrib/purescript-js-bigints/blob/36a7d8ac75a7230043ae511f3145f9ed130954a9/LICENSE),
  [bower.json](https://github.com/purescript-contrib/purescript-js-bigints/blob/36a7d8ac75a7230043ae511f3145f9ed130954a9/bower.json)
  and [spago.dhall](https://github.com/purescript-contrib/purescript-js-bigints/blob/36a7d8ac75a7230043ae511f3145f9ed130954a9/spago.dhall)
  specify MIT.
- [package.json](https://github.com/purescript-contrib/purescript-js-bigints/blob/36a7d8ac75a7230043ae511f3145f9ed130954a9/package.json)
  specifies ISC. The registry package's `purs.json` records `MIT AND ISC`.

The inherited JavaScript `fromStringAsImpl` comment identifies a copied
[Stack Overflow example](https://stackoverflow.com/questions/55646698/base-36-to-bigint).
That comment is retained. Restoring the upstream license does not resolve the
example's attribution and licensing requirements or the metadata discrepancy.

The Go adaptation in `src/JS/BigInt.go` is distributed under the MIT terms in
[LICENSE](LICENSE), as selected by this port's maintainer.
For this Go file: `SPDX-License-Identifier: MIT`.
Its `FromStringAsImpl` calls `math/big.Int.SetString`. The retained upstream
copyright attribution identifies inherited material; this notice does not
attribute authorship of the Go adaptation to upstream. The explicit MIT
choice for the Go adaptation does not resolve the inherited metadata and
example-attribution questions above.
