# Yandex Geocoding API

[![GoDoc](https://godoc.org/github.com/FlameInTheDark/go-yandex-geocoding?status.svg)](https://godoc.org/github.com/FlameInTheDark/go-yandex-geocoding)
[![Go Report Card](https://goreportcard.com/badge/github.com/FlameInTheDark/go-yandex-geocoding)](https://goreportcard.com/report/github.com/FlameInTheDark/go-yandex-geocoding)
[![Build Status](https://travis-ci.org/FlameInTheDark/go-yandex-geocoding.svg?branch=master)](https://travis-ci.org/FlameInTheDark/go-yandex-geocoding)
[![codecov](https://codecov.io/gh/FlameInTheDark/go-yandex-geocoding/branch/master/graph/badge.svg)](https://codecov.io/gh/FlameInTheDark/go-yandex-geocoding)

### Usage

```go
    package main
    
    import (
    	"fmt"
    	"github.com/FlameInTheDark/go-yandex-geocoding"
    )

    func main() {
    	ygi := yageocoding.New("your_api_key")
    	
    	meters, err := ygi.RangeBtw("Chelyabinsk, Zakharenko, 12", "Chelyabinsk, Chicherina, 25")
    	if err != nil {
    		fmt.Println(err.Error())
    		return
    	}
    	
    	fmt.Println("Distance btw two addresses ", meters, " meters")
    	
    	result, err := ygi.Find("Russia, Chelyabinsk, Kuznetcova, 2")
    	if err != nil {
    		fmt.Println(err.Error())
    		return
    	}
    	fmt.Println(result.Address())
    	fmt.Println(result.Latitude())
    	fmt.Println(result.Longitude())
    	fmt.Println(result.CountryCode())
    	members := *result.Members()
    	// For one element of list of found elements
    	fmt.Println(members[0].PostalCode())
    }
```