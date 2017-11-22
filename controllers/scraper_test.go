package controllers

import (
    _ "fmt"
    "testing"
    "time"
    _ "github.com/LucaTony/SIMS/models"
)

//TestCheckDue is a test function to test the CheckDue function
//with all the possible cases
func TestCheckDue (t *testing.T) {

    testTime1 := time.Now().Add(-23*time.Hour)
    testTime2 := time.Now().Add(-25*time.Hour)

    testText1 := "## as df Test 123 +*ç %&()=?"
    testText2 := "## as df Test 123 +*ç"
    testText3 := "aaa bb sdfj asdfsadf asdf adsfas df"

    //Case: Text is same and new, no update needed
    if CheckDue(testText1, testText1, testTime1) != false { t.Error("Shouldn't update, because same") }

    //Case: Text is same but old, no update needed
    if CheckDue(testText1, testText1, testTime2) != false { t.Error("Shouldn't update, because same") }

    //Case: Text is different and new, no update needed
    if CheckDue(testText1, testText2, testTime1) != false { t.Error("Shouldn't update, change is too new") }

    //Case: Text is different a lot different and old, update not allowed
    if CheckDue(testText1, testText3, testTime2) != false { t.Error("Shouldn't update, change is too much") }

    //Case: Text is different and old, update needed
    if CheckDue(testText1, testText2, testTime2) != true { t.Error("Should update") }
}

