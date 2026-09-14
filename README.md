# gopurs-js-bigints

Go adaptation of the `JS.BigInt` PureScript module. The public types and
operations are in `src/JS/BigInt.purs`; `BigInt.go` supplies the Go FFI and
`BigInt.js` the JavaScript FFI.

## Local Go development

Follow the [local Go development guide](../gopurs/README.md#develop-one-library-locally)
for the TAST compiler, Spago and sibling checkouts. Edit `spago.go.yaml`;
`spago.yaml` is its tracked link. Tests enter through `test/Main.purs`
(`Test.Main`), and `./bin/test` builds and runs them with Go.

The package has no npm manifest. API semantics and the provenance/license
context remain assigned to lot 12 in the [maintenance plan](../gopurs/todo.md).
