# PostgreSQL Go query builder benchmark

The benchmark code here was based on [golang-sql-builder-benchmark](https://github.com/elgris/golang-sql-builder-benchmark).

It compares the following:

* [Squirrel](https://github.com/lann/squirrel)
* [pgq](https//github.com/henvic/pgq)
* [sqrl](https://github.com/elgris/sqrl)

Details:

* pgq is PostgreSQL specific and requires you to use a driver that communicates with PostgreSQL using its native interface instead of the slower database/sql, such as [pgx](https://github.com/jackc/pgx).
* pgq and sqrl are forks of the Squirrel query builder.
* sqrl was written with an aim on query builder performance at the cost of thread safety.
* pgq was written with an aim of providing the best experience for PostgreSQL.
* The API for pgq, sqrl, and Squirrel are pretty much the same, with some minor exceptions.

# Benchmarks

In short:

* pgq and sqrl numbers are pretty close.
* Squirrel execution time was at least twice slower than pgq and sqrl.
* Squirrel allocates twice as many bytes per operation as pgq and sqrl.
* Squirrel does about 3 times more allocations.
* pgq is slightly slower than sqrl, but allocates less memory per operation.

<details>
<summary>Results of the benchmark execution</summary>

```
$ go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/henvic/pgqbenchmark
cpu: Intel(R) Core(TM) i7-7820HQ CPU @ 2.90GHz
BenchmarkSelectSimple/squirrel-8         	  191125	      5585 ns/op	    2792 B/op	      56 allocs/op
BenchmarkSelectSimple/pgq-8              	  716359	      1429 ns/op	     960 B/op	      18 allocs/op
BenchmarkSelectSimple/sqrl-8             	  770060	      1453 ns/op	    1296 B/op	      20 allocs/op
BenchmarkSelectConditional/squirrel-8    	  133291	      8732 ns/op	    4333 B/op	      87 allocs/op
BenchmarkSelectConditional/pgq-8         	  538165	      1973 ns/op	    1432 B/op	      23 allocs/op
BenchmarkSelectConditional/sqrl-8        	  619050	      1877 ns/op	    1736 B/op	      24 allocs/op
BenchmarkSelectComplex/squirrel-8        	   41720	     28862 ns/op	   13148 B/op	     290 allocs/op
BenchmarkSelectComplex/pgq-8             	  128432	      8095 ns/op	    4793 B/op	      93 allocs/op
BenchmarkSelectComplex/sqrl-8            	  155815	      7484 ns/op	    5281 B/op	      93 allocs/op
BenchmarkSelectSubquery/squirrel-8       	   53664	     21922 ns/op	   10748 B/op	     220 allocs/op
BenchmarkSelectSubquery/pgq-8            	  182692	      6625 ns/op	    4593 B/op	      79 allocs/op
BenchmarkSelectSubquery/sqrl-8           	  201844	      5657 ns/op	    4745 B/op	      75 allocs/op
BenchmarkSelectMoreComplex/squirrel-8    	   29647	     39905 ns/op	   19222 B/op	     415 allocs/op
BenchmarkSelectMoreComplex/pgq-8         	   94238	     12471 ns/op	    8251 B/op	     150 allocs/op
BenchmarkSelectMoreComplex/sqrl-8        	   93992	     12038 ns/op	    8963 B/op	     146 allocs/op
BenchmarkInsert/squirrel-8               	  169060	      6817 ns/op	    3417 B/op	      77 allocs/op
BenchmarkInsert/pgq-8                    	  568092	      2097 ns/op	    1288 B/op	      21 allocs/op
BenchmarkInsert/sqrl-8                   	  861246	      1476 ns/op	    1176 B/op	      18 allocs/op
BenchmarkUpdate/squirrel-8               	  141913	      8397 ns/op	    4009 B/op	      91 allocs/op
BenchmarkUpdate/pgq-8                    	  519693	      2207 ns/op	    1344 B/op	      29 allocs/op
BenchmarkUpdate/sqrl-8                   	  657476	      1690 ns/op	    1296 B/op	      26 allocs/op
BenchmarkUpdateMap/squirrel-8            	  135675	      8410 ns/op	    4081 B/op	      93 allocs/op
BenchmarkUpdateMap/pgq-8                 	  450721	      2544 ns/op	    1411 B/op	      31 allocs/op
BenchmarkUpdateMap/sqrl-8                	  559718	      2023 ns/op	    1347 B/op	      28 allocs/op
BenchmarkDelete/squirrel-8               	  185583	      6336 ns/op	    2728 B/op	      65 allocs/op
BenchmarkDelete/pgq-8                    	 1746862	       684.0 ns/op	     432 B/op	      11 allocs/op
BenchmarkDelete/sqrl-8                   	 2194473	       533.1 ns/op	     576 B/op	      10 allocs/op
PASS
ok  	github.com/henvic/pgqbenchmark	35.768s
```

</details>

Disclaimer: This benchmark was published by the author of pgq, with tests adapted from a benchmark of author of sqrl.
