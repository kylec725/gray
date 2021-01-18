package write

import (
    "testing"
    "fmt"
    "os"

    "github.com/kylec725/graytorrent/torrent"
    "github.com/stretchr/testify/assert"
)

const debugWrite = false

func TestNewWriteSingle(t *testing.T) {
    assert := assert.New(t)

    to := torrent.Torrent{Source: "../tmp/change.torrent"}
    err := to.Setup()
    assert.Nil(err, "torrent Setup() error")

    // Remove the torrent's filename if it exists
    if _, err := os.Stat(to.Name); err == nil {
        err = os.Remove(to.Name)
        if err != nil {
            panic("Removing test file failed")
        }
    }

    err = NewWrite(to)
    if assert.Nil(err) {
        if debugWrite {
            fmt.Println("File created:", to.Name)
        }

        // Test that creating an identical file throws an error
        err = NewWrite(to)
        assert.NotNil(err)
    }
}

func TestNewWriteMulti(t *testing.T) {
    assert := assert.New(t)

    to := torrent.Torrent{Source: "../tmp/batonroad.torrent"}
    err := to.Setup()
    assert.Nil(err, "torrent Setup() error")

    // Remove the torrent's filename if it exists
    if _, err := os.Stat(to.Name); err == nil {
        err = os.RemoveAll(to.Name)
        if err != nil {
            panic("Removing test file failed")
        }
    }

    err = NewWrite(to)
    if assert.Nil(err) {
        if debugWrite {
            fmt.Println("File created:", to.Name)
        }

        // Test that creating an identical file throws an error
        err = NewWrite(to)
        assert.NotNil(err)
    }
}

func TestPieceSize(t *testing.T) {
    to := torrent.Torrent{Source: "../tmp/change.torrent"}
    err := to.Setup()
    assert.Nil(t, err, "torrent Setup() error")

    assert.Equal(t, 193972, pieceSize(&to, 149))
}

func TestAddBlock(t *testing.T) {
    assert := assert.New(t)

    to := torrent.Torrent{Source: "../tmp/change.torrent"}
    err := to.Setup()
    assert.Nil(err, "torrent Setup() error")

    index := 8
    begin := 0
    piece := make([]byte, to.PieceLength)
    block := []byte("hello")

    if debugWrite {
        fmt.Println("BlockLength:", len(block))
        fmt.Println("PieceLength:", len(piece))
    }

    err = AddBlock(&to, index, begin, block, piece)
    assert.Nil(err)
    assert.Equal(block, piece[begin:begin + len(block)])
}
