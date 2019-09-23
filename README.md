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
