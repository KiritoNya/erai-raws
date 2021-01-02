package eraiRaws

import (
	"github.com/mmcdole/gofeed"
)

type Link string

const (
	RssMagnetLink_1080         Link = "https://www.erai-raws.info/rss-1080-magnet"
	RssMagnetLink_720          Link = "https://www.erai-raws.info/rss-720-magnet"
	RssMagnetLink_480          Link = "https://www.erai-raws.info/rss-480-magnet"
	RssMagnetMovieLink_1080    Link = "https://www.erai-raws.info/rss-movies-1080-magnet"
	RssMagnetMovieLink_720     Link = "https://www.erai-raws.info/rss-movies-720-magnet"
	RssMagnetMovieLink_480     Link = "https://www.erai-raws.info/rss-movies-480-magnet"
	RssMagnetBatchLink_1080    Link = "https://www.erai-raws.info/rss-batch-1080-magnet"
	RssMagnetBatchLink_720     Link = "https://www.erai-raws.info/rss-batch-720-magnet"
	RssMagnetBatchLink_480     Link = "https://www.erai-raws.info/rss-batch-480-magnet"
	RssMagnetRawsLink_1080     Link = "https://www.erai-raws.info/rss-raws-1080-magnet"
	RssMagnetRawsLink_720      Link = "https://www.erai-raws.info/rss-raws-720-magnet"
	RssMagnetEncodedLink_1080  Link = "https://www.erai-raws.info/rss-encoded-1080-magnet"
	RssMagnetEncodedLink_720   Link = "https://www.erai-raws.info/rss-encoded-720-magnet"
	RssTorrentLink_1080        Link = "https://www.erai-raws.info/rss-1080-torrent"
	RssTorrentLink_720         Link = "https://www.erai-raws.info/rss-720-torrent"
	RssTorrentLink_480         Link = "https://www.erai-raws.info/rss-480-torrent"
	RssTorrentMovieLink_1080   Link = "https://www.erai-raws.info/rss-movies-1080-torrent"
	RssTorrentMovieLink_720    Link = "https://www.erai-raws.info/rss-movies-720-torrent"
	RssTorrentMovieLink_480    Link = "https://www.erai-raws.info/rss-movies-480-torrent"
	RssTorrentBatchLink_1080   Link = "https://www.erai-raws.info/rss-batch-1080-torrent"
	RssTorrentBatchLink_720    Link = "https://www.erai-raws.info/rss-batch-720-torrent"
	RssTorrentBatchLink_480    Link = "https://www.erai-raws.info/rss-batch-480-torrent"
	RssTorrentRawsLink_1080    Link = "https://www.erai-raws.info/rss-raws-1080-torrent"
	RssTorrentRawsLink_720     Link = "https://www.erai-raws.info/rss-raws-720-torrent"
	RssTorrentEncodedLink_1080 Link = "https://www.erai-raws.info/rss-encoded-1080-torrent"
	RssTorrentEncodedLink_720  Link = "https://www.erai-raws.info/rss-encoded-720-torrent"
	RssDailyEpisode            Link = "https://www.erai-raws.info/rss-eps"
	RssNotifications           Link = "https://www.erai-raws.info/rss-notifications"
	RssMovies                  Link = "https://www.erai-raws.info/rss-movies"
	RssBatch                   Link = "https://www.erai-raws.info/rss-batch"
	RssRaws                    Link = "https://www.erai-raws.info/rss-raws"
	RssEncoded                 Link = "https://www.erai-raws.info/rss-encoded"
)

/*type RssTorrentEpisode struct {
	Title   string    `xml:"title"`
	Link    string    `xml:"link"`
	PubDate time.Time `xml:"pubDate"`
}

type RssPostEpisode struct {
	Title string `xml:"title"`
	Link string `xml:"link"`
	PubDate string `xml:"pubDate"`
	Guid string `xml:"guid"`
	Enclosure struct {
		Url string `xml:"url"`
		Lenght int64 `xml:"lenght"`
		Type string `xml:"type"`
	} `xml:"enclosure"`
	Description string `xml:"description"`
}*/
func NewRssTorrent(link Link) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(string(link))
	if err != nil {
		return nil, err
	}
	return feed, nil
}
