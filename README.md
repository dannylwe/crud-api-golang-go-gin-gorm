[![Codacy Badge](https://app.codacy.com/project/badge/Grade/f022e00be11b4e59bd4a3dc36c46725c)](https://www.codacy.com/manual/dannylwe/crud-api-golang-go-gin-gorm?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=dannylwe/crud-api-golang-go-gin-gorm&amp;utm_campaign=Badge_Grade)

#### Installation and Run
- run `git clone https://github.com/dannylwe/crud-api-golang-go-gin-gorm`  
- run `go mod download`  
- run `go run main.go`  

#### API
path: /api/v1/todos  

/people
- `GET`: Get all people  
- `POST`: Post a person  

/people/:id
- `GET`: Get a single person  
- `PUT`: Update a single person  
- `DELETE`: Delete a single person  

#### Post Payload
`{
	"FirstName": "Bruce",
	"LastName": "Banner"
}`

#### Structure
```
├── README.md
├── go.mod
├── go.sum
└── todo
    ├── common
    │   └── database.go
    ├── gorm.db
    ├── handlers
    │   └── person.go
    ├── main.go
    ├── migrations
    │   └── person.go
    ├── models
    │   └── person.go
    └── routers
        └── people.go
```
