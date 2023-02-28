# api-server #
Simple api server for CRUD operation <br>
## Technology Used ##
- Golang
- JWT Authentication
- Cobra CLI
## Running the server ##
<pre> go run . start </pre>
<pre> go run . start -p port-number </pre>
## Data model ##
<pre><code>
// For string a cricket players basic information
type Player struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Country string  `json:"country"`
	Batting Batting `json:"batting"`
	Bowling Bowling `json:"bowling"`
}

// For a players batting information
type Batting struct {
	Innings int `json:"innings"`
	Runs    int `json:"runs"`
}

// For a players bowling information
type Bowling struct {
	Innings int `json:"innings"`
	Wickets int `json:"wickets"`
}
</code></pre>

