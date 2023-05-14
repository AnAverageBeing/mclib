package world

import "github.com/AnAverageBeing/mclib/level"

type EventsListener struct {
	LoadChunk   func(pos level.ChunkPos) error
	UnloadChunk func(pos level.ChunkPos) error
}
