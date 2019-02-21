package api

import (
	"context"
	"log"
	"googlemaps.github.io/maps"
)

func Maps(ori string, dest string) []maps.Route{
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyDvq_9Eg7Reuu83LQUIzu-e-tKZE7GkIj4"))
	if err != nil {
		log.Fatalf("fatal error1: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin: 		ori,
		Destination:	dest,
	}
	route, _, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error2: %s", err)
	}

	return route
}