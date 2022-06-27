package pkg

import (
	"fyne.io/fyne/v2/widget"
	"testing"
)

func TestWaterMarkWithVideo(t *testing.T) {
	WaterMark("https://v.douyin.com/Ya2aQpp/", "/Users/liutao/Desktop", widget.NewLabel(""))
}

func TestWaterMarkWithImage(t *testing.T) {
	WaterMark("https://v.douyin.com/Ya2crCd/", "/Users/liutao/Desktop", widget.NewLabel(""))
}
