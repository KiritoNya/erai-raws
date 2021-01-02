package eraiRaws

import (
	"testing"
)

func TestNewRssMagnet(t *testing.T) {
	_, err := NewRssMagnet(RssMagnetLink_1080)
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
}

func TestRssMagnet_GetTitleRss(t *testing.T) {
	result := "Erai-raws RSS"

	rm, err := NewRssMagnet(RssMagnetLink_1080)
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	title := rm.GetTitleRss()

	if title != result {
		t.Log("not obtain", result, "but obtain", title)
	} else {
		t.Log("[OK]")
	}
}

func TestRssMagnet_GetDescriptionRss(t *testing.T) {
	result := "Seasonal releases 1080p (Magnet)"

	rm, err := NewRssMagnet(RssMagnetLink_1080)
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	title := rm.GetDescriptionRss()

	if title != result {
		t.Log("not obtain", result, "but obtain", title)
	} else {
		t.Log("[OK]")
	}
}

func TestRssMagnet_GetLinkRss(t *testing.T) {
	result := "https://www.erai-raws.info"

	rm, err := NewRssMagnet(RssMagnetLink_1080)
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	title := rm.GetLinkRss()

	if title != result {
		t.Log("not obtain", result, "but obtain", title)
	} else {
		t.Log("[OK]")
	}
}

func TestRssMagnet_GetItemsTitle(t *testing.T) {
	result := []string{
		"[1080p] D4DJ: First Mix – 09 (VRV)",     //Changes constantly
		"[1080p] Dogeza de Tanondemita – 12 END", //Changes constantly
	}

	rm, err := NewRssMagnet(RssMagnetLink_1080)
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	titles := rm.GetItemsTitle()

	titles = titles[:2]

	for i, title := range titles {
		if title != result[i] {
			t.Fatal("not obtain", result[i], "but obtain", title)
		} else {
			t.Log("item", i, "[OK]")
		}
	}
}

func TestRssMagnet_GetItemsLinks(t *testing.T) {
	result := []string{
		"magnet:?xt=urn:btih:2E6E0BD82CC270A819C00180850B7A3ABC3C24D7&dn=%5bErai-raws%5d%20D4DJ%20-%20First%20Mix%20-%2009%20%5bVRV%5d%5b1080p%5d%5bMultiple%20Subtitle%5d.mkv&tr=http%3a%2f%2fnyaa.tracker.wf%3a7777%2fannounce&tr=http%3a%2f%2fanidex.moe%3a6969%2fannounce&tr=http%3a%2f%2ftracker.anirena.com%3a80%2fannounce&tr=udp%3a%2f%2fopen.stealth.si%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2fexodus.desync.com%3a6969%2fannounce&tr=wss%3a%2f%2ftracker.openwebtorrent.com&tr=http%3a%2f%2fopen.acgnxtracker.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=udp%3a%2f%2f9.rarbg.to%3a2910%2fannounce&tr=udp%3a%2f%2f9.rarbg.me%3a2720%2fannounce&tr=udp%3a%2f%2ftracker.cyberia.is%3a6969%2fannounce&tr=udp%3a%2f%2ftracker3.itzmx.com%3a6961%2fannounce&tr=udp%3a%2f%2ftracker.torrent.eu.org%3a451%2fannounce&tr=udp%3a%2f%2ftracker.tiny-vps.com%3a6969%2fannounce&tr=udp%3a%2f%2fretracker.lanta-net.ru%3a2710%2fannounce", //Changes constantly
		"magnet:?xt=urn:btih:800FA436A80044B9EDB574773B5BE059B6159187&dn=%5bErai-raws%5d%20Dogeza%20de%20Tanondemita%20-%2012%20END%20%5b1080p%5d%5bMultiple%20Subtitle%5d.mkv&tr=http%3a%2f%2fnyaa.tracker.wf%3a7777%2fannounce&tr=http%3a%2f%2fanidex.moe%3a6969%2fannounce&tr=http%3a%2f%2ftracker.anirena.com%3a80%2fannounce&tr=udp%3a%2f%2fopen.stealth.si%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2fexodus.desync.com%3a6969%2fannounce&tr=wss%3a%2f%2ftracker.openwebtorrent.com&tr=http%3a%2f%2fopen.acgnxtracker.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=udp%3a%2f%2f9.rarbg.to%3a2920%2fannounce&tr=udp%3a%2f%2f9.rarbg.me%3a2810%2fannounce&tr=udp%3a%2f%2ftracker.cyberia.is%3a6969%2fannounce&tr=udp%3a%2f%2ftracker3.itzmx.com%3a6961%2fannounce&tr=udp%3a%2f%2ftracker.torrent.eu.org%3a451%2fannounce&tr=udp%3a%2f%2ftracker.tiny-vps.com%3a6969%2fannounce&tr=udp%3a%2f%2fretracker.lanta-net.ru%3a2710%2fannounce", //Changes constantly
	}

	rm, err := NewRssMagnet(RssMagnetLink_1080)
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	links := rm.GetItemsLinks()

	links = links[:2]

	for i, link := range links {
		if link != result[i] {
			t.Fatal("not obtain", result[i], "but obtain", link)
		} else {
			t.Log("item", i, "[OK]")
		}
	}
}

func TestRssMagnet_GetItemsPubDate(t *testing.T) {
	result := []string{
		"2021-01-01 14:17:32 +0000 UTC", //Changes constantly
		"2020-12-30 15:41:41 +0000 UTC", //Changes constantly
	}

	rm, err := NewRssMagnet(RssMagnetLink_1080)
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	dates := rm.GetItemsPubDate()

	dates = dates[:2]

	for i, date := range dates {
		if date.String() != result[i] {
			t.Fatal("not obtain", result[i], "but obtain", date.String())
		} else {
			t.Log("item", i, "[OK]")
		}
	}
}

func TestRssMagnet_GetItemTitle(t *testing.T) {
	result := "[1080p] D4DJ: First Mix – 09 (VRV)" //Changes continuously

	rm, err := NewRssMagnet(RssMagnetLink_1080)
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}

	title := rm.GetItemTitle(0) //last items

	if title != result {
		t.Error("not obtain", result, "but obtain", title)
	} else {
		t.Log("[OK]")
	}
}

func TestRssMagnet_GetItemLink(t *testing.T) {
	result := "magnet:?xt=urn:btih:2E6E0BD82CC270A819C00180850B7A3ABC3C24D7&dn=%5bErai-raws%5d%20D4DJ%20-%20First%20Mix%20-%2009%20%5bVRV%5d%5b1080p%5d%5bMultiple%20Subtitle%5d.mkv&tr=http%3a%2f%2fnyaa.tracker.wf%3a7777%2fannounce&tr=http%3a%2f%2fanidex.moe%3a6969%2fannounce&tr=http%3a%2f%2ftracker.anirena.com%3a80%2fannounce&tr=udp%3a%2f%2fopen.stealth.si%3a80%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce&tr=udp%3a%2f%2fexodus.desync.com%3a6969%2fannounce&tr=wss%3a%2f%2ftracker.openwebtorrent.com&tr=http%3a%2f%2fopen.acgnxtracker.com%3a80%2fannounce&tr=udp%3a%2f%2ftracker.coppersurfer.tk%3a6969%2fannounce&tr=udp%3a%2f%2f9.rarbg.to%3a2910%2fannounce&tr=udp%3a%2f%2f9.rarbg.me%3a2720%2fannounce&tr=udp%3a%2f%2ftracker.cyberia.is%3a6969%2fannounce&tr=udp%3a%2f%2ftracker3.itzmx.com%3a6961%2fannounce&tr=udp%3a%2f%2ftracker.torrent.eu.org%3a451%2fannounce&tr=udp%3a%2f%2ftracker.tiny-vps.com%3a6969%2fannounce&tr=udp%3a%2f%2fretracker.lanta-net.ru%3a2710%2fannounce" //Changes continuously

	rm, err := NewRssMagnet(RssMagnetLink_1080)
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}

	link := rm.GetItemLink(0) //last items

	if link != result {
		t.Error("not obtain", result, "but obtain", link)
	} else {
		t.Log("[OK]")
	}
}

func TestRssMagnet_GetItemPubDate(t *testing.T) {
	result := "2021-01-01 14:17:32 +0000 UTC" //Changes continuously

	rm, err := NewRssMagnet(RssMagnetLink_1080)
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}

	date := rm.GetItemPubDate(0) //last items

	if date.String() != result {
		t.Error("not obtain", result, "but obtain", date.String())
	} else {
		t.Log("[OK]")
	}
}
