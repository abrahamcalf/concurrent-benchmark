Work In Progress...

# concurrent-benchmark

[![test][test-badge]][test]
[![lint][lint-badge]][lint]
[![codeql][codeql-badge]][codeql]
[![report-card][report-card-badge]][report-card]
[![npm][npm-badge]][npm]
[![license][license-badge]][repo]

> Benchmark scripts concurrently.

## Description

Right now I am just playing around, but I am planning (maybe) to make something serious in the future.

`concurrent-benchmark` has the alias `cb`.

## Install

```bash
$ yarn global add concurrent-benchmark
```

## Usage

#### CLI

Let's take `fixture.js` as an example.
It has a `setTimeout` that waits for 2 seconds.

```bash
$ time cb \
    'node fixture.js' \
    'node fixture.js' \
    'node fixture.js' \
    'node fixture.js'
```

So instead of taking using 1 thread (~8s) we are going to use 4 threads (~2s), one for each command.
The output of this will get a few console.logs be something like:

```bash
real    0m2.289s
user    0m0.157s
sys     0m0.237s
```

Where `real` is the time it took to run the command.

#### Package

Not yet implemented.

```js
import cb from 'concurrent-benchmark';

cb.run([
  {
    name: 'foo',
    script: 'echo foo',
  },
  {
    name: 'bar',
    script: 'echo bar',
  },
]);
```

[npm]: https://npmjs.org/package/concurrent-benchmark
[repo]: https://github.com/abranhe/concurrent-benchmark
[test]: https://github.com/abranhe/concurrent-benchmark/actions/workflows/test.yml
[lint]: https://github.com/abranhe/concurrent-benchmark/actions/workflows/lint.yml
[codeql]: https://github.com/abranhe/concurrent-benchmark/actions/workflows/codeql.yml
[report-card]: https://goreportcard.com/report/github.com/abranhe/concurrent-benchmark
[npm-badge]: https://img.shields.io/npm/v/concurrent-benchmark
[license-badge]: https://img.shields.io/npm/l/concurrent-benchmark
[test-badge]: https://github.com/abranhe/concurrent-benchmark/actions/workflows/test.yml/badge.svg
[lint-badge]: https://github.com/abranhe/concurrent-benchmark/actions/workflows/lint.yml/badge.svg
[codeql-badge]: https://github.com/abranhe/concurrent-benchmark/actions/workflows/codeql.yml/badge.svg
[report-card-badge]: https://goreportcard.com/badge/github.com/abranhe/concurrent-benchmark
