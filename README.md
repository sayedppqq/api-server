# api-server #
Simple api server for CRUD operation <br>
## Technology Used ##
- Golang
- JWT Authentication
- Cobra CLI
## Running the server ##
### Running the server from direct source code ##
```git clone https://github.com/sayedppqq/api-server.git``` <br>

Go to the api-server directory and run <br>
```go mod tidy``` <br>
```go mod vendor``` <br>
```go run . start``` or ```go run . start -p port-number```
### Running the server from docker image ###
```docker pull sayedppqq/api-server```<br>
```docker run -dp <choosen port>:8080 sayedppqq/api-server```
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

## api calls ##

| Method |                     url                      | payload | actions                                     |
|--------|:--------------------------------------------:|--------:|---------------------------------------------|
| GET    |         ```http://localhost:8080/```         |      No | call home page and get new token            |
| GET    |   ```http://localhost:8080/api/players```    |      No | call all the players exist in data model    |
| GET    | ```http://localhost:8080/api/players/{id}``` |      No | call by a specific player id                |
| POST   | ```http://localhost:8080/api/players/add```  |     Yes | add new player to the data slice            |
| PUT    | ```http://localhost:8080/api/update/{id}```  |     Yes | update a specific player by calling with id |
| DELETE | ```http://localhost:8080/api/delete/{id}```  |      No | delete a specific by calling its id         |


## CAUTIONS ##
- All the api call is authorized by JWT token. So we need to pass token as a valid "token" header while calling methods(except for homepage request) by postman or curl.
- A token can be generated by call homepage initially.
