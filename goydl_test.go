package goydl_test

import (
	"testing"

	"strings"

	"time"

	"github.com/BrianAllred/goydl"
)

func TestBoolOption(t *testing.T) {
	youtubeDl := goydl.NewYoutubeDl()
	youtubeDl.Options.ExtractAudio.Value = true
	if !strings.Contains(youtubeDl.Options.OptionsToCliParameters(), "-x") {
		t.Fail()
	}
}

func TestTimeOption(t *testing.T) {
	youtubeDl := goydl.NewYoutubeDl()
	youtubeDl.Options.Date.Value = time.Date(2016, time.March, 19, 0, 0, 0, 0, time.UTC)
	if !strings.Contains(youtubeDl.Options.OptionsToCliParameters(), "--date 20160319") {
		t.Fail()
	}
}

func TestFileSizeRateOption1(t *testing.T) {
	youtubeDl := goydl.NewYoutubeDl()
	youtubeDl.Options.BufferSize.Value = goydl.FileSizeRateFromString("5.5M")
	if !strings.Contains(youtubeDl.Options.OptionsToCliParameters(), "--buffer-size 5.5M") {
		t.Fail()
	}
}

func TestFileSizeRateOption2(t *testing.T) {
	youtubeDl := goydl.NewYoutubeDl()
	youtubeDl.Options.BufferSize.Value = goydl.FileSizeRateFromValues(5.5, 'M')
	if !strings.Contains(youtubeDl.Options.OptionsToCliParameters(), "--buffer-size 5.5M") {
		t.Fail()
	}
}

func TestIntOption(t *testing.T) {
	youtubeDl := goydl.NewYoutubeDl()
	youtubeDl.Options.SocketTimeout.Value = 5
	if !strings.Contains(youtubeDl.Options.OptionsToCliParameters(), "--socket-timeout 5") {
		t.Fail()
	}
}

func TestIntOptionNegativeIsInfinite(t *testing.T) {
	youtubeDl := goydl.NewYoutubeDl()
	youtubeDl.Options.Retries.Value = -1
	if !strings.Contains(youtubeDl.Options.OptionsToCliParameters(), "-R infinite") {
		t.Fail()
	}
}

func TestStringOption(t *testing.T) {
	youtubeDl := goydl.NewYoutubeDl()
	youtubeDl.Options.Username.Value = "testUser"
	if !strings.Contains(youtubeDl.Options.OptionsToCliParameters(), "-u testUser") {
		t.Fail()
	}
}

func TestDownloadInfo(t *testing.T) {
	youtubeDl := goydl.NewYoutubeDl()
	youtubeDl.VideoURL = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
	info, err := youtubeDl.GetInfo()

	if err != nil {
		t.Fail()
	}

	if info.Title != "Rick Astley - Never Gonna Give You Up" {
		t.Fail()
	}
}
