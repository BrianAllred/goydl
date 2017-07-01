/*
 * @Author: brian.allred
 * @Date: 2017-06-28 22:04:28
 * @Last Modified by:   brian.allred
 * @Last Modified time: 2017-06-28 22:04:28

 * Copyright (c) 2017-06-28 22:04:28 brian.allred
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

// Info is an object representing the JSON returned from a -J download (dump-single-json)
type Info struct {
	ID          string      `json:"id"`
	Uploader    string      `json:"uploader"`
	UploaderID  string      `json:"uploader_id"`
	UploaderURL string      `json:"uploader_url"`
	UploadDate  string      `json:"upload_date"`
	License     string      `json:"license"`
	Creator     interface{} `json:"creator"`
	Title       string      `json:"title"`
	AltTitle    interface{} `json:"alt_title"`
	Thumbnail   string      `json:"thumbnail"`
	Description string      `json:"description"`
	Categories  []string    `json:"categories"`
	Tags        []string    `json:"tags"`
	Subtitles   struct {
	} `json:"subtitles"`
	AutomaticCaptions struct {
	} `json:"automatic_captions"`
	Duration      int         `json:"duration"`
	AgeLimit      int         `json:"age_limit"`
	Annotations   interface{} `json:"annotations"`
	Chapters      interface{} `json:"chapters"`
	WebpageURL    string      `json:"webpage_url"`
	ViewCount     int         `json:"view_count"`
	LikeCount     int         `json:"like_count"`
	DislikeCount  int         `json:"dislike_count"`
	AverageRating float64     `json:"average_rating"`
	Formats       []struct {
		Ext         string      `json:"ext"`
		FormatNote  string      `json:"format_note"`
		Acodec      string      `json:"acodec"`
		Abr         int         `json:"abr,omitempty"`
		Container   string      `json:"container,omitempty"`
		FormatID    string      `json:"format_id"`
		URL         string      `json:"url"`
		ManifestURL string      `json:"manifest_url,omitempty"`
		Width       interface{} `json:"width,omitempty"`
		Height      interface{} `json:"height,omitempty"`
		Tbr         float64     `json:"tbr,omitempty"`
		Asr         int         `json:"asr,omitempty"`
		Fps         interface{} `json:"fps,omitempty"`
		Language    interface{} `json:"language,omitempty"`
		Filesize    int         `json:"filesize,omitempty"`
		Vcodec      string      `json:"vcodec"`
		Format      string      `json:"format"`
		Protocol    string      `json:"protocol"`
		HTTPHeaders struct {
			UserAgent      string `json:"User-Agent"`
			AcceptCharset  string `json:"Accept-Charset"`
			Accept         string `json:"Accept"`
			AcceptEncoding string `json:"Accept-Encoding"`
			AcceptLanguage string `json:"Accept-Language"`
		} `json:"http_headers"`
		PlayerURL  string `json:"player_url,omitempty"`
		Resolution string `json:"resolution,omitempty"`
	} `json:"formats"`
	IsLive             interface{} `json:"is_live"`
	StartTime          interface{} `json:"start_time"`
	EndTime            interface{} `json:"end_time"`
	Series             interface{} `json:"series"`
	SeasonNumber       interface{} `json:"season_number"`
	EpisodeNumber      interface{} `json:"episode_number"`
	Extractor          string      `json:"extractor"`
	WebpageURLBasename string      `json:"webpage_url_basename"`
	ExtractorKey       string      `json:"extractor_key"`
	Playlist           interface{} `json:"playlist"`
	PlaylistIndex      interface{} `json:"playlist_index"`
	Thumbnails         []struct {
		URL string `json:"url"`
		ID  string `json:"id"`
	} `json:"thumbnails"`
	DisplayID          string      `json:"display_id"`
	RequestedSubtitles interface{} `json:"requested_subtitles"`
	RequestedFormats   []struct {
		Ext         string      `json:"ext"`
		Height      int         `json:"height,omitempty"`
		FormatNote  string      `json:"format_note"`
		Vcodec      string      `json:"vcodec"`
		FormatID    string      `json:"format_id"`
		URL         string      `json:"url"`
		ManifestURL string      `json:"manifest_url,omitempty"`
		Width       int         `json:"width,omitempty"`
		Tbr         float64     `json:"tbr"`
		Asr         interface{} `json:"asr,omitempty"`
		Fps         int         `json:"fps,omitempty"`
		Language    interface{} `json:"language,omitempty"`
		Filesize    int         `json:"filesize"`
		Acodec      string      `json:"acodec"`
		Format      string      `json:"format"`
		Protocol    string      `json:"protocol"`
		HTTPHeaders struct {
			UserAgent      string `json:"User-Agent"`
			AcceptCharset  string `json:"Accept-Charset"`
			Accept         string `json:"Accept"`
			AcceptEncoding string `json:"Accept-Encoding"`
			AcceptLanguage string `json:"Accept-Language"`
		} `json:"http_headers"`
		PlayerURL string `json:"player_url,omitempty"`
		Abr       int    `json:"abr,omitempty"`
	} `json:"requested_formats"`
	Format         string      `json:"format"`
	FormatID       string      `json:"format_id"`
	Width          int         `json:"width"`
	Height         int         `json:"height"`
	Resolution     interface{} `json:"resolution"`
	Fps            int         `json:"fps"`
	Vcodec         string      `json:"vcodec"`
	Vbr            interface{} `json:"vbr"`
	StretchedRatio interface{} `json:"stretched_ratio"`
	Acodec         string      `json:"acodec"`
	Abr            int         `json:"abr"`
	Ext            string      `json:"ext"`
	Fulltitle      string      `json:"fulltitle"`
	Filename       string      `json:"_filename"`
}
