# goydl

A simple youtube-dl library for go.

[![CircleCI](https://circleci.com/gh/BrianAllred/NYoutubeDL.svg?style=svg)](https://circleci.com/gh/BrianAllred/NYoutubeDL) [![Go Report Card](https://goreportcard.com/badge/github.com/brianallred/goydl)](https://goreportcard.com/report/github.com/brianallred/goydl) [![Sourcegraph](https://sourcegraph.com/github.com/BrianAllred/goydl/-/badge.svg)](https://sourcegraph.com/github.com/BrianAllred/goydl?badge)

See the [main page](https://rg3.github.io/youtube-dl/) for youtube-dl for more information.

### Get the package
`go get github.com/BrianAllred/goydl`

### Use the code
See the [documentation](https://github.com/rg3/youtube-dl/blob/master/README.md#readme) for youtube-dl first to understand what it does and how it does it.

1. Create a new youtubeDl client:

        youtubeDl := goydl.NewYoutubeDl()

2. The Options object contains the various youtube-dl download parameters:

        youtubeDl.Options.Output.Value = "/path/to/downloads/video.mp3"
        youtubeDl.Options.ExtractAudio.Value = true
        youtubeDl.Options.AudioFormat.Value = "mp3"

        // Or update the binary
        youtubeDl.Options.Update.Value = true

        // Optional, required if binary is not in $PATH
        youtubeDl.YoutubeDlPath = "/path/to/youtube-dl"

3. Listen to console output (optional, but recommended):

        go io.Copy(os.Stdout, youtubeDl.Stdout)
        go io.Copy(os.Stderr, youtubeDl.Stderr)

4. Start the download:

        cmd, err := youtubeDl.Download("http://videosite.com/videoURL")

        if err != nil {
            log.Fatal(err)
        }

5. Check the video's info:

        fmt.Printf("Title: %s\n", youtubeDl.Info.Title)

6. Wait for the download to finish:

        // Synchronously:
        cmd.Wait()
        
        // Asynchronously:
        defer cmd.Wait()