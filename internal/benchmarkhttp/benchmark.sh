#!/bin/sh
# Before running the benchmark use `go run .` to start the target server.

bombardier --print i,r localhost:8080/sanity > results.txt
bombardier --print i,r localhost:8080/sanity-gzip >> results.txt
bombardier --print i,r localhost:8080/bytes >> results.txt
bombardier --print i,r localhost:8080/bytes-gzip >> results.txt
bombardier --print i,r localhost:8080/template >> results.txt
bombardier --print i,r localhost:8080/template-gzip >> results.txt
