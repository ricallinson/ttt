# Tic Tac Toe

[![Build Status](https://travis-ci.org/ricallinson/ttt.svg)](https://travis-ci.org/ricallinson/ttt)

## Test

    go test -cover

## Benchmark

    go test -bench=. -run=BENCH_ONLY

## Line Coverage

    go test -coverprofile=coverage.out; go tool cover -html=coverage.out
