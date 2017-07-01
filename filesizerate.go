/*
 * @Author: brian.allred
 * @Date: 2017-06-28 22:03:59
 * @Last Modified by:   brian.allred
 * @Last Modified time: 2017-06-28 22:03:59

 * Copyright (c) 2017-06-28 22:03:59 brian.allred
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
	"log"
	"strconv"
)

// Object representing a filesize or filerate (eg. 4K, 10M, etc)
type fileSizeRate struct {
	SizeRate float64
	Unit     rune
}

// FileSizeRateFromString creates a new FileSizeRate from a string
func FileSizeRateFromString(fileSizeRateString string) fileSizeRate {
	sizeRate, err := strconv.ParseFloat(fileSizeRateString[0:len(fileSizeRateString)-1], 64)

	if err != nil {
		log.Fatal(err)
	}

	var unit rune

	for index, runeValue := range fileSizeRateString {
		if index == len(fileSizeRateString)-1 {
			unit = runeValue
		}
	}

	return fileSizeRate{sizeRate, unit}
}

// FileSizeRateFromValues creates a new FileSizeRate from values (float64, rune)
func FileSizeRateFromValues(sizeRate float64, unit rune) fileSizeRate {
	return fileSizeRate{sizeRate, unit}
}

// Return the string representation of a FileSizeRate
func (fsr fileSizeRate) String() string {
	return strconv.FormatFloat(fsr.SizeRate, 'f', -1, 64) + string(fsr.Unit)
}

// Returns if the given FileSizeRate is nil/invalid
func (fsr fileSizeRate) IsNil() bool {
	return fsr.SizeRate == 0 || fsr.Unit == 0
}
