package command

import (
	"github.com/AnAverageBeing/mclib/data/packetid"
	pk "github.com/AnAverageBeing/mclib/net/packet"
)

type Client interface {
	SendPacket(p pk.Packet)
}

// ClientJoin implement server.Component for Graph
func (g *Graph) ClientJoin(client Client) {
	client.SendPacket(pk.Marshal(
		packetid.ClientboundCommands, g,
	))
}
