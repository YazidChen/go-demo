module yazid.com/hello

go 1.16

require (
	github.com/google/go-cmp v0.5.5
	rsc.io/quote v1.5.2
	yazid.com/greetings v0.0.0-00010101000000-000000000000
)

replace yazid.com/greetings => ../greetings

replace yazid.com/hello => ../hello

replace yazid.com/hello/morestrings => ../hello/morestrings
