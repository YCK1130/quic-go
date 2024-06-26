package wire

import (
	"bytes"

	"github.com/YCK1130/quic-go/internal/protocol"
	"github.com/YCK1130/quic-go/quicvarint"
)

// A DataBlockedFrame is a DATA_BLOCKED frame
type DataBlockedFrame struct {
	MaximumData protocol.ByteCount
}

func parseDataBlockedFrame(r *bytes.Reader, _ protocol.Version) (*DataBlockedFrame, error) {
	offset, err := quicvarint.Read(r)
	if err != nil {
		return nil, err
	}
	return &DataBlockedFrame{MaximumData: protocol.ByteCount(offset)}, nil
}

func (f *DataBlockedFrame) Append(b []byte, version protocol.Version) ([]byte, error) {
	b = append(b, dataBlockedFrameType)
	return quicvarint.Append(b, uint64(f.MaximumData)), nil
}

// Length of a written frame
func (f *DataBlockedFrame) Length(version protocol.Version) protocol.ByteCount {
	return 1 + protocol.ByteCount(quicvarint.Len(uint64(f.MaximumData)))
}
