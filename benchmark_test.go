// Package pgqbenchmark is an adaption of https://github.com/elgris/golang-sql-builder-benchmark
// with a focus to compare Squirrel and pgq when used with PostgreSQL.
package pgqbenchmark

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/elgris/sqrl"
	"github.com/henvic/pgq"
)

// Use the Dollar placeholder format as required by PostgreSQL.
var (
	squirrelDollar = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	//lint:ignore U1000 bug; see https://github.com/dominikh/go-tools/issues/633
	sqrlDollar = sqrl.StatementBuilder.PlaceholderFormat(sqrl.Dollar)
)

func BenchmarkSelectSimple(b *testing.B) {
	testCases := []struct {
		name string
		f    func(b *testing.B)
	}{
		{
			name: "squirrel",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					squirrelDollar.Select("id").
						From("tickets").
						Where("subdomain_id = ? and (state = ? or state = ?)",
							1, "open", "spam").
						ToSql()
				}
			},
		},
		{
			name: "pgq",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					pgq.Select("id").
						From("tickets").
						Where("subdomain_id = ? and (state = ? or state = ?)",
							1, "open", "spam").
						SQL()
				}
			},
		},
		{
			name: "sqrl",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					sqrlDollar.Select("id").
						From("tickets").
						Where("subdomain_id = ? and (state = ? or state = ?)",
							1, "open", "spam").
						ToSql()
				}
			},
		},
	}
	for _, tc := range testCases {
		b.Run(tc.name, tc.f)
	}
}

func BenchmarkSelectConditional(b *testing.B) {
	testCases := []struct {
		name string
		f    func(b *testing.B)
	}{
		{
			name: "squirrel",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					qb := squirrelDollar.Select("id").
						From("tickets").
						Where("subdomain_id = ? and (state = ? or state = ?)",
							1, "open", "spam")

					if n%2 == 0 {
						qb = qb.GroupBy("subdomain_id").
							Having("number = ?", 1).
							OrderBy("state").
							Limit(7).
							Offset(8)
					}

					qb.ToSql()
				}
			},
		},
		{
			name: "pgq",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					qb := pgq.Select("id").
						From("tickets").
						Where("subdomain_id = ? and (state = ? or state = ?)",
							1, "open", "spam")

					if n%2 == 0 {
						qb = qb.GroupBy("subdomain_id").
							Having("number = ?", 1).
							OrderBy("state").
							Limit(7).
							Offset(8)
					}

					qb.SQL()
				}
			},
		},
		{
			name: "sqrl",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					qb := sqrlDollar.Select("id").
						From("tickets").
						Where("subdomain_id = ? and (state = ? or state = ?)",
							1, "open", "spam")

					if n%2 == 0 {
						qb.GroupBy("subdomain_id").
							Having("number = ?", 1).
							OrderBy("state").
							Limit(7).
							Offset(8)
					}

					qb.ToSql()
				}
			},
		},
	}
	for _, tc := range testCases {
		b.Run(tc.name, tc.f)
	}
}

func BenchmarkSelectComplex(b *testing.B) {
	testCases := []struct {
		name string
		f    func(b *testing.B)
	}{
		{
			name: "squirrel",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					squirrelDollar.Select("a", "b", "z", "y", "x").
						Distinct().
						From("c").
						Where("d = ? OR e = ?", 1, "wat").
						Where(squirrel.Eq{"f": 2, "x": "hi"}).
						Where(map[string]any{"g": 3}).
						Where(squirrel.Eq{"h": []int{1, 2, 3}}).
						GroupBy("i").
						GroupBy("ii").
						GroupBy("iii").
						Having("j = k").
						Having("jj = ?", 1).
						Having("jjj = ?", 2).
						OrderBy("l").
						OrderBy("l").
						OrderBy("l").
						Limit(7).
						Offset(8).
						ToSql()
				}
			},
		},
		{
			name: "pgq",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					pgq.Select("a", "b", "z", "y", "x").
						Distinct().
						From("c").
						Where("d = ? OR e = ?", 1, "wat").
						Where(pgq.Eq{"f": 2, "x": "hi"}).
						Where(map[string]any{"g": 3}).
						Where(pgq.Eq{"h": []int{1, 2, 3}}).
						GroupBy("i").
						GroupBy("ii").
						GroupBy("iii").
						Having("j = k").
						Having("jj = ?", 1).
						Having("jjj = ?", 2).
						OrderBy("l").
						OrderBy("l").
						OrderBy("l").
						Limit(7).
						Offset(8).
						SQL()
				}
			},
		},
		{
			name: "sqrl",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					sqrlDollar.Select("a", "b", "z", "y", "x").
						Distinct().
						From("c").
						Where("d = ? OR e = ?", 1, "wat").
						Where(sqrl.Eq{"f": 2, "x": "hi"}).
						Where(map[string]any{"g": 3}).
						Where(sqrl.Eq{"h": []int{1, 2, 3}}).
						GroupBy("i").
						GroupBy("ii").
						GroupBy("iii").
						Having("j = k").
						Having("jj = ?", 1).
						Having("jjj = ?", 2).
						OrderBy("l").
						OrderBy("l").
						OrderBy("l").
						Limit(7).
						Offset(8).
						ToSql()
				}
			},
		},
	}
	for _, tc := range testCases {
		b.Run(tc.name, tc.f)
	}
}

func BenchmarkSelectSubquery(b *testing.B) {
	testCases := []struct {
		name string
		f    func(b *testing.B)
	}{
		{
			name: "squirrel",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					subSelect := squirrelDollar.Select("id").
						From("tickets").
						Where("subdomain_id = ? and (state = ? or state = ?)",
							1, "open", "spam")

					squirrelDollar.Select("a", "b").
						From("c").
						Distinct().
						Column(squirrel.Alias(subSelect, "subq")).
						Where(squirrel.Eq{"f": 2, "x": "hi"}).
						Where(map[string]any{"g": 3}).
						OrderBy("l").
						OrderBy("l").
						Limit(7).
						Offset(8).
						ToSql()
				}
			},
		},
		{
			name: "pgq",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					subSelect := pgq.Select("id").
						From("tickets").
						Where("subdomain_id = ? and (state = ? or state = ?)",
							1, "open", "spam")

					pgq.Select("a", "b").
						From("c").
						Distinct().
						Column(pgq.Alias{Expr: subSelect, As: "subq"}).
						Where(pgq.Eq{"f": 2, "x": "hi"}).
						Where(map[string]any{"g": 3}).
						OrderBy("l").
						OrderBy("l").
						Limit(7).
						Offset(8).
						SQL()
				}
			},
		},
		{
			name: "sqrl",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					subSelect := sqrlDollar.Select("id").
						From("tickets").
						Where("subdomain_id = ? and (state = ? or state = ?)",
							1, "open", "spam")

					sqrlDollar.Select("a", "b").
						From("c").
						Distinct().
						Column(sqrl.Alias(subSelect, "subq")).
						Where(sqrl.Eq{"f": 2, "x": "hi"}).
						Where(map[string]any{"g": 3}).
						OrderBy("l").
						OrderBy("l").
						Limit(7).
						Offset(8).
						ToSql()
				}
			},
		},
	}
	for _, tc := range testCases {
		b.Run(tc.name, tc.f)
	}
}

func BenchmarkSelectMoreComplex(b *testing.B) {
	testCases := []struct {
		name string
		f    func(b *testing.B)
	}{
		{
			name: "squirrel",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					squirrelDollar.Select("a", "b").
						Prefix("WITH prefix AS ?", 0).
						Distinct().
						Columns("c").
						Column("IF(d IN ("+squirrel.Placeholders(3)+"), 1, 0) as stat_column",
							1, 2, 3).
						Column(squirrel.Expr("a > ?", 100)).
						Column(squirrel.Eq{"b": []int{101, 102, 103}}).
						From("e").
						JoinClause("CROSS JOIN j1").
						Join("j2").
						LeftJoin("j3").
						RightJoin("j4").
						Where("f = ?", 4).
						Where(squirrel.Eq{"g": 5}).
						Where(map[string]any{"h": 6}).
						Where(squirrel.Eq{"i": []int{7, 8, 9}}).
						Where(squirrel.Or{
							squirrel.Expr("j = ?", 10),
							squirrel.And{
								squirrel.Eq{"k": 11},
								squirrel.Expr("true"),
							},
						}).
						GroupBy("l").
						Having("m = n").
						OrderBy("o ASC", "p DESC").
						Limit(12).
						Offset(13).
						Suffix("FETCH FIRST ? ROWS ONLY", 14).
						ToSql()
				}
			},
		},
		{
			name: "pgq",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					pgq.Select("a", "b").
						Prefix("WITH prefix AS ?", 0).
						Distinct().
						Columns("c").
						Column("IF(d IN ("+pgq.Placeholders(3)+"), 1, 0) as stat_column",
							1, 2, 3).
						Column(pgq.Expr("a > ?", 100)).
						Column(pgq.Eq{"b": []int{101, 102, 103}}).
						From("e").
						JoinClause("CROSS JOIN j1").
						Join("j2").
						LeftJoin("j3").
						RightJoin("j4").
						Where("f = ?", 4).
						Where(pgq.Eq{"g": 5}).
						Where(map[string]any{"h": 6}).
						Where(pgq.Eq{"i": []int{7, 8, 9}}).
						Where(pgq.Or{
							pgq.Expr("j = ?", 10),
							pgq.And{
								pgq.Eq{"k": 11},
								pgq.Expr("true"),
							},
						}).
						GroupBy("l").
						Having("m = n").
						OrderBy("o ASC", "p DESC").
						Limit(12).
						Offset(13).
						Suffix("FETCH FIRST ? ROWS ONLY", 14).
						SQL()
				}
			},
		},
		{
			name: "sqrl",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					sqrlDollar.Select("a", "b").
						Prefix("WITH prefix AS ?", 0).
						Distinct().
						Columns("c").
						Column("IF(d IN ("+sqrl.Placeholders(3)+"), 1, 0) as stat_column",
							1, 2, 3).
						Column(sqrl.Expr("a > ?", 100)).
						Column(sqrl.Eq{"b": []int{101, 102, 103}}).
						From("e").
						JoinClause("CROSS JOIN j1").
						Join("j2").
						LeftJoin("j3").
						RightJoin("j4").
						Where("f = ?", 4).
						Where(sqrl.Eq{"g": 5}).
						Where(map[string]any{"h": 6}).
						Where(sqrl.Eq{"i": []int{7, 8, 9}}).
						Where(sqrl.Or{
							sqrl.Expr("j = ?", 10),
							sqrl.And{
								sqrl.Eq{"k": 11},
								sqrl.Expr("true"),
							},
						}).
						GroupBy("l").
						Having("m = n").
						OrderBy("o ASC", "p DESC").
						Limit(12).
						Offset(13).
						Suffix("FETCH FIRST ? ROWS ONLY", 14).
						ToSql()
				}
			},
		},
	}
	for _, tc := range testCases {
		b.Run(tc.name, tc.f)
	}
}

func BenchmarkInsert(b *testing.B) {
	testCases := []struct {
		name string
		f    func(b *testing.B)
	}{
		{
			name: "squirrel",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					squirrelDollar.Insert("mytable").
						Columns("id", "a", "b", "price", "created", "updated").
						Values(1, "test_a", "test_b", 100.05, "2014-01-05", "2015-01-05").
						ToSql()
				}
			},
		},
		{
			name: "pgq",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					pgq.Insert("mytable").
						Columns("id", "a", "b", "price", "created", "updated").
						Values(1, "test_a", "test_b", 100.05, "2014-01-05", "2015-01-05").
						SQL()
				}
			},
		},
		{
			name: "sqrl",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					sqrl.Insert("mytable").
						Columns("id", "a", "b", "price", "created", "updated").
						Values(1, "test_a", "test_b", 100.05, "2014-01-05", "2015-01-05").
						ToSql()
				}
			},
		},
	}
	for _, tc := range testCases {
		b.Run(tc.name, tc.f)
	}
}

func BenchmarkUpdate(b *testing.B) {
	testCases := []struct {
		name string
		f    func(b *testing.B)
	}{
		{
			name: "squirrel",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					squirrelDollar.Update("mytable").
						Set("foo", 1).
						Set("bar", squirrel.Expr("COALESCE(bar, 0) + 1")).
						Set("c", 2).
						Where("id = ?", 9).
						ToSql()
				}
			},
		},
		{
			name: "pgq",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					pgq.Update("mytable").
						Set("foo", 1).
						Set("bar", pgq.Expr("COALESCE(bar, 0) + 1")).
						Set("c", 2).
						Where("id = ?", 9).
						SQL()
				}
			},
		},
		{
			name: "sqrl",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					sqrl.Update("mytable").
						Set("foo", 1).
						Set("bar", sqrl.Expr("COALESCE(bar, 0) + 1")).
						Set("c", 2).
						Where("id = ?", 9).
						Limit(10).
						Offset(20).
						ToSql()
				}
			},
		},
	}
	for _, tc := range testCases {
		b.Run(tc.name, tc.f)
	}
}

func BenchmarkUpdateMap(b *testing.B) {
	testCases := []struct {
		name string
		f    func(b *testing.B)
	}{
		{
			name: "squirrel",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					squirrelDollar.Update("mytable").
						SetMap(map[string]any{
							"b":   1,
							"c":   2,
							"bar": squirrel.Expr("COALESCE(bar, 0) + 1"),
						}).
						Where("id = ?", 9).
						ToSql()
				}
			},
		},
		{
			name: "pgq",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					pgq.Update("mytable").
						SetMap(map[string]any{
							"b":   1,
							"c":   2,
							"bar": pgq.Expr("COALESCE(bar, 0) + 1"),
						}).
						Where("id = ?", 9).
						SQL()
				}
			},
		},
		{
			name: "sqrl",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					sqrl.Update("mytable").
						SetMap(map[string]any{
							"b":   1,
							"c":   2,
							"bar": sqrl.Expr("COALESCE(bar, 0) + 1"),
						}).
						Where("id = ?", 9).
						ToSql()
				}
			},
		},
	}
	for _, tc := range testCases {
		b.Run(tc.name, tc.f)
	}
}

func BenchmarkDelete(b *testing.B) {
	testCases := []struct {
		name string
		f    func(b *testing.B)
	}{
		{
			name: "squirrel",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					squirrelDollar.Delete("test_table").
						Where("b = ?", 1).
						OrderBy("c").
						Limit(2).
						Offset(3).
						ToSql()
				}
			},
		},
		{
			name: "pgq",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					pgq.Delete("test_table").
						Where("b = ?", 1).
						OrderBy("c").
						SQL()
				}
			},
		},
		{
			name: "sqrl",
			f: func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					sqrl.Delete("test_table").
						Where("b = ?", 1).
						OrderBy("c").
						ToSql()
				}
			},
		},
	}
	for _, tc := range testCases {
		b.Run(tc.name, tc.f)
	}
}
