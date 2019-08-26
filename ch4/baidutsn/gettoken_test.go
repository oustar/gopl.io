package baidutsn

import (
	"testing"
)

func TestGetToken(t *testing.T) {
	key, err := GetToken()
	if err == nil {
		if len(key) != 70 {
			t.Errorf("the length of token is wrong,%d\n", len(key))
		}
	}
}
