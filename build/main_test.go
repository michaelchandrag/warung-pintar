package main

import (
   "encoding/json"
   "net/http"
   "net/http/httptest"
   "testing"
   "io"
   "bytes"

   "github.com/stretchr/testify/assert"

   router "github.com/michaelchandrag/warung-pintar/module/router"
   storage "github.com/michaelchandrag/warung-pintar/module/storage"
)
func performRequest(r http.Handler, method, path string, payload io.Reader) *httptest.ResponseRecorder {
   req, _ := http.NewRequest(method, path, payload)
   w := httptest.NewRecorder()
   r.ServeHTTP(w, req)
   return w
}

func TestSend(t *testing.T) {
   type request struct {
      Text     string      `json:"text"`
   }

   payload := request{
      Text: "Testing text.",
   }

   payloadJson, _ := json.Marshal(&payload)
   b := bytes.NewBuffer(payloadJson)

   router := router.SetupRouter()
   w := performRequest(router, "POST", "/send", b)
   assert.Equal(t, http.StatusOK, w.Code)

   type response struct {
      Data        storage.Message      `json:"data"`
      Success     bool                 `json:"success"`
      Message     string               `json:"message"`
      Error       interface{}          `json:"error,omitempty"`
   }

   var result response
   err := json.Unmarshal([]byte(w.Body.String()), &result)

   assert.Nil(t, err)
   assert.NotNil(t, result.Data.ID)
   assert.NotEqual(t, result.Data.ID, 0)
   assert.Equal(t, result.Data.Text, payload.Text)
   assert.Equal(t, result.Success, true)
}

func TestFind(t *testing.T) {
   router := router.SetupRouter()
   w := performRequest(router, "GET", "/messages", nil)
   assert.Equal(t, http.StatusOK, w.Code)

   type response struct {
      Data        []storage.Message    `json:"data"`
      Success     bool                 `json:"success"`
      Message     string               `json:"message"`
      Error       interface{}          `json:"error,omitempty"`
   }

   var result response
   err := json.Unmarshal([]byte(w.Body.String()), &result)

   assert.Nil(t, err)
   assert.NotNil(t, result.Data)
   assert.Equal(t, result.Success, true)
}