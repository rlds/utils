package utils_test

import (
	"testing"

	. "../utils"
)

func TestIsIP(t *testing.T) {
	if !IsIP("127.0.0.1") {
		t.Error("should be true")
	}
	if IsIP("257.0.1") {
		t.Error("should be false")
	}
}

func TestIsEmail(t *testing.T) {
	if !IsEmail("jinshan@wacai.com") {
		t.Error("should be true")
	}
	if IsEmail("jinshan") {
		t.Error("should be false")
	}
}

func TestIsTelephone(t *testing.T) {
	if !IsTelephone("15657584405") {
		t.Error("should be true")
	}
	if IsTelephone("123456789") {
		t.Error("should be false")
	}
}
