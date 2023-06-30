# Bikram Sambat Progress

<div align="center">
  <img src="./tweet.png" alt="Bikram Sambat Tweet Sample" />
  <p>A tweet bot that tweets yearly progress of the Nepali Calendar a.k.a <a href="https://en.wikipedia.org/wiki/Vikram_Samvat">Bikram Sambat.</a></p>
  <p><em>Follow <a href="https://twitter.com/bikram_sambat">Bikram Sambat</a> on Twitter.</em></p>
</div>

## About

### Prerequisite

- [GNU Make](https://ftp.gnu.org/old-gnu/Manuals/make-3.79.1/html_chapter/make_1.html)
- [Docker](https://docs.docker.com/install/)

### Usage

Export your [Twitter secrets](https://developer.twitter.com).

```bash
# Twitter Secrets
export TWITTER_CONSUMER_KEY="WEbDazCqR6fVERc8SuD4tK5c"
export TWITTER_CONSUMER_SECRET="BPhmzeZqJ5nfAGGDeEQM8Xrh"
export TWITTER_ACCESS_TOKEN="9gdteNLPQVw5C9vcBfHR8Kkj"
export TWITTER_ACCESS_TOKEN_SECRET="5apcG4rbwSKJvXPtPb2yWk54"
```

And,

```bash
$ make build
$ make run

2020/04/06 15:09:26 Bikram Sambat Progress: 98.320859
2020/04/06 15:09:27 Bikram Sambat Progress: 98.320862
2020/04/06 15:09:28 Bikram Sambat Progress: 98.320865
2020/04/06 15:09:29 Bikram Sambat Progress: 98.320868
2020/04/06 15:09:30 Bikram Sambat Progress: 98.320871
2020/04/06 15:09:31 Bikram Sambat Progress: 98.320875
```

## License

[MIT](LICENSE)
