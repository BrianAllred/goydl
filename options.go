/*
 * @Author: brian.allred
 * @Date: 2017-06-28 22:04:57
 * @Last Modified by:   brian.allred
 * @Last Modified time: 2017-06-28 22:04:57

 * Copyright (c) 2017-06-28 22:04:57 brian.allred
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
	"reflect"
	"strings"
)

// Object representing all of the options available for a youtube-dl download
type options struct {
	// General
	Update                BoolOption
	IgnoreErrors          BoolOption
	AbortOnError          BoolOption
	DumpUserAgent         BoolOption
	ListExtractors        BoolOption
	ExtractorDescriptions BoolOption
	ForceGenericExtractor BoolOption
	DefaultSearch         StringOption
	IgnoreConfig          BoolOption
	ConfigLocation        StringOption
	FlatPlaylist          BoolOption
	MarkWatched           BoolOption
	NoMarkWatched         BoolOption
	NoColor               BoolOption

	// Network
	Proxy         StringOption
	SocketTimeout UintOption
	SourceAddress IpOption
	ForceIpv4     BoolOption
	ForceIpv6     BoolOption

	// Geo Restriction
	GeoVerificationProxy StringOption

	// Geo Restriction (Experimental)
	GeoBypass        BoolOption
	NoGeoBypass      BoolOption
	GeoBypassCountry StringOption

	// Video Selection
	PlaylistStart   IntOption
	PlaylistEnd     IntOption
	PlaylistItems   StringOption
	MatchTitle      StringOption
	RejectTitle     StringOption
	MaxDownloads    UintOption
	MinFilesize     FileSizeRateOption
	MaxFilesize     FileSizeRateOption
	Date            TimeOption
	DateBefore      TimeOption
	DateAfter       TimeOption
	MinViews        UintOption
	MaxViews        UintOption
	MatchFilter     StringOption
	NoPlaylist      BoolOption
	YesPlaylist     BoolOption
	AgeLimit        UintOption
	DownloadArchive StringOption

	// Video Selection (Experimental)
	IncludeAds BoolOption

	// Download Options
	LimitRate                   FileSizeRateOption
	Retries                     IntOption
	FragmentRetries             IntOption
	SkipUnavailableFragments    BoolOption
	AbortOnUnavailableFragments BoolOption
	KeepFragments               BoolOption
	BufferSize                  FileSizeRateOption
	NoResizeBuffer              BoolOption
	PlaylistReverse             BoolOption
	PlaylistRandom              BoolOption
	HlsPreferNative             BoolOption
	HlsPreferFfmpeg             BoolOption
	HlsUseMpegTs                BoolOption
	ExternalDownloader          StringOption
	ExternalDownloaderArgs      StringOption

	// Download Options (Experimental)
	XAttrSetFilesize BoolOption

	// Filesystem Options
	BatchFile         StringOption
	ID                BoolOption
	Output            StringOption
	AutonumberStart   UintOption
	RestrictFilenames BoolOption
	NoOverwrites      BoolOption
	Continue          BoolOption
	NoContinue        BoolOption
	NoPart            BoolOption
	NoMtime           BoolOption
	WriteDescription  BoolOption
	WriteInfoJSON     BoolOption
	WriteAnnotations  BoolOption
	LoadInfoJSON      StringOption
	Cookies           StringOption
	CacheDir          StringOption
	NoCacheDir        BoolOption
	RmCacheDir        BoolOption

	// Thumbnail Images Options
	WriteThumbnail     BoolOption
	WriteAllThumbnails BoolOption
	ListThumbnails     BoolOption

	// Verbosity / Simulation Options
	Quiet          BoolOption
	NoWarnings     BoolOption
	Simulate       BoolOption
	SkipDownload   BoolOption
	GetURL         BoolOption
	GetTitle       BoolOption
	GetID          BoolOption
	GetThumbnail   BoolOption
	GetDescription BoolOption
	GetDuration    BoolOption
	GetFilename    BoolOption
	GetFormat      BoolOption
	DumpJSON       BoolOption
	DumpSingleJSON BoolOption
	PrintJSON      BoolOption
	Newline        BoolOption
	NoProgress     BoolOption
	ConsoleTitle   BoolOption
	Verbose        BoolOption
	DumpPages      BoolOption
	WritePages     BoolOption
	PrintTraffic   BoolOption
	CallHome       BoolOption
	NoCallHome     BoolOption

	// Workaround Options
	NoCheckCertificate BoolOption
	PreferInsecure     BoolOption
	UserAgent          StringOption
	Referer            StringOption
	Headers            StringArrayOption
	BiDiWorkaround     BoolOption
	SleepInterval      UintOption
	MaxSleepInterval   UintOption

	// Workaround Options (Experimental)
	Encoding StringOption

	// Video Format Options
	Format                  StringOption
	AllFormats              BoolOption
	PreferFreeFormats       BoolOption
	ListFormats             BoolOption
	YoutubeSkipDashManifest BoolOption
	MergeOutputFormat       StringOption

	// Subtitle Options
	WriteSub     BoolOption
	WriteAutoSub BoolOption
	AllSubs      BoolOption
	ListSubs     BoolOption
	SubFormat    StringOption
	SubLang      StringOption

	// Authentication Options
	Username      StringOption
	Password      StringOption
	TwoFactor     StringOption
	NetRc         BoolOption
	VideoPassword StringOption

	// Adobe Pass Options
	ApMso      StringOption
	ApUsername StringOption
	ApPassword StringOption
	ApListMso  BoolOption

	// Post-processing Options
	ExtractAudio      BoolOption
	AudioFormat       StringOption
	AudioQuality      StringOption
	RecodeVideo       StringOption
	PostprocessorArgs StringOption
	KeepVideo         BoolOption
	NoPostOverwrites  BoolOption
	EmbedSubs         BoolOption
	EmbedThumbnail    BoolOption
	AddMetadata       BoolOption
	MetadataFromTitle StringOption
	XAttrs            BoolOption
	Fixup             StringOption
	PreferAvconv      BoolOption
	PreferFfmpeg      BoolOption
	FfmpegLocation    StringOption
	Exec              StringOption
	ConvertSubs       StringOption
}

// Returns a new options object with the various options initialized with their params
func NewOptions() options {
	opts := options{}

	// General options
	opts.Update = BoolOption{param: "-U"}
	opts.IgnoreErrors = BoolOption{param: "-i"}
	opts.AbortOnError = BoolOption{param: "--abort-on-error"}
	opts.DumpUserAgent = BoolOption{param: "--dump-user-agent"}
	opts.ListExtractors = BoolOption{param: "--list-extractors"}
	opts.ExtractorDescriptions = BoolOption{param: "--extractor-descriptions"}
	opts.ForceGenericExtractor = BoolOption{param: "--force-generic-extractor"}
	opts.DefaultSearch = StringOption{param: "--default-search"}
	opts.IgnoreConfig = BoolOption{param: "--ignore-config"}
	opts.ConfigLocation = StringOption{param: "--config-location"}
	opts.FlatPlaylist = BoolOption{param: "--flat-playlist"}
	opts.MarkWatched = BoolOption{param: "--mark-watched"}
	opts.NoMarkWatched = BoolOption{param: "--no-mark-watched"}
	opts.NoColor = BoolOption{param: "--no-color"}

	// Network options
	opts.Proxy = StringOption{param: "--proxy"}
	opts.SocketTimeout = UintOption{param: "--socket-timeout"}
	opts.SourceAddress = IpOption{param: "--source-address"}
	opts.ForceIpv4 = BoolOption{param: "-4"}
	opts.ForceIpv6 = BoolOption{param: "-6"}

	// Geo Restriction options
	opts.GeoVerificationProxy = StringOption{param: "--geo-verification-proxy"}
	opts.GeoBypass = BoolOption{param: "--geo-bypass"}
	opts.NoGeoBypass = BoolOption{param: "--no-geo-bypass"}
	opts.GeoBypassCountry = StringOption{param: "--geo-bypass-country"}

	// Video Selection options
	opts.PlaylistStart = IntOption{param: "--playlist-start", Value: 1, defaultValue: 1}
	opts.PlaylistEnd = IntOption{param: "--playlist-end", Value: -1, defaultValue: -1}
	opts.PlaylistItems = StringOption{param: "--playlist-items"}
	opts.MatchTitle = StringOption{param: "--match-title"}
	opts.RejectTitle = StringOption{param: "--reject-title"}
	opts.MaxDownloads = UintOption{param: "--max-downloads", Value: 0, defaultValue: 0}
	opts.MinFilesize = FileSizeRateOption{param: "--min-filesize"}
	opts.MaxFilesize = FileSizeRateOption{param: "--max-filesize"}
	opts.Date = TimeOption{param: "--date"}
	opts.DateBefore = TimeOption{param: "--datebefore"}
	opts.DateAfter = TimeOption{param: "--dateafter"}
	opts.MinViews = UintOption{param: "--min-views"}
	opts.MaxViews = UintOption{param: "--max-views"}
	opts.MatchFilter = StringOption{param: "--match-filter"}
	opts.NoPlaylist = BoolOption{param: "--no-playlist"}
	opts.YesPlaylist = BoolOption{param: "--yes-playlist"}
	opts.AgeLimit = UintOption{param: "--age-limit"}
	opts.DownloadArchive = StringOption{param: "--download-archive"}
	opts.IncludeAds = BoolOption{param: "--include-ads"}

	// Download options
	opts.LimitRate = FileSizeRateOption{param: "-r"}
	opts.Retries = IntOption{param: "-R", Value: 10, defaultValue: 10}
	opts.FragmentRetries = IntOption{param: "--fragment-retries", Value: 10, defaultValue: 10}
	opts.SkipUnavailableFragments = BoolOption{param: "--skip-unavailable-fragments"}
	opts.AbortOnUnavailableFragments = BoolOption{param: "--abort-on-unavailable-fragments"}
	opts.KeepFragments = BoolOption{param: "--keep-fragments"}
	opts.BufferSize = FileSizeRateOption{param: "--buffer-size"}
	opts.NoResizeBuffer = BoolOption{param: "--no-resize-buffer"}
	opts.PlaylistReverse = BoolOption{param: "--playlist-reverse"}
	opts.PlaylistRandom = BoolOption{param: "--playlist-random"}
	opts.XAttrSetFilesize = BoolOption{param: "--xattr-set-filesize"}
	opts.HlsPreferNative = BoolOption{param: "--hls-prefer-native"}
	opts.HlsPreferFfmpeg = BoolOption{param: "--hls-prefer-ffmpeg"}
	opts.HlsUseMpegTs = BoolOption{param: "--hls-use-mpegts"}
	opts.ExternalDownloader = StringOption{param: "--external-downloader"}
	opts.ExternalDownloaderArgs = StringOption{param: "--external-downloader-args"}

	// Filesystem options
	opts.BatchFile = StringOption{param: "-a"}
	opts.ID = BoolOption{param: "--id"}
	opts.Output = StringOption{param: "-o"}
	opts.AutonumberStart = UintOption{param: "--autonumber-start", Value: 1, defaultValue: 1}
	opts.RestrictFilenames = BoolOption{param: "--restrict-filenames"}
	opts.NoOverwrites = BoolOption{param: "-w"}
	opts.Continue = BoolOption{param: "-c"}
	opts.NoContinue = BoolOption{param: "--no-continue"}
	opts.NoPart = BoolOption{param: "--no-part"}
	opts.NoMtime = BoolOption{param: "--no-mtime"}
	opts.WriteDescription = BoolOption{param: "--write-description"}
	opts.WriteInfoJSON = BoolOption{param: "--write-info-json"}
	opts.WriteAnnotations = BoolOption{param: "--write-annotations"}
	opts.LoadInfoJSON = StringOption{param: "--load-info-json"}
	opts.Cookies = StringOption{param: "--cookies"}
	opts.CacheDir = StringOption{param: "--cache-dir"}
	opts.NoCacheDir = BoolOption{param: "--no-cache-dir"}
	opts.RmCacheDir = BoolOption{param: "--rm-cache-dir"}

	// Thumbnail images options
	opts.WriteThumbnail = BoolOption{param: "--write-thumbnail"}
	opts.WriteAllThumbnails = BoolOption{param: "--write-all-thumbnails"}
	opts.ListThumbnails = BoolOption{param: "--list-thumbnails"}

	// Verbosity / Simulation options
	opts.Quiet = BoolOption{param: "-q"}
	opts.NoWarnings = BoolOption{param: "--no-warnings"}
	opts.Simulate = BoolOption{param: "-s"}
	opts.SkipDownload = BoolOption{param: "--skip-download"}
	opts.GetURL = BoolOption{param: "-g"}
	opts.GetTitle = BoolOption{param: "-e"}
	opts.GetID = BoolOption{param: "--get-id"}
	opts.GetThumbnail = BoolOption{param: "--get-thumbnail"}
	opts.GetDescription = BoolOption{param: "--get-description"}
	opts.GetDuration = BoolOption{param: "--get-duration"}
	opts.GetFilename = BoolOption{param: "--get-filename"}
	opts.GetFormat = BoolOption{param: "--get-format"}
	opts.DumpJSON = BoolOption{param: "-j"}
	opts.DumpSingleJSON = BoolOption{param: "-J"}
	opts.PrintJSON = BoolOption{param: "--print-json"}
	opts.Newline = BoolOption{param: "--newline"}
	opts.NoProgress = BoolOption{param: "--no-progress"}
	opts.ConsoleTitle = BoolOption{param: "--console-title"}
	opts.Verbose = BoolOption{param: "-v"}
	opts.DumpPages = BoolOption{param: "--dump-pages"}
	opts.WritePages = BoolOption{param: "--write-pages"}
	opts.PrintTraffic = BoolOption{param: "--print-traffic"}
	opts.CallHome = BoolOption{param: "-C"}
	opts.NoCallHome = BoolOption{param: "--no-call-home"}

	// Workaround options
	opts.Encoding = StringOption{param: "--encoding"}
	opts.NoCheckCertificate = BoolOption{param: "--no-check-certificate"}
	opts.PreferInsecure = BoolOption{param: "--prefer-insecure"}
	opts.UserAgent = StringOption{param: "--user-agent"}
	opts.Referer = StringOption{param: "--referer"}
	opts.Headers = StringArrayOption{param: "--add-header"}
	opts.BiDiWorkaround = BoolOption{param: "--bidi-workaround"}
	opts.SleepInterval = UintOption{param: "--sleep-interval"}
	opts.MaxSleepInterval = UintOption{param: "--max-sleep-interval"}

	// Video Format options
	opts.Format = StringOption{param: "-f"}
	opts.AllFormats = BoolOption{param: "--all-formats"}
	opts.PreferFreeFormats = BoolOption{param: "--prefer-free-formats"}
	opts.ListFormats = BoolOption{param: "-F"}
	opts.YoutubeSkipDashManifest = BoolOption{param: "--youtube-skip-dash-manifest"}
	opts.MergeOutputFormat = StringOption{param: "--merge-output-format"}

	// Subtitle options
	opts.WriteSub = BoolOption{param: "--write-sub"}
	opts.WriteAutoSub = BoolOption{param: "--write-auto-sub"}
	opts.AllSubs = BoolOption{param: "--all-subs"}
	opts.ListSubs = BoolOption{param: "--list-subs"}
	opts.SubFormat = StringOption{param: "--sub-format"}
	opts.SubLang = StringOption{param: "--sub-lang"}

	// Authentication options
	opts.Username = StringOption{param: "-u"}
	opts.Password = StringOption{param: "-p"}
	opts.TwoFactor = StringOption{param: "-2"}
	opts.NetRc = BoolOption{param: "-n"}
	opts.VideoPassword = StringOption{param: "--video-password"}

	// Adobe Pass options
	opts.ApMso = StringOption{param: "--ap-mso"}
	opts.ApUsername = StringOption{param: "--ap-username"}
	opts.ApPassword = StringOption{param: "--ap-password"}
	opts.ApListMso = BoolOption{param: "--ap-list-mso"}

	// Post-processing options
	opts.ExtractAudio = BoolOption{param: "-x"}
	opts.AudioFormat = StringOption{param: "--audio-format"}
	opts.AudioQuality = StringOption{param: "--audio-quality"}
	opts.RecodeVideo = StringOption{param: "--recode-video"}
	opts.PostprocessorArgs = StringOption{param: "--postprocessor-args"}
	opts.KeepVideo = BoolOption{param: "-k"}
	opts.NoPostOverwrites = BoolOption{param: "--no-post-overwrites"}
	opts.EmbedSubs = BoolOption{param: "--embed-subs"}
	opts.EmbedThumbnail = BoolOption{param: "--embed-thumbnail"}
	opts.AddMetadata = BoolOption{param: "--add-metadata"}
	opts.MetadataFromTitle = StringOption{param: "--metadata-from-title"}
	opts.XAttrs = BoolOption{param: "--xattrs"}
	opts.Fixup = StringOption{param: "--fixup"}
	opts.PreferAvconv = BoolOption{param: "--prefer-avconv"}
	opts.PreferFfmpeg = BoolOption{param: "--prefer-ffmpeg"}
	opts.FfmpegLocation = StringOption{param: "--ffmpeg-location"}
	opts.Exec = StringOption{param: "--exec"}
	opts.ConvertSubs = StringOption{param: "--convert-subs"}

	return opts
}

// Returns the string representation of all the option fields
func (opts *options) OptionsToCliParameters() string {
	var buffer bytes.Buffer

	optReflect := reflect.ValueOf(opts).Elem()

	for i := 0; i < optReflect.NumField(); i++ {
		option := optReflect.Field(i).Interface().(Option)
		buffer.WriteString(option.OptionString() + " ")
	}

	output := strings.Join(strings.Fields(buffer.String()), " ")

	if output == "" || output == " " {
		return ""
	}

	return output
}
