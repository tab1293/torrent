package torrent

import (
	"net"

	"github.com/anacrolix/dht/v2/krpc"

	"github.com/tab1293/torrent/peer_protocol"
)

// Peer connection info, handed about publicly.
type PeerInfo struct {
	Id     [20]byte
	Addr   net.Addr
	Source PeerSource
	// Peer is known to support encryption.
	SupportsEncryption bool
	peer_protocol.PexPeerFlags
	// Whether we can ignore poor or bad behaviour from the peer.
	Trusted bool
}

func (me PeerInfo) equal(other PeerInfo) bool {
	return me.Id == other.Id &&
		me.Addr.String() == other.Addr.String() &&
		me.Source == other.Source &&
		me.SupportsEncryption == other.SupportsEncryption &&
		me.PexPeerFlags == other.PexPeerFlags &&
		me.Trusted == other.Trusted
}

// Generate PeerInfo from peer exchange
func (me *PeerInfo) FromPex(na krpc.NodeAddr, fs peer_protocol.PexPeerFlags) {
	me.Addr = ipPortAddr{append([]byte(nil), na.IP...), na.Port}
	me.Source = PeerSourcePex
	// If they prefer encryption, they must support it.
	if fs.Get(peer_protocol.PexPrefersEncryption) {
		me.SupportsEncryption = true
	}
	me.PexPeerFlags = fs
}

func (me PeerInfo) addr() IpPort {
	return IpPort{IP: addrIpOrNil(me.Addr), Port: uint16(addrPortOrZero(me.Addr))}
}
