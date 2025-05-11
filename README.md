# 🗺️ sitemap

A simple and extensible Go package for generating **XML sitemaps**.

This package helps you create Google-compliant sitemap XML files dynamically in Go. You can group URLs into submaps (for large sites), and write them to disk.

---

## 🚀 Installation

```bash
go get github.com/Kratiwitz/sitemap
```

---

## ✨ Usage

```go
package main

import (
    "github.com/yourusername/sitemap"
)

func main() {
    host := "https://example.com"

    sitemapLocatedPaths := "sitemaps" // in your website http://www.example.com/sitemaps
    // main sitemap will be like this
    // https://www.example.com/sitemap.xml
    // other sitemaps will be created like below
    // https://www.example.com/sitemaps/sitemap-1.xml

    // Initialize sitemap generator
    sm := sitemap.NewSitemap(host, sitemapLocatedPaths)

    // Create a sub-sitemap (useful for splitting large sets)
    submap := sm.AddMap()

    // Add a URL entry
    submap.Add(sitemap.Url{
        Loc:        "tower-of-god",
        Changefreq: "weekly",
        Priority:   "0.8",
    })

    // Write all sitemap files to disk
    err := sm.Render()
    if err != nil {
        panic(err)
    }
}
```

---

## 🧾 Generated XML Example

Main sitemap.xml

```xml
<?xml version="1.0" encoding="UTF-8"?> <sitemapindex xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/siteindex.xsd" xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"><sitemap>
  <loc>https://example.com/sitemaps/sitemap-1.xml</loc>
  <lastmod>2025-05-11T11:38:09+03:00</lastmod>
</sitemap>
</sitemapindex>
```

Sub sitemap.xml

```xml
<?xml version="1.0" encoding="UTF-8"?> <urlset xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd" xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:image="http://www.google.com/schemas/sitemap-image/1.1" xmlns:video="http://www.google.com/schemas/sitemap-video/1.1" xmlns:geo="http://www.google.com/geo/schemas/sitemap/1.0" xmlns:news="http://www.google.com/schemas/sitemap-news/0.9" xmlns:mobile="http://www.google.com/schemas/sitemap-mobile/1.0" xmlns:pagemap="http://www.google.com/schemas/sitemap-pagemap/1.0" xmlns:xhtml="http://www.w3.org/1999/xhtml">
<url>
  <loc>https://example.com/solo-max-level-newbie</loc>
  <lastmod>2025-05-11T11:38:09+03:00</lastmod>
  <changefreq>weekly</changefreq>
  <priority>0.8</priority>
</url>
</urlset>
```

---

## 📄 License

MIT

---

## 🤝 Contributing

Pull requests and issues are welcome. Please open one if you'd like to improve the library or report a bug.

