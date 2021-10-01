// не до конца понятная штука

// внутренний(можно сказать локальный, не публичный) тест, для тестирования тестового сайта /Hello
package serverwithapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	// библиотека для проверки(в данном случае проверяет, находится ли по тому адресу слово "Hello")
	// при полож ответе выводится "ok"
	"github.com/stretchr/testify/assert"
)

// функция для тестирования
func TestServerAPI_HandleHello(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}
