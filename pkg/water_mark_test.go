package pkg

import "testing"

func TestWaterMarkWithVideo(t *testing.T) {
	WaterMark("https://v.douyin.com/Ya2aQpp/", "/Users/liutao/Desktop")
}

func TestWaterMarkWithImage(t *testing.T) {
	WaterMark("https://v.douyin.com/Ya2crCd/", "/Users/liutao/Desktop")
}
