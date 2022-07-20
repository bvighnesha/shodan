# shodan

[![Build Status](https://travis-ci.org/B-V-R/shodan.svg?branch=master)](https://travis-ci.org/B-V-R/shodan)
[![GitHub issues](https://img.shields.io/github/issues/B-V-R/shodan.svg?style=flat-square)]()
[![license](https://img.shields.io/github/license/B-V-R/shodan.svg?style=flat-square)]()

Example:

```Go
package main

import (
	"context"
	"fmt"
	"vighnesha.in/shodan"
)

func main() {
	s := shodan.Configure("API_KEY")
	str, err := s.RestAPI().Search().Ports(context.Background())
	fmt.Println(str, err)
}

```

refer https://vighnesh.org/shodan for Documentation and Examples.


