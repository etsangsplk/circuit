// Copyright 2013 The Go Circuit Project
// Use of this source code is governed by the license for
// The Go Circuit Project, found in the LICENSE file.
//
// Authors:
//   2013 Petar Maymounkov <p@gocircuit.org>

package client

import (
	"io"

	"github.com/gocircuit/circuit/kit/element/valve"
)

type Chan interface {
	Send() (io.WriteCloser, error)
	IsDone() bool
	Scrub()
	Close() error
	Recv() (io.ReadCloser, error)
	Cap() int
	Stat() Stat
}

type Stat struct {
	Cap int
	Opened bool
	Closed bool
	Aborted bool
	NumSend int
	NumRecv int
}

func statstat(s valve.Stat) Stat {
	return Stat{
		Cap: s.Cap,
		Opened: s.Opened,
		Closed: s.Closed,
		Aborted: s.Aborted,
		NumSend: s.NumSend,
		NumRecv: s.NumRecv,
	}
}

type yvalveChan struct {
	valve.YValve
}

func (y yvalveChan) Stat() Stat {
	return statstat(y.YValve.Stat())
}
