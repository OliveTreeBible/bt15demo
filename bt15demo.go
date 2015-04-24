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

type ChunkInfo struct {
  totalWords int
  chunk string
}

func excuteDemoForURL(documentURL string) {
  doc := fetchTheDocument(documentURL)
  
  chunks := chunkTheText(doc)
  hashes := hashTheChunks(chunks)
  success := indexTheHashes(hashes)
  if success {
	fmt.Println("Done!")
  }
}

func fetchTheDocument(documentURL string) (*goquery.Document) {
  doc, err := goquery.NewDocument(documentURL)
  
  if err != nil {
    log.Fatal(err)
  }
  return doc
}

func prepTheDatabase() {
  c, _ := sqlite3.Open("sqlite.db")
  c.Exec("DROP TABLE IF EXISTS block_index")
  c.Exec("CREATE TABLE block_index (block_id INTEGER PRIMARY KEY, block_hash VARCHAR(255) NOT NULL, total_words INTEGER NOT NULL)")
  c.Exec("CREATE INDEX block_hash_idx ON block_index (block_hash)")
  // go use(c)
}

func chunkTheText(doc *goquery.Document) ([]ChunkInfo) {
  var chunks []ChunkInfo
  breakAccountedFor := false
  chunk := ""
  totalWords := 0
  
  doc.Find("p").Each(func(i int, s *goquery.Selection) {
    text := s.Text()
	
    segmenter := segment.NewWordSegmenterDirect([]byte(text))
    for segmenter.Segment() {
        tokenBytes := segmenter.Bytes()
        tokenType := segmenter.Type()
		
		
		if tokenType != None   {
				chunk += segmenter.Text()
				breakAccountedFor = false
				totalWords++
		} else if tokenType == None {
				if !breakAccountedFor {
						chunk += " "
						breakAccountedFor = true
				}
		} else {
				fmt.Printf("Unknown TokenType(%d) encountered for token: %s", tokenType, tokenBytes)
		}
    }
	
    if err := segmenter.Err(); err != nil {
        log.Fatal(err)
    }
	
	if totalWords >= 50 {
	  chunks = append(chunks, ChunkInfo{totalWords, chunk})
	  // Reset for the next round
	  totalWords = 0
	  chunk = ""
	}
  })
  return chunks
}

func hashTheChunks(chunks []ChunkInfo) ([]string) {
  var hashes []string = make([]string, 0, len(chunks))
  
  for i, chunkInfo := range chunks {
	hasher := sha256.New()
    hasher.Write([]byte(chunkInfo.chunk))
    hash := hasher.Sum(nil)
    hashes = append(hashes, hex.EncodeToString(hash))
	fmt.Printf("\n%s\n====> %d words %s\n", chunkInfo.chunk, chunkInfo.totalWords, hashes[i])
  }
  
  return hashes
}

func indexTheHashes(hashes []string) (bool) {
  // TODO:
  return false
}

func main() {
    excuteDemoForURL("https://www.gutenberg.org/files/39452/39452-h/39452-h.htm")
}
