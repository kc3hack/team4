package api

import (
	"context"
	"log"
	"googlemaps.github.io/maps"
)

func MapsWithWayPoint(ori string, dest string, mode maps.Mode, way []string) []maps.Route{
	c, err := maps.NewClient(maps.WithAPIKey(""))
	if err != nil {
		log.Fatalf("fatal error1: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin: 		ori,
		Destination:	dest,
		Mode:			mode,
		Waypoints:		way,
	}
	route, _, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error2: %s", err)
	}

	return route
}

func Maps(ori string, dest string, mode maps.Mode) []maps.Route{
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyDvq_9Eg7Reuu83LQUIzu-e-tKZE7GkIj4"))
	if err != nil {
		log.Fatalf("fatal error1: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin: 		ori,
		Destination:	dest,
		Mode:			mode,
	}
	route, _, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error2: %s", err)
	}

	return route
}
