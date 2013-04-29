package swiftype

import (
    "os"
    "testing"
)

func Test_NewClientWithApiKey(t *testing.T) {
    c := NewClientWithApiKey("key", "host")

    if (c != nil) {
        t.Log("NewClient creation passed")
    } else {
        t.Error("Failed to create a new client")
    }

    if (c.api_key != "key") {
        t.Error("Failed to set api_key correctly")   
    }
}

func Test_Engines(t *testing.T) {
    api_key := os.Getenv("SWIFTYPE_API_KEY")
    c := NewClientWithApiKey(api_key, DEFAULT_API_HOST)

    engines := c.Engines()

    if (engines == nil) {
        t.Error("Failed to get engines correctly")
    }

    t.Log("Engines: ", engines)
}

func Test_Engine(t *testing.T) {
    api_key := os.Getenv("SWIFTYPE_API_KEY")
    c := NewClientWithApiKey(api_key, DEFAULT_API_HOST)

    engine := c.Engine("testing")

    if (engine != nil) {
        t.Log("Engine: ", engine)
    } else {
        t.Error("Failed to get engine correctly")
    }

}

func Test_Search(t *testing.T) {
    api_key := os.Getenv("SWIFTYPE_API_KEY")
    c := NewClientWithApiKey(api_key, DEFAULT_API_HOST)

    results := c.Search("testing", "Great")

    if (results == nil) {
        t.Error("Failed to get search results correctly")
    } else {
        t.Log("Results: ", results)    
    }

}
