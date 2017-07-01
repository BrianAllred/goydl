/*
 * @Author: brian.allred
 * @Date: 2017-06-28 22:04:54
 * @Last Modified by:   brian.allred
 * @Last Modified time: 2017-06-28 22:04:54

 * Copyright (c) 2017-06-28 22:04:54 brian.allred
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
	"bytes"
	"net"
	"strconv"
	"time"
)

// Option interface represents a CLI parameter that can be passed to youtube-dl
// The String method returns what would be passed to youtube-dl
type Option interface {
	OptionString() string
}

type UintOption struct {
	param        string
	Value        uint
	defaultValue uint
}

type StringOption struct {
	param string
	Value string
}

type IpOption struct {
	param string
	Value net.IP
}

type FileSizeRateOption struct {
	param string
	Value fileSizeRate
}

type TimeOption struct {
	param string
	Value time.Time
}
type IntOption struct {
	param        string
	Value        int
	defaultValue int
}

type BoolOption struct {
	param        string
	Value        bool
	defaultValue bool
}

type StringArrayOption struct {
	param  string
	Values []string
}

func (stringArrOpt StringArrayOption) OptionString() string {
	var buffer bytes.Buffer

	if len(stringArrOpt.Values) > 0 {
		for _, value := range stringArrOpt.Values {
			buffer.WriteString(stringArrOpt.param + " " + value)
		}
	}

	return buffer.String()
}

func (uintOpt UintOption) OptionString() string {
	if uintOpt.Value != uintOpt.defaultValue {
		return uintOpt.param + " " + strconv.FormatUint(uint64(uintOpt.Value), 10)
	}

	return ""
}

func (stringOpt StringOption) OptionString() string {
	if stringOpt.Value != "" {
		return stringOpt.param + " " + stringOpt.Value
	}

	return ""
}

func (ipOpt IpOption) OptionString() string {
	if ipOpt.Value != nil {
		return ipOpt.param + " " + ipOpt.Value.String()
	}

	return ""
}

func (fsrOpt FileSizeRateOption) OptionString() string {
	if !fsrOpt.Value.IsNil() {
		return fsrOpt.param + " " + fsrOpt.Value.String()
	}

	return ""
}

func (timeOpt TimeOption) OptionString() string {
	if !timeOpt.Value.IsZero() {
		// youtube-dl date format is YYYYMMDD
		return timeOpt.param + " " + timeOpt.Value.Format("20060102")
	}

	return ""
}

func (boolOpt BoolOption) OptionString() string {
	if boolOpt.Value != boolOpt.defaultValue {
		return boolOpt.param
	}

	return ""
}

func (intOpt IntOption) OptionString() string {
	if intOpt.Value != intOpt.defaultValue {
		// For IntOptions, negative values represent "infinite"
		if intOpt.Value < 0 {
			return intOpt.param + " infinite"
		}

		return intOpt.param + " " + strconv.Itoa(intOpt.Value)
	}

	return ""
}
