package context

// CtxKey used for add context key
type CtxKey string

func (c CtxKey) String() string {
	return "ADV_context_key:" + string(c)
}
