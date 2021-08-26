package common

// State indicates the state and thus behavior of a torrent goroutine
type State uint8

const (
	// Started torrent state
	Started State = iota
	// Stopped torrent state
	Stopped
)
