module github.com/tokopedia/gripmock

go 1.22

require (
	github.com/go-chi/chi/v5 v5.0.12
	github.com/golang/protobuf v1.5.4
	github.com/lithammer/fuzzysearch v1.1.8
	github.com/stretchr/testify v1.7.0
	github.com/tokopedia/gripmock/protogen/example v0.0.0
	google.golang.org/grpc v1.64.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/tokopedia/gripmock/protogen v0.0.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto v0.0.0-20240604185151-ef581f913117 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240604185151-ef581f913117 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

// this is for generated server to be able to run
replace github.com/tokopedia/gripmock/protogen/example v0.0.0 => ./protogen/example

// this is for example client to be able to run
replace github.com/tokopedia/gripmock/protogen v0.0.0 => ./protogen
