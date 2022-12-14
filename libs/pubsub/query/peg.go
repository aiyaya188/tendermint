package query

// Normally I would use go run,
// but the "Code generated by" comment for peg includes the full arg0,
// which includes an unpredictable temporary directory,
// resulting in a nondeterminstic generated source file.
// Using go build is the workaround as detailed in https://github.com/pointlander/peg/issues/129.

//go:generate go build -o ./.bin/peg github.com/pointlander/peg
//go:generate ./.bin/peg -inline -switch query.peg
