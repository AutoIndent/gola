// Code generated by gola 0.0.1; DO NOT EDIT.

package blogs

import "github.com/olachat/gola/corelib"

type orderBy int

type idxQuery[T any] struct {
}

// order by enum & interface
const (
	IdAsc orderBy = iota
	IdDesc
	UserIdAsc
	UserIdDesc
	SlugAsc
	SlugDesc
	TitleAsc
	TitleDesc
	CategoryIdAsc
	CategoryIdDesc
	IsPinnedAsc
	IsPinnedDesc
	IsVipAsc
	IsVipDesc
	CountryAsc
	CountryDesc
	CreatedAtAsc
	CreatedAtDesc
	UpdatedAtAsc
	UpdatedAtDesc
)

func (q *idxQuery[T]) OrderBy(args ...orderBy) corelib.ReadQuery[T] {
	return q
}

func (q *idxQuery[T]) All() []*T {
	return nil
}

func (q *idxQuery[T]) Limit(limit, offset int) []*T {
	return nil
}

type order[T any] interface {
	OrderBy(args ...orderBy) corelib.ReadQuery[T]
}

type iQuery1[T any] interface {
	WhereCategoryIdEQ(val int) *idxQuery8[T]
	WhereCategoryIdIN(vals ...int) *idxQuery8[T]
	WhereCountryEQ(val string) *idxQuery1[T]
	WhereCountryIN(vals ...string) *idxQuery1[T]
	WhereSlugEQ(val string) *idxQuery3[T]
	WhereSlugIN(vals ...string) *idxQuery3[T]
	WhereUserIdEQ(val int) *idxQuery4[T]
	WhereUserIdIN(vals ...int) *idxQuery4[T]
	order[T]
	corelib.ReadQuery[T]
}

// Find methods
func Select[T any]() *idxQuery[T] {
	return new(idxQuery[T])
}
func (q *idxQuery[T]) WhereCategoryIdEQ(val int) *idxQuery8[T] {
	return q
}

func (q *idxQuery[T]) WhereCategoryIdIN(vals ...int) *idxQuery8[T] {
	return q
}

func (q *idxQuery[T]) WhereCountryEQ(val string) *idxQuery1[T] {
	return q
}

func (q *idxQuery[T]) WhereCountryIN(vals ...string) *idxQuery1[T] {
	return q
}

func (q *idxQuery[T]) WhereSlugEQ(val string) *idxQuery3[T] {
	return q
}

func (q *idxQuery[T]) WhereSlugIN(vals ...string) *idxQuery3[T] {
	return q
}

func (q *idxQuery[T]) WhereUserIdEQ(val int) *idxQuery4[T] {
	return q
}

func (q *idxQuery[T]) WhereUserIdIN(vals ...int) *idxQuery4[T] {
	return q
}
