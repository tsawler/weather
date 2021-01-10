# Weather

Gets current weather status from Open Weather Maps.

Installation:

```
go get -u github.com/tsawler/weather
```

Sample usage:

~~~go
package main

import (
	"fmt"
	"github.com/tsawler/weather_course/weather"
	"log"
	"net/http"
	"time"
)

func main() {
	myWeather := weather.API{
		Client:  &http.Client{Timeout: 10 * time.Second},
		Key:     "3a17a33b2c242e1bc22f0303ff5c7ebc",
		City:    "Fredericton",
		Country: "ca",
	}

	x, err := myWeather.CurrentWeather()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Current weather is", x.Simple[0].Description)
	fmt.Println(fmt.Sprintf("Temp: %0.2f, but feels like %0.2f", x.Temperature.Temp, x.Temperature.FeelsLike))
}
~~~