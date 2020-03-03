package checkgreeting

//GreetingData holds a string called Greeting which is used in SayGreeting function
type GreetingData struct {
	Greeting string
}

//SayGreeting is a receiver function on the GreetingData struct for getting a greeting
func (g *GreetingData) SayGreeting() string {
	if notEmptyWords(g.Greeting) {
		return g.Greeting
	}
	return "Greetings should be more than empty words..."
}

//NotEmptyWords tests whether the greeting has meaning
func notEmptyWords(input string) bool {
	if len(input) > 0 {
		return true
	}
	return false
}
