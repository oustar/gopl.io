package baidutsn

import (
	"io/ioutil"
	"testing"
)

func TestTransfer(t *testing.T) {
	audio, err := Transfer("我很好")
	if err == nil {
		err = ioutil.WriteFile("test.mp3", audio, 0644)
		if err != nil {
			t.Error("write file error")
		}
	} else {
		t.Errorf("transfer error %v", err)
	}
}
