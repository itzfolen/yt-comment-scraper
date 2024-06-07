# Simple Golang CLI utility for scraping youtube comment

With this tool you will be able to easily download all public comments from YouTube

## Installation

Install the [Golang](https://go.dev/dl/) compiler for your platform

Clone the project

```bash
$ git clone https://github.com/itzfolen/yt-comment-scraper.git
$ cd yt-comment-scraper
```

Install the necessary Golang libraries

```bash
$ go get google.golang.org/api/option google.golang.org/api/youtube/v3
```

## Preparing for run

[Get YouTube Data API v3](https://console.cloud.google.com/marketplace/product/google/youtube.googleapis.com?q=search&referrer=search&project=fast-envoy-414818&pli=1)

Paste your API key in man.go on line 19 

```golang
apiKey := "INPUT_YOU_YOUTUBE_DATA_API_V3_KEY"
```

## Run

In order to launch our project, we need to specify the video ID of the video on YouTube

Example: We have a link https://www.youtube.com/watch?v=446E-r0rXHI and the video ID of our video will be `446E-r0rXHI`

```bash
$ go run main.go 446E-r0rXHI
```

## Result

As a result, we will get the comments.csv file in our folder

## License

[GNU GPLv3](https://choosealicense.com/licenses/gpl-3.0/)
