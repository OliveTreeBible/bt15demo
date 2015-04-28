package main

import (
    "github.com/PuerkitoBio/goquery"
    "github.com/blevesearch/segment"
	"code.google.com/p/go-sqlite/go1/sqlite3"
    "fmt"
    "crypto/sha256"
	"encoding/hex"
    "log"
	"flag"
	"strings"
)

const (
	None = iota
	Number
	Letter
	Kana
	Ideo
)

type ChunkInfo struct {
  positionIndex int
  totalWords int
  chunk string
  hash string
}

type DocumentInfo struct {
  documentURL string
  databaseFileName string
  databaseHandle *sqlite3.Conn
  htmlDocument *goquery.Document
  chunks []ChunkInfo
}

func chunkTheText(docInfo *DocumentInfo) bool {
  failed := false
  breakAccountedFor := false
  chunk := ""
  totalWords := 0
  positionIndex := -1
  
  docInfo.htmlDocument.Find("p,pre,div,span,h1,h2,h3").Each(func(i int, s *goquery.Selection) {
	if !failed {
	// A bit of a bug here as this does not handle tags embedded in the text at
	// all so for instance if you have:
	//		'... some text<span>blah blah</span>more text...'
	// you will get '... some textblah blahmore text...' which most likely isn't
	// what you want. This is an area that would need to be discussed in detail
	// if this approach was to be standardized for sharng annotations or for
	// citations.
    text := s.Text() 
	if	len(chunk) > 0 && !strings.HasSuffix(chunk, " ") {
	  chunk += " "
	}
	positionIndex = i//docInfo.htmlDocument.IndexOfSelection(s)
	
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
        failed = true
    } else if totalWords >= 50 {
	  docInfo.chunks = append(docInfo.chunks, ChunkInfo{positionIndex, totalWords, chunk, ""})
	  // Reset for the next round
	  breakAccountedFor = false
	  totalWords = 0
	  chunk = ""
	}
	}
  })
  
  // In case there is a straggler 
  if !failed && len(chunk) > 0 && totalWords > 0 {
	docInfo.chunks = append(docInfo.chunks, ChunkInfo{positionIndex, totalWords, chunk, ""})
  }
  
  return !failed
}

func hashTheChunks(docInfo *DocumentInfo) (bool){
  for index := range docInfo.chunks {
	hasher := sha256.New()
    hasher.Write([]byte(docInfo.chunks[index].chunk))
    hash := hasher.Sum(nil)
    docInfo.chunks[index].hash = hex.EncodeToString(hash)
	fmt.Printf("\n%s\n====> %d words %s\n", docInfo.chunks[index].chunk, docInfo.chunks[index].totalWords, docInfo.chunks[index].hash)
  }
  return true
}

func indexTheHashes(docInfo *DocumentInfo) (bool) {
  fmt.Println("Chunks ", len(docInfo.chunks))
  for _, chunk := range docInfo.chunks {
	args := sqlite3.NamedArgs{"$block_hash": chunk.hash, "$total_words": chunk.totalWords, "$position_index":chunk.positionIndex}
	docInfo.databaseHandle.Exec("INSERT INTO block_index (block_hash, total_words, position_index) VALUES($block_hash, $total_words, $position_index)", args)
  }
  return true
}

func prepDocInfo(docInfo *DocumentInfo) (error) {
  var err error
  flag.StringVar(&docInfo.documentURL, "document-url", "https://www.gutenberg.org/files/39452/39452-h/39452-h.htm", "The input document URL that you wish to process.")
  flag.StringVar(&docInfo.databaseFileName, "database-filename", "btdemo15.sqlite", "The output filename for the datbase index.")
  flag.Parse()
  
  docInfo.htmlDocument, err = goquery.NewDocument(docInfo.documentURL)
  if err == nil {
	docInfo.databaseHandle, err = sqlite3.Open(docInfo.databaseFileName)
	if err == nil {
	  docInfo.databaseHandle.Exec("DROP TABLE IF EXISTS block_index")
	  docInfo.databaseHandle.Exec("CREATE TABLE block_index (block_id INTEGER PRIMARY KEY ASC AUTOINCREMENT, block_hash VARCHAR(255) NOT NULL, total_words INTEGER NOT NULL, position_index INTEGER DEFAULT -1)")
	  docInfo.databaseHandle.Exec("CREATE INDEX block_hash_idx ON block_index (block_hash)")
	}
  }
  
  return err
}

func main() {
  docInfo := DocumentInfo{}
  
  if err := prepDocInfo(&docInfo); err == nil {
	if chunkTheText(&docInfo) &&
	   hashTheChunks(&docInfo) &&
	   indexTheHashes(&docInfo) {
		fmt.Println("Success!")
	}
  } else {
	log.Fatal(err)
  }
  docInfo.databaseHandle.Close()
}
