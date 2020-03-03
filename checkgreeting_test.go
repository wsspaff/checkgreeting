package checkgreeting

import (
	"testing"
)

func TestNotEmptyWords(t *testing.T) {

	//Test that good greeting returns correctly
	myData := &GreetingData{Greeting: "somewhere over the rainbow"}

	wantGreetingStr1 := "somewhere over the rainbow"
	gotGreetingStr1 := myData.SayGreeting()
	if gotGreetingStr1 != wantGreetingStr1 {
		t.Errorf("Wrong response returned - wanted: %s, got: %s", wantGreetingStr1, gotGreetingStr1)
	}

	//Test that missing greeting returns correctly
	myData = &GreetingData{Greeting: ""}

	wantGreetingStr1 = "Greetings should be more than empty words..."
	gotGreetingStr1 = myData.SayGreeting()
	if gotGreetingStr1 != wantGreetingStr1 {
		t.Errorf("Wrong response returned - wanted: %s, got: %s", wantGreetingStr1, gotGreetingStr1)
	}

}
