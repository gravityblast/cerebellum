package main

import (
  "net/http"
  "net/http/httptest"
)

func init() {
  router.RequestLogFunc = func(int, *http.Request) {}
}

func newTestRequest(method, path string) *httptest.ResponseRecorder  {
  request, _ := http.NewRequest(method, path, nil)
  recorder := httptest.NewRecorder()
  router.ServeHTTP(recorder, request)

  return recorder
}
