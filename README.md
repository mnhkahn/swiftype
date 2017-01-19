swiftype v1
========

This is a Golang client for swiftype. Only support search.

### Getting started

To get the package, execute:

	go get gopkg.in/mnhkahn/swiftype.v1

### For more details, see the API documentation.

[godoc](https://godoc.org/github.com/mnhkahn/swiftype)

### Example

	import "gopkg.in/mnhkahn/swiftype.v1"

	var (
		SWIFTYPE        *swiftype.Client
		SWIFTYPE_APIKEY = "YOUR OWN API KEY"
		SWIFTYPE_HOST   = "api.swiftype.com"
		SWIFTYPE_ENGINE = "YOUR OWN ENGINE"
	)

	func InitSwiftype() error {
		SWIFTYPE = swiftype.NewClientWithApiKey(SWIFTYPE_APIKEY, SWIFTYPE_HOST)
		return nil
	}

	func Search(q string) *swiftype.SwiftypeResult {
		data, err := SWIFTYPE.Search(SWIFTYPE_ENGINE, q)
		if err != nil {
			panic(err)
		}
		return data
	}
