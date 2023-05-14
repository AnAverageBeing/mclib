package basic

import (
	"github.com/AnAverageBeing/mclib/data/packetid"
	pk "github.com/AnAverageBeing/mclib/net/packet"
)

func (p *Player) handleKeepAlivePacket(packet pk.Packet) error {
	var KeepAliveID pk.Long
	if err := packet.Scan(&KeepAliveID); err != nil {
		return Error{err}
	}
	// Response
	err := p.c.Conn.WritePacket(pk.Packet{
		ID:   int32(packetid.ServerboundKeepAlive),
		Data: packet.Data,
	})
	if err != nil {
		return Error{err}
	}
	return nil
}
