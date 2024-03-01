package geecache

import "github.com/xwxb/MyGeeCache/pb"

// PeerPicker is the interface that must be implemented to locate
// the peer that owns a specific key.
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented by a peer.
// 可以做 http 和 RPC 两种实现，参数不同的情况尝试在下层去进行解决
type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}
