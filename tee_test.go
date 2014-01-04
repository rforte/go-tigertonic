package tigertonic

import (
	"net/http"
	"net/url"
	"testing"
)

func TestTeeResponseWriter(t *testing.T) {
	w0 := &testResponseWriter{}
	w := NewTeeResponseWriter(w0)
	r, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	Marshaled(func(*url.URL, http.Header) (int, http.Header, *testResponse, error) {
		return http.StatusOK, http.Header{"X-Foo": []string{"bar"}}, &testResponse{"bar"}, nil
	}).ServeHTTP(w, r)
	if w0.Status != w.Status {
		t.Fatal(w.Status)
	}
	if w0.Header().Get("X-Foo") != w.Header().Get("X-Foo") {
		t.Fatal(w.Status)
	}
	if w0.Body.String() != w.Body.String() {
		t.Fatal(w.Body.String())
	}
}
