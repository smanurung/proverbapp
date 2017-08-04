// avoid to using package main here (library-driven development principle)
package bibel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	Username     = "ed6ueZHsxRgY6qIGdiQn7Pn2aZwNDbdSY6fYL5bs"
	Password     = "x" // this doesn't matter.
	BiblesOrgURL = "https://bibles.org/v2"

	ProverbsVerseTotal = 31
)

type Bibel struct{}

type Verse struct {
	Ref string `json:"reference"`
	Txt string `json:"text"`
}

type ChapterResponse struct {
	Response struct {
		Verses []Verse `json:"verses"`
	} `json:"response"`
}

func (b *Bibel) HandleProverbs() {
	now := time.Now().Unix()
	random := now%ProverbsVerseTotal + 1

	endpoint := fmt.Sprintf("%s/chapters/eng-MSG:Prov.%d/verses.js", BiblesOrgURL, random)
	verseRequest, e := http.NewRequest("GET", endpoint, nil)
	if e != nil {
		log.Println(e)
		return
	}
	verseRequest.SetBasicAuth(Username, Password)

	client := http.Client{}
	resp, e := client.Do(verseRequest)
	if e != nil {
		log.Println(e)
		return
	}
	defer resp.Body.Close()

	encoded, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Println(e)
		return
	}

	var chapter ChapterResponse
	e = json.Unmarshal(encoded, &chapter)
	if e != nil {
		log.Println(e)
		return
	}

	random = now%int64(len(chapter.Response.Verses)) + 1

	fmt.Printf("%s\n", chapter.Response.Verses[random].Ref)
	fmt.Printf("------\n")
	fmt.Printf("%s\n", chapter.Response.Verses[random].Txt)
	return
}
