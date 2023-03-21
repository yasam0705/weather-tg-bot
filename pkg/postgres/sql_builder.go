package postgres

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type builder struct {
	*goqu.DialectWrapper
}

func newBuilder() *builder {
	d := goqu.Dialect("postgres")
	return &builder{DialectWrapper: &d}
}

func (b *builder) Equal(k string, v interface{}) exp.BooleanExpression {
	return goqu.C(k).Eq(v)
}

func (b *builder) OrderByDesc(k string) exp.OrderedExpression {
	return goqu.I(k).Desc()
}

func (b *builder) OrderByAsc(k string) exp.OrderedExpression {
	return goqu.I(k).Asc()
}

func (b *builder) Count(k interface{}) exp.SQLFunctionExpression {
	return goqu.COUNT(k)
}
