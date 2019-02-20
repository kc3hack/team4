package api

import (
	"context"
	"log"

	"googlemaps.github.io/maps"
)

func Maps() []maps.Route{
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyDvq_9Eg7Reuu83LQUIzu-e-tKZE7GkIj4"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:      "Sydney",
		Destination: "Perth",
	}
	route, _, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	return route
}