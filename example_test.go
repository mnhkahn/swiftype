package swiftype

import "github.com/mnhkahn/swiftype"

func Example() {
	SWIFTYPE := *swiftype.Client
	SWIFTYPE_APIKEY := "YOUR OWN API KEY"
	SWIFTYPE_HOST := "api.swiftype.com"
	SWIFTYPE_ENGINE := "YOUR OWN ENGINE"

	SWIFTYPE := swiftype.NewClientWithApiKey(SWIFTYPE_APIKEY, SWIFTYPE_HOST)

	data, err := SWIFTYPE.Search(SWIFTYPE_ENGINE, q)
	if err != nil {
		panic(err)
	}
	_ = data
}
