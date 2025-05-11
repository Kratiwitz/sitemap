package sitemap

import (
	"testing"
)

func TestSitemap_Render(t *testing.T) {
	tests := []struct {
		name    string
		sitemap func() *Sitemap
		wantErr bool
	}{
		{
			name: "test only main xml",
			sitemap: func() *Sitemap {
				sm := NewSitemap("https://example.com", "sitemaps")

				return sm
			},
			wantErr: false,
		},
		{
			name: "test with sub xmls",
			sitemap: func() *Sitemap {
				sm := NewSitemap("https://example.com", "sitemaps")

				_ = sm.AddMap()
				_ = sm.AddMap()

				return sm
			},
			wantErr: false,
		},
		{
			name: "test with sub xmls",
			sitemap: func() *Sitemap {
				sm := NewSitemap("https://example.com", "sitemaps")

				subm := sm.AddMap()
				subm.Add(Url{
					Loc:        "jujutsu-kaizen",
					Changefreq: "weekly",
					Priority:   "0.8",
				})

				subm2 := sm.AddMap()
				subm2.Add(Url{
					Loc:        "jujutsu-kaizen",
					Changefreq: "weekly",
					Priority:   "0.8",
				})

				return sm
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := tt.sitemap().Render()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Render() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Render() succeeded unexpectedly")
			}
		})
	}
}
