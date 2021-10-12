package torrent

import (
	"context"

	"github.com/kylec725/graytorrent/internal/common"
	"github.com/kylec725/graytorrent/rpc"
	pb "github.com/kylec725/graytorrent/rpc"
	"github.com/pkg/errors"
)

// server.go contains implementations of the required grpc server functions

// ErrTorrentNotFound indicates that a torrent was not found
var ErrTorrentNotFound = errors.New("Torrent not found")

// List current managed torrents
func (s *Session) List(in *pb.Empty, stream pb.TorrentService_ListServer) error {
	for _, to := range s.torrentList {
		stream.Send(&pb.Torrent{
			Id:          to.ID,
			Name:        to.Info.Name,
			InfoHash:    to.InfoHash[:], // NOTE: may need to check if infohash is set first
			TotalLength: uint32(to.Info.TotalLength),
			Left:        uint32(to.Info.Left),
			DownRate:    uint32(to.DownRate()),
			UpRate:      uint32(to.UpRate()),
			State:       rpc.Torrent_State(to.State()),
		})
	}
	return nil
}

// Add a new torrent to be managed
func (s *Session) Add(ctx context.Context, in *pb.AddRequest) (*pb.TorrentReply, error) {
	to, err := s.AddTorrent(ctx, in.File) // TODO: need to check if torrent is already managed
	if err != nil {
		return nil, err
	}
	return &pb.TorrentReply{Id: to.ID, Name: to.Info.Name, InfoHash: to.InfoHash[:]}, nil
}

// Remove a torrent from being managed
func (s *Session) Remove(ctx context.Context, in *pb.TorrentRequest) (*pb.TorrentReply, error) {
	var infoHash [20]byte
	copy(infoHash[:], in.GetInfoHash())

	if to, ok := s.torrentList[infoHash]; ok {
		s.RemoveTorrent(to)
		return &pb.TorrentReply{Id: to.ID, Name: to.Info.Name, InfoHash: to.InfoHash[:]}, nil
	}
	// Check ID instead
	for _, to := range s.torrentList {
		if to.ID == in.GetId() {
			s.RemoveTorrent(to)
			return &pb.TorrentReply{Id: to.ID, Name: to.Info.Name, InfoHash: to.InfoHash[:]}, nil
		}
	}

	return nil, ErrTorrentNotFound
}

// Start a torrent's download/upload
func (s *Session) Start(ctx context.Context, in *pb.TorrentRequest) (*pb.TorrentReply, error) {
	var infoHash [20]byte
	copy(infoHash[:], in.GetInfoHash())

	if to, ok := s.torrentList[infoHash]; ok {
		newCtx := context.WithValue(context.Background(), common.KeyPort, s.port) // NOTE: using ctx causes to.Start() to end immediately
		go to.Start(newCtx)
		return &pb.TorrentReply{Id: to.ID, Name: to.Info.Name, InfoHash: to.InfoHash[:]}, nil
	}
	// Check ID instead
	for _, to := range s.torrentList {
		if to.ID == in.GetId() {
			newCtx := context.WithValue(context.Background(), common.KeyPort, s.port) // NOTE: using ctx causes to.Start() to end immediately
			go to.Start(newCtx)
			return &pb.TorrentReply{Id: to.ID, Name: to.Info.Name, InfoHash: to.InfoHash[:]}, nil
		}
	}

	return nil, ErrTorrentNotFound
}

// Stop a torrent's download/upload
func (s *Session) Stop(ctx context.Context, in *pb.TorrentRequest) (*pb.TorrentReply, error) {
	var infoHash [20]byte
	copy(infoHash[:], in.GetInfoHash())

	if to, ok := s.torrentList[infoHash]; ok {
		to.Stop()
		return &pb.TorrentReply{Id: to.ID, Name: to.Info.Name, InfoHash: to.InfoHash[:]}, nil
	}
	// Check ID instead
	for _, to := range s.torrentList {
		if to.ID == in.GetId() {
			to.Stop()
			return &pb.TorrentReply{Id: to.ID, Name: to.Info.Name, InfoHash: to.InfoHash[:]}, nil
		}
	}

	return nil, ErrTorrentNotFound
}
