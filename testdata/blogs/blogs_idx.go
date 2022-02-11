// Code generated by gola 0.0.1; DO NOT EDIT.

package blogs

import (
	"strings"

	"github.com/olachat/gola/corelib"
)

type orderBy int

type idxQuery[T any] struct {
	whereSql    string
	orders      []string
	whereParams []interface{}
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
	q.orders = make([]string, len(args))
	for i, arg := range args {
		switch arg {
		case IdAsc:
			q.orders[i] = "id asc"
		case IdDesc:
			q.orders[i] = "id desc"
		case UserIdAsc:
			q.orders[i] = "user_id asc"
		case UserIdDesc:
			q.orders[i] = "user_id desc"
		case SlugAsc:
			q.orders[i] = "slug asc"
		case SlugDesc:
			q.orders[i] = "slug desc"
		case TitleAsc:
			q.orders[i] = "title asc"
		case TitleDesc:
			q.orders[i] = "title desc"
		case CategoryIdAsc:
			q.orders[i] = "category_id asc"
		case CategoryIdDesc:
			q.orders[i] = "category_id desc"
		case IsPinnedAsc:
			q.orders[i] = "is_pinned asc"
		case IsPinnedDesc:
			q.orders[i] = "is_pinned desc"
		case IsVipAsc:
			q.orders[i] = "is_vip asc"
		case IsVipDesc:
			q.orders[i] = "is_vip desc"
		case CountryAsc:
			q.orders[i] = "country asc"
		case CountryDesc:
			q.orders[i] = "country desc"
		case CreatedAtAsc:
			q.orders[i] = "created_at asc"
		case CreatedAtDesc:
			q.orders[i] = "created_at desc"
		case UpdatedAtAsc:
			q.orders[i] = "updated_at asc"
		case UpdatedAtDesc:
			q.orders[i] = "updated_at desc"
		}
	}
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

type orderReadQuery[T any] interface {
	order[T]
	corelib.ReadQuery[T]
}

type iQuery[T any] interface {
	WhereCategoryIdEQ(val int) iQuery1[T]
	WhereCategoryIdIN(vals ...int) iQuery1[T]
	WhereCountryEQ(val string) iQuery4[T]
	WhereCountryIN(vals ...string) iQuery4[T]
	WhereSlugEQ(val string) orderReadQuery[T]
	WhereSlugIN(vals ...string) orderReadQuery[T]
	WhereUserIdEQ(val int) iQuery9[T]
	WhereUserIdIN(vals ...int) iQuery9[T]
	orderReadQuery[T]
}
type iQuery1[T any] interface {
	AndIsPinnedEQ(val int8) iQuery2[T]
	AndIsPinnedIN(vals ...int8) iQuery2[T]
	orderReadQuery[T]
}

type iQuery2[T any] interface {
	AndIsVipEQ(val int8) orderReadQuery[T]
	AndIsVipIN(vals ...int8) orderReadQuery[T]
	orderReadQuery[T]
}

type iQuery4[T any] interface {
	AndCategoryIdEQ(val int) iQuery5[T]
	AndCategoryIdIN(vals ...int) iQuery5[T]
	AndIsVipEQ(val int8) orderReadQuery[T]
	AndIsVipIN(vals ...int8) orderReadQuery[T]
	orderReadQuery[T]
}

type iQuery5[T any] interface {
	AndIsVipEQ(val int8) orderReadQuery[T]
	AndIsVipIN(vals ...int8) orderReadQuery[T]
	orderReadQuery[T]
}

type iQuery9[T any] interface {
	AndIsPinnedEQ(val int8) iQuery10[T]
	AndIsPinnedIN(vals ...int8) iQuery10[T]
	AndIsVipEQ(val int8) orderReadQuery[T]
	AndIsVipIN(vals ...int8) orderReadQuery[T]
	orderReadQuery[T]
}

type iQuery10[T any] interface {
	AndCategoryIdEQ(val int) orderReadQuery[T]
	AndCategoryIdIN(vals ...int) orderReadQuery[T]
	orderReadQuery[T]
}

type idxQuery1[T any] struct {
	*idxQuery[T]
}

func (q *idxQuery1[T]) AndIsPinnedEQ(val int8) iQuery2[T] {
	q.whereSql = " and category_id = ?"
	q.whereParams = append(q.whereParams, val)
	return &idxQuery2[T]{q.idxQuery}
}

func (q *idxQuery1[T]) AndIsPinnedIN(vals ...int8) iQuery2[T] {
	q.whereSql += " and category_id in (" + corelib.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return &idxQuery2[T]{q.idxQuery}
}

type idxQuery2[T any] struct {
	*idxQuery[T]
}

func (q *idxQuery2[T]) AndIsVipEQ(val int8) orderReadQuery[T] {
	q.whereSql = " and is_pinned = ?"
	q.whereParams = append(q.whereParams, val)
	return q.idxQuery
}

func (q *idxQuery2[T]) AndIsVipIN(vals ...int8) orderReadQuery[T] {
	q.whereSql += " and is_pinned in (" + corelib.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return q.idxQuery
}

type idxQuery4[T any] struct {
	*idxQuery[T]
}

func (q *idxQuery4[T]) AndCategoryIdEQ(val int) iQuery5[T] {
	q.whereSql = " and country = ?"
	q.whereParams = append(q.whereParams, val)
	return &idxQuery5[T]{q.idxQuery}
}

func (q *idxQuery4[T]) AndCategoryIdIN(vals ...int) iQuery5[T] {
	q.whereSql += " and country in (" + corelib.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return &idxQuery5[T]{q.idxQuery}
}

func (q *idxQuery4[T]) AndIsVipEQ(val int8) orderReadQuery[T] {
	q.whereSql = " and country = ?"
	q.whereParams = append(q.whereParams, val)
	return q.idxQuery
}

func (q *idxQuery4[T]) AndIsVipIN(vals ...int8) orderReadQuery[T] {
	q.whereSql += " and country in (" + corelib.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return q.idxQuery
}

type idxQuery5[T any] struct {
	*idxQuery[T]
}

func (q *idxQuery5[T]) AndIsVipEQ(val int8) orderReadQuery[T] {
	q.whereSql = " and category_id = ?"
	q.whereParams = append(q.whereParams, val)
	return q.idxQuery
}

func (q *idxQuery5[T]) AndIsVipIN(vals ...int8) orderReadQuery[T] {
	q.whereSql += " and category_id in (" + corelib.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return q.idxQuery
}

type idxQuery9[T any] struct {
	*idxQuery[T]
}

func (q *idxQuery9[T]) AndIsPinnedEQ(val int8) iQuery10[T] {
	q.whereSql = " and user_id = ?"
	q.whereParams = append(q.whereParams, val)
	return &idxQuery10[T]{q.idxQuery}
}

func (q *idxQuery9[T]) AndIsPinnedIN(vals ...int8) iQuery10[T] {
	q.whereSql += " and user_id in (" + corelib.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return &idxQuery10[T]{q.idxQuery}
}

func (q *idxQuery9[T]) AndIsVipEQ(val int8) orderReadQuery[T] {
	q.whereSql = " and user_id = ?"
	q.whereParams = append(q.whereParams, val)
	return q.idxQuery
}

func (q *idxQuery9[T]) AndIsVipIN(vals ...int8) orderReadQuery[T] {
	q.whereSql += " and user_id in (" + corelib.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return q.idxQuery
}

type idxQuery10[T any] struct {
	*idxQuery[T]
}

func (q *idxQuery10[T]) AndCategoryIdEQ(val int) orderReadQuery[T] {
	q.whereSql = " and is_pinned = ?"
	q.whereParams = append(q.whereParams, val)
	return q.idxQuery
}

func (q *idxQuery10[T]) AndCategoryIdIN(vals ...int) orderReadQuery[T] {
	q.whereSql += " and is_pinned in (" + corelib.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return q.idxQuery
}

// Find methods
func Select[T any]() iQuery[T] {
	return new(idxQuery[T])
}

func (q *idxQuery[T]) WhereCategoryIdEQ(val int) iQuery1[T] {
	q.whereSql += " where category_id = ?"
	q.whereParams = append(q.whereParams, val)
	return &idxQuery1[T]{q}
}

func (q *idxQuery[T]) WhereCategoryIdIN(vals ...int) iQuery1[T] {
	q.whereSql = " where category_id in (" + corelib.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return &idxQuery1[T]{q}
}

func (q *idxQuery[T]) WhereCountryEQ(val string) iQuery4[T] {
	q.whereSql += " where country = ?"
	q.whereParams = append(q.whereParams, val)
	return &idxQuery4[T]{q}
}

func (q *idxQuery[T]) WhereCountryIN(vals ...string) iQuery4[T] {
	q.whereSql = " where country in (" + corelib.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return &idxQuery4[T]{q}
}

func (q *idxQuery[T]) WhereSlugEQ(val string) orderReadQuery[T] {
	q.whereSql += " where slug = ?"
	q.whereParams = append(q.whereParams, val)
	return q
}

func (q *idxQuery[T]) WhereSlugIN(vals ...string) orderReadQuery[T] {
	q.whereSql = " where slug in (" + corelib.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return q
}

func (q *idxQuery[T]) WhereUserIdEQ(val int) iQuery9[T] {
	q.whereSql += " where user_id = ?"
	q.whereParams = append(q.whereParams, val)
	return &idxQuery9[T]{q}
}

func (q *idxQuery[T]) WhereUserIdIN(vals ...int) iQuery9[T] {
	q.whereSql = " where user_id in (" + corelib.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return &idxQuery9[T]{q}
}

func (q *idxQuery[T]) GetWhere() (whereSql string, params []interface{}) {
	var orderSql string
	if len(q.orders) > 0 {
		orderSql = " order by " + strings.Join(q.orders, ",")
	}
	return q.whereSql + orderSql, q.whereParams
}
