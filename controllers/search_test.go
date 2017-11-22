package controllers

import(
    "testing"
    "strings"
    "sort"
    _ "github.com/LucaTony/SIMS/models"
)

//TestHome is a test function to test the Home
func TestHome(t *testing.T) {
    searches := []DataSend{}
    tempSearch := DataSend{ //TODO: Direct
        Title: "TestTitle",
        Body:  "Vg37 eBT8 6Fw3 yVD5 3YrN 5AgH rg9A 7Mtt 5inS YvJ4",
        Url:   "TestUrl",
        Score: 0,
    }
    searches = append(searches, tempSearch)

    var mySearch = "Vg37 eBT8 6Fw3 yVD5 3YrN 5AgH rg9A 7Mtt 5inS YvJ4"
    var mySend = []DataSend {}
    var testScore = 0

    words := strings.Fields(mySearch)
    for i := range searches {
        showentry := false
        score := 0
        for w := range words {
            if (strings.Count(searches[i].Body, words[w]) != 0) {
                score += strings.Count(searches[i].Body, words[w]) // add the amount one searched word occurs in a result
                showentry = true
            }
        }
        if (showentry) {
            tempSend := DataSend{ //TODO: Direct
                Title: searches[i].Title,
                Body:  searches[i].Body,
                Url:   searches[i].Url,
                Score: score,
            }
            mySend = append(mySend, tempSend)
            //fmt.Println("added")
        }
        //if (strings.Contains(strings.ToLower(w), strings.ToLower(searches[i].Body))) { // a searched word 
    }

    sort.Sort(ByScore(mySend))
    testScore = mySend[0].Score

    if (testScore != 10) {
        t.Error("Basic search failed: Score should be 10")
    }
}
