swiftype v1.0
========

This is a Golang client for swiftype. Only support search.

### doc

[godoc](https://godoc.org/github.com/mnhkahn/swiftype)

### example

	import "github.com/mnhkahn/swiftype"

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