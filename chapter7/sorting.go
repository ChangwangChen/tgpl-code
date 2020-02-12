package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}

	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"

	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	fmt.Fprintln(tw)
	tw.Flush()
}

//定义按照 Artist 排序的类型
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//定义通用的排序类型
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (c customSort) Len() int           { return len(c.t) }
func (c customSort) Less(i, j int) bool { return c.less(c.t[i], c.t[j]) }
func (c customSort) Swap(i, j int)      { c.t[i], c.t[j] = c.t[j], c.t[i] }

func main() {
	var tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m58s")},
		{"Go Ahead", "Changwang", "As I Am", 2008, length("4m38s")},
		{"Ready To Go", "Chen", "Smash", 2019, length("6m38s")},
	}
	printTracks(tracks)

	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	sort.Reverse(byArtist(tracks))
	printTracks(tracks)

	//使用 customSort 类型变量来减少类型的定义
	byYear := customSort{
		t: tracks,
		less: func(x, y *Track) bool {
			return x.Year < y.Year
		},
	}

	sort.Sort(byYear)
	printTracks(byYear.t)

	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		return x.Length < y.Length
	}})
	printTracks(tracks)
}
