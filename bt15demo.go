package main

import (
    "github.com/PuerkitoBio/goquery"
    "github.com/blevesearch/segment"
	"code.google.com/p/go-sqlite/go1/sqlite3"
    "fmt"
    "crypto/sha256"
	"encoding/hex"
    "log"   
)

const (
	None = iota
	Number
	Letter
	Kana
	Ideo
)

func ExampleScrape(url string) {
  doc, err := goquery.NewDocument(url)
  
  c, _ := sqlite3.Open("sqlite.db")
  c.Exec("DROP TABLE IF EXISTS block_index")
  c.Exec("CREATE TABLE block_index (block_id INTEGER PRIMARY KEY, block_hash VARCHAR(255) NOT NULL, total_words INTEGER NOT NULL)")
  c.Exec("CREATE INDEX block_hash_idx ON block_index (block_hash)")
  // go use(c)
  
  if err != nil {
    log.Fatal(err)
  }

  breakAccountedFor := false
  normalizedString := ""
  blockWordCount := 0
  
  doc.Find("p").Each(func(i int, s *goquery.Selection) {
    text := s.Text()
	blockWordCount = 0
	normalizedString = ""
	
    segmenter := segment.NewWordSegmenterDirect([]byte(text))
    for segmenter.Segment() {
        tokenBytes := segmenter.Bytes()
        tokenType := segmenter.Type()
		
		
		if tokenType != None   {
				normalizedString += segmenter.Text()
				breakAccountedFor = false
				blockWordCount++
		} else if tokenType == None {
				if !breakAccountedFor {
						normalizedString += " "
						breakAccountedFor = true
				}
		} else {
				fmt.Printf("Unknown TokenType(%d) encountered for token: %s", tokenType, tokenBytes)
		}
    }
	
    if err := segmenter.Err(); err != nil {
        log.Fatal(err)
    }
	
    hash := sha256.New()
    hash.Write([]byte(normalizedString))
    md := hash.Sum(nil)
    mdStr := hex.EncodeToString(md)
    fmt.Printf("\n%s\n====> %d words %s\n", normalizedString, blockWordCount, mdStr)
  })
}

func main() {
    ExampleScrape("https://www.gutenberg.org/files/39452/39452-h/39452-h.htm")
}
