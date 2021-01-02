package eraiRaws

import (
	"github.com/mmcdole/gofeed"
	"time"
)

type Link string
type LinkMagnet string
type LinkTorrent string

const (
	RssDailyEpisode            Link        = "https://www.erai-raws.info/rss-eps"
	RssNotifications           Link        = "https://www.erai-raws.info/rss-notifications"
	RssMovies                  Link        = "https://www.erai-raws.info/rss-movies"
	RssBatch                   Link        = "https://www.erai-raws.info/rss-batch"
	RssRaws                    Link        = "https://www.erai-raws.info/rss-raws"
	RssEncoded                 Link        = "https://www.erai-raws.info/rss-encoded"
	RssMagnetLink_1080         LinkMagnet  = "https://www.erai-raws.info/rss-1080-magnet"
	RssMagnetLink_720          LinkMagnet  = "https://www.erai-raws.info/rss-720-magnet"
	RssMagnetLink_480          LinkMagnet  = "https://www.erai-raws.info/rss-480-magnet"
	RssMagnetMovieLink_1080    LinkMagnet  = "https://www.erai-raws.info/rss-movies-1080-magnet"
	RssMagnetMovieLink_720     LinkMagnet  = "https://www.erai-raws.info/rss-movies-720-magnet"
	RssMagnetMovieLink_480     LinkMagnet  = "https://www.erai-raws.info/rss-movies-480-magnet"
	RssMagnetBatchLink_1080    LinkMagnet  = "https://www.erai-raws.info/rss-batch-1080-magnet"
	RssMagnetBatchLink_720     LinkMagnet  = "https://www.erai-raws.info/rss-batch-720-magnet"
	RssMagnetBatchLink_480     LinkMagnet  = "https://www.erai-raws.info/rss-batch-480-magnet"
	RssMagnetRawsLink_1080     LinkMagnet  = "https://www.erai-raws.info/rss-raws-1080-magnet"
	RssMagnetRawsLink_720      LinkMagnet  = "https://www.erai-raws.info/rss-raws-720-magnet"
	RssMagnetEncodedLink_1080  LinkMagnet  = "https://www.erai-raws.info/rss-encoded-1080-magnet"
	RssMagnetEncodedLink_720   LinkMagnet  = "https://www.erai-raws.info/rss-encoded-720-magnet"
	RssTorrentLink_1080        LinkTorrent = "https://www.erai-raws.info/rss-1080-torrent"
	RssTorrentLink_720         LinkTorrent = "https://www.erai-raws.info/rss-720-torrent"
	RssTorrentLink_480         LinkTorrent = "https://www.erai-raws.info/rss-480-torrent"
	RssTorrentMovieLink_1080   LinkTorrent = "https://www.erai-raws.info/rss-movies-1080-torrent"
	RssTorrentMovieLink_720    LinkTorrent = "https://www.erai-raws.info/rss-movies-720-torrent"
	RssTorrentMovieLink_480    LinkTorrent = "https://www.erai-raws.info/rss-movies-480-torrent"
	RssTorrentBatchLink_1080   LinkTorrent = "https://www.erai-raws.info/rss-batch-1080-torrent"
	RssTorrentBatchLink_720    LinkTorrent = "https://www.erai-raws.info/rss-batch-720-torrent"
	RssTorrentBatchLink_480    LinkTorrent = "https://www.erai-raws.info/rss-batch-480-torrent"
	RssTorrentRawsLink_1080    LinkTorrent = "https://www.erai-raws.info/rss-raws-1080-torrent"
	RssTorrentRawsLink_720     LinkTorrent = "https://www.erai-raws.info/rss-raws-720-torrent"
	RssTorrentEncodedLink_1080 LinkTorrent = "https://www.erai-raws.info/rss-encoded-1080-torrent"
	RssTorrentEncodedLink_720  LinkTorrent = "https://www.erai-raws.info/rss-encoded-720-torrent"
)

type RSS interface {
	GetTitleRss() string
	GetLinkRss() string
}

//RssMagnet is a type that contains the RSS with the magnet url.
type RssMagnet struct {
	feed *gofeed.Feed
}

//NewRssMagnet create an object of type RssMagnet.
func NewRssMagnet(link LinkMagnet) (RssMagnet, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(string(link))
	if err != nil {
		return RssMagnet{}, err
	}
	return RssMagnet{feed: feed}, nil
}

//GetNumberItems is a method of obtaining the number of items of the RSS file.
func (rm *RssMagnet) GetNumberItems() int {
	return len(rm.feed.Items)
}

//GetTitleRss is a method of obtaining the title of the RSS file.
func (rm *RssMagnet) GetTitleRss() string {
	return rm.feed.Title
}

//GetLinkRss is a method of obtaining the link of site.
func (rm *RssMagnet) GetLinkRss() string {
	return rm.feed.Link
}

//GetDescriptionRss is a method of obtaining the description of rss.
func (rm *RssMagnet) GetDescriptionRss() string {
	return rm.feed.Description
}

//GetItemsTitle is a method that returns a slice of string with the titles of the various items inside.
func (rm *RssMagnet) GetItemsTitle() (titles []string) {
	for _, item := range rm.feed.Items {
		titles = append(titles, item.Title)
	}
	return titles
}

//GetItemsLink is a method that returns a slice of string with the links of the various items inside.
func (rm *RssMagnet) GetItemsLinks() (links []string) {
	for _, item := range rm.feed.Items {
		links = append(links, item.Link)
	}
	return links
}

//GetItemsPubDate is a method that returns a slice of string with the public dates of the various items inside.
func (rm *RssMagnet) GetItemsPubDate() (dates []time.Time) {
	for _, item := range rm.feed.Items {
		dates = append(dates, *item.PublishedParsed)
	}
	return dates
}

//GetItemTitle is a method that returns an string with the title of the item defined by the number of items sent as a parameter..
func (rm *RssMagnet) GetItemTitle(numItem int) string {
	return rm.feed.Items[numItem].Title
}

//GetItemLink is a method that returns an string with the link of the item defined by the number of items sent as a parameter..
func (rm *RssMagnet) GetItemLink(numItem int) string {
	return rm.feed.Items[numItem].Link
}

//GetItemPubDate is a method that returns an string with the link of the item defined by the number of items sent as a parameter..
func (rm *RssMagnet) GetItemPubDate(numItem int) time.Time {
	return *rm.feed.Items[numItem].PublishedParsed
}
