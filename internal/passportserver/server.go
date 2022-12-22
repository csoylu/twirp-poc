package passportserver

import (
	"context"
	"math/rand"

	"github.com/csoylu/twirp-poc/rpc/passport"
	"github.com/twitchtv/twirp"
)

func New() *randomHaberdasher {
	return new(randomHaberdasher)
}

type randomHaberdasher struct{}

func (h *randomHaberdasher) MakeHat(ctx context.Context, size *passport.Size) (*passport.Hat, error) {
	// When returning an error, it's best to use the error constructors defined in
	// the Twirp package so that the client gets a well-structured error response.
	if size.Inches <= 0 {
		return nil, twirp.InvalidArgumentError("Inches", "I can't make a hat that small!")
	}
	colors := []string{"white", "black", "brown", "red", "blue"}
	names := []string{"bowler", "baseball cap", "top hat", "derby"}
	return &passport.Hat{
		Size:  size.Inches,
		Color: colors[rand.Intn(len(colors))],
		Name:  names[rand.Intn(len(names))],
	}, nil
}
