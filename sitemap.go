package sitemap

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var mainXmlHead string = `<?xml version="1.0" encoding="UTF-8"?> <sitemapindex xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/siteindex.xsd" xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`

var mainXmlBottom string = `</sitemapindex>
`

var mainXmlRecordTemplate string = `<sitemap>
  <loc>%s</loc>
  <lastmod>%s</lastmod>
</sitemap>
`

var xmlHead string = `<?xml version="1.0" encoding="UTF-8"?> <urlset xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd" xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:image="http://www.google.com/schemas/sitemap-image/1.1" xmlns:video="http://www.google.com/schemas/sitemap-video/1.1" xmlns:geo="http://www.google.com/geo/schemas/sitemap/1.0" xmlns:news="http://www.google.com/schemas/sitemap-news/0.9" xmlns:mobile="http://www.google.com/schemas/sitemap-mobile/1.0" xmlns:pagemap="http://www.google.com/schemas/sitemap-pagemap/1.0" xmlns:xhtml="http://www.w3.org/1999/xhtml">`

var urlSectionTemplate string = `
<url>
  <loc>%s</loc>
  <lastmod>%s</lastmod>
  <changefreq>%s</changefreq>
  <priority>%s</priority>
</url>
`

var (
	timeLayout             string = "2006-01-02T15:04:05-07:00"
	mainSitemapFileName    string = "sitemap.xml"
	subSitemapFileTemplate string = "sitemap-%d.xml"
	urlTemplate            string = "%s/%s/%s"
	locTemplate            string = "%s/%s"
	saveFolder             string = "maps"
	urlSetClose            string = "</urlset>"
)

type Sitemap struct {
	Submaps      []*Submap
	mu           sync.Mutex
	host         string
	sitemapsPath string
}

func NewSitemap(host string, sitemapsPath string) *Sitemap {
	return &Sitemap{
		Submaps:      []*Submap{},
		host:         host,
		sitemapsPath: sitemapsPath,
	}
}

type Submap struct {
	Urls    []Url
	Sitemap *Sitemap
	index   int
	mu      *sync.Mutex
}

type Url struct {
	Loc        string
	Changefreq string
	Priority   string
}

func (s *Sitemap) AddMap() *Submap {
	s.mu.Lock()
	defer s.mu.Unlock()

	submap := &Submap{
		Sitemap: s,
		Urls:    []Url{},
		index:   len(s.Submaps) + 1,
		mu:      &s.mu,
	}

	s.Submaps = append(s.Submaps, submap)

	return submap
}

func (s *Submap) Add(url Url) *Submap {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Urls = append(s.Urls, url)

	return s
}

func (s *Submap) filename() string {
	return fmt.Sprintf(subSitemapFileTemplate, s.index)
}

func (s *Sitemap) Render() error {
	if len(s.Submaps) == 0 {
		return nil
	}

	var sb strings.Builder

	sb.WriteString(mainXmlHead)

	for _, submap := range s.Submaps {
		if len(submap.Urls) == 0 {
			continue
		}

		url := fmt.Sprintf(urlTemplate, s.host, s.sitemapsPath, submap.filename())
		submapRecord := fmt.Sprintf(mainXmlRecordTemplate, url, getFormattedCurrentTime())
		sb.WriteString(submapRecord)

		err := submap.render()
		if err != nil {
			return err
		}
	}

	sb.WriteString(mainXmlBottom)

	filePath, err := getFullFileName(mainSitemapFileName)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, []byte(sb.String()), 0644)
	if err != nil {
		return err
	}

	return nil
}

func getFormattedCurrentTime() string {
	return time.Now().Format(timeLayout)
}

func getFullFileName(filename string) (string, error) {
	fullPath := filepath.Join(saveFolder, filename)

	err := os.MkdirAll(saveFolder, os.ModePerm)
	if err != nil {
		return "", err
	}

	return fullPath, nil
}

func (s *Submap) render() error {
	if len(s.Urls) == 0 {
		return nil
	}

	var sb strings.Builder

	sb.WriteString(xmlHead)

	for _, url := range s.Urls {
		loc := fmt.Sprintf(locTemplate, s.Sitemap.host, url.Loc)

		urlSection := fmt.Sprintf(
			urlSectionTemplate,
			loc,
			getFormattedCurrentTime(),
			url.Changefreq,
			url.Priority,
		)

		sb.WriteString(urlSection)
	}

	sb.WriteString(urlSetClose)

	filePath, err := getFullFileName(s.filename())
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, []byte(sb.String()), 0644)
	if err != nil {
		return err
	}

	return nil
}
