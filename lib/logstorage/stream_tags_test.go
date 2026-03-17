package logstorage

import (
	"testing"
)

func TestStreamTagsUnmarshalString_Success(t *testing.T) {
	f := func(s string) {
		t.Helper()

		var st StreamTags
		if err := st.unmarshalString(s); err != nil {
			t.Fatalf("unexpected error in unmarshalString(%s): %s", s, err)
		}
		result := st.String()
		if result != s {
			t.Fatalf("unexpected result\ngot\n%s\nwant\n%s", result, s)
		}
	}

	f(`{}`)
	f(`{foo="bar"}`)
	f(`{a="b",c="d"}`)
}

func TestStreamTagsUnmarshalString_Failure(t *testing.T) {
	f := func(s string) {
		t.Helper()

		var st StreamTags
		if err := st.unmarshalString(s); err == nil {
			t.Fatalf("expecting non-nil error in unmarshalString(%s)", s)
		}
	}

	f(``)
	f(`{`)
	f(`{foo}`)
	f(`{"foo":"bar"}`)
	f(`{foo=abc`)
	f(`{foo="abc`)
	f(`{foo="abc"`)
	f(`{foo="abc",`)
	f(`{foo="abc",bar}`)
}
