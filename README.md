# Simple Golang CLI utility for scraping youtube comment

With this tool you will be able to easily download all public comments from YouTube

## Installation

Clone the project

```bash
$ git clone https://github.com/itzfolen/yt-comment-scraper.git
$ cd yt-comment-scraper
```

Install the necessary Golang libraries

```bash
$ go get google.golang.org/api/option google.golang.org/api/youtube/v3
```

## Usage

[Get YouTube Data API v3](https://console.cloud.google.com/marketplace/product/google/youtube.googleapis.com?q=search&referrer=search&project=fast-envoy-414818&pli=1)

Paste your API key in man.go on line 19 

```golang
apiKey := "INPUT_YOU_YOUTUBE_DATA_API_V3_KEY"
```

## License

[GNU GPLv3](https://choosealicense.com/licenses/gpl-3.0/)
