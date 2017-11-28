/*
 * @Author: brian.allred
 * @Date: 2017-06-28 22:05:04
 * @Last Modified by:   brian.allred
 * @Last Modified time: 2017-06-28 22:05:04

 * Copyright (c) 2017-06-28 22:05:04 brian.allred
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package goydl

import (
	"encoding/json"
	"errors"
	"io"
	"os/exec"
	"strings"
)

// Represents the youtube-dl object
type youtubeDl struct {
	VideoURL       string
	YoutubeDlPath  string
	Info           Info
	Stderr, Stdout io.ReadCloser
	Options        options
	Err            error
	cmd            *exec.Cmd
}

// NewYoutubeDl returns a newly instantiated youtubeDl object
func NewYoutubeDl() youtubeDl {
	return youtubeDl{YoutubeDlPath: "youtube-dl", Info: Info{}, Options: NewOptions(), cmd: new(exec.Cmd)}
}

// Download will download either the given url or the youtubeDl.VideoUrl from youtube
func (ydl *youtubeDl) Download(urls ...string) (*exec.Cmd, error) {
	if ydl.Options.Update.Value {
		return ydl.Update()
	}

	if len(urls) > 1 {
		return nil, errors.New("invalid argument")
	}

	if urls != nil {
		ydl.VideoURL = urls[0]
	}

	// Get the info first
	if _, err := ydl.GetInfo(); err != nil {
		return nil, err
	}

	ydl.cmd = exec.Command(ydl.YoutubeDlPath, ydl.VideoURL)

	ydl.cmd.Args = append(ydl.cmd.Args, strings.Fields(ydl.Options.OptionsToCliParameters())...)

	// Setup stderr pipe
	ydl.Stderr, ydl.Err = ydl.cmd.StderrPipe()

	if ydl.Err != nil {
		return nil, ydl.Err
	}

	// Setup stdout pipe
	ydl.Stdout, ydl.Err = ydl.cmd.StdoutPipe()

	if ydl.Err != nil {
		return nil, ydl.Err
	}

	// Return the command and any error from starting the command
	return ydl.cmd, ydl.cmd.Start()
}

// GetInfo gets the info of a particular video or playlist
func (ydl *youtubeDl) GetInfo() (Info, error) {
	// Setup command with '-J' argument
	cmd := exec.Command(ydl.YoutubeDlPath, "-J", ydl.VideoURL)
	stdOut, err := cmd.StdoutPipe()

	if err != nil {
		return Info{}, err
	}

	if err := cmd.Start(); err != nil {
		return Info{}, err
	}

	if err := json.NewDecoder(stdOut).Decode(&ydl.Info); err != nil {
		return Info{}, err
	}

	return ydl.Info, cmd.Wait()
}

// Update updates the youtube-dl binary
func (ydl *youtubeDl) Update() (*exec.Cmd, error) {
	ydl.cmd = exec.Command(ydl.YoutubeDlPath, "-U")

	ydl.Stderr, ydl.Err = ydl.cmd.StderrPipe()

	if ydl.Err != nil {
		return nil, ydl.Err
	}

	ydl.Stdout, ydl.Err = ydl.cmd.StdoutPipe()

	if ydl.Err != nil {
		return nil, ydl.Err
	}

	return ydl.cmd, ydl.cmd.Start()
}
