package types

import (
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	tagId          = "watson"
	attrAlwaysOmit = "-"
	attrOmitEmpty  = "omitempty"
	attrInline     = "inline"
)

type tag struct {
	name       string
	f          *reflect.StructField
	omitempty  bool
	alwaysomit bool
	inline     bool
}

func findField(key string, obj reflect.Value) (*tag, bool) {
	t := obj.Type()
	size := obj.NumField()
	for i := 0; i < size; i++ {
		f := t.Field(i)
		tag := parseTag(&f)
		if tag.Key() == key {
			return tag, true
		}
	}
	return nil, false
}

func inlineFields(obj reflect.Value) []*tag {
	t := obj.Type()
	tags := make([]*tag, 0)
	size := obj.NumField()
	for i := 0; i < size; i++ {
		f := t.Field(i)
		tag := parseTag(&f)
		if tag.Inline() {
			tags = append(tags, tag)
		}
	}
	return tags
}

func parseTag(f *reflect.StructField) *tag {
	tag := &tag{f: f}
	name := f.Tag.Get(tagId)
	if name == "" {
		return tag
	}
	attrs := strings.Split(name, ",")
	first := true
	for _, attr := range attrs {
		if first {
			if attr == attrAlwaysOmit {
				tag.alwaysomit = true
			} else {
				tag.name = attr
			}
			first = false
			continue
		}
		switch attr {
		case attrOmitEmpty:
			tag.omitempty = true
		case attrInline:
			tag.inline = true
		}
	}
	return tag
}

func (t *tag) Key() string {
	if t.name == "" {
		return strings.ToLower(t.f.Name)
	}
	return t.name
}

func (t *tag) ShouldAlwaysOmit() bool {
	if t.alwaysomit {
		return true
	}
	r, _ := utf8.DecodeRuneInString(t.f.Name)
	return unicode.IsLower(r)
}

func (t *tag) OmitEmpty() bool {
	return t.omitempty
}

func (t *tag) Inline() bool {
	return t.inline
}

func (t *tag) FieldOf(v reflect.Value) reflect.Value {
	return v.FieldByIndex(t.f.Index)
}
