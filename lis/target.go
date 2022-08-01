package lis

import (
	"strings"
)

type Target struct {
	key string
	ty  string // type
}

func NewTarget(key, ty string) Target {
	return Target{key, ty}
}

func (me *Target) String() string {
	if me == nil {
		return ""
	}
	var str strings.Builder
	str.Grow(256)
	str.WriteString("[Target(Key:")
	str.WriteString(me.key)
	str.WriteString(",Type:")
	str.WriteString(me.ty)
	str.WriteString(")]")
	return str.String()
}

func (me *Target) Key() string {
	if me == nil {
		return ""
	}
	return me.key
}

func (me *Target) Type() string {
	if me == nil {
		return ""
	}
	return me.ty
}
