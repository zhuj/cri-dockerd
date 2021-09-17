// +build !dockerless !windows

package remote

import (
	"net"

	"github.com/coreos/go-systemd/v22/activation"
	"github.com/pkg/errors"
)

func listenFd(addr string) (net.Listener, error) {
	var (
		err       error
		listeners []net.Listener
	)

	listeners, err = activation.Listeners()
	if err != nil {
		return nil, err
	}

	if len(listeners) == 0 {
		return nil, errors.New("no sockets found via socket activation: make sure the service was started by systemd")
	}

	if addr == "" {
		return listeners[0], nil
	}

	//TODO: systemd fd selection (default is 3)
	return nil, errors.New("not supported yet")
}
