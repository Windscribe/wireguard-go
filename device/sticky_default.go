// +build !linux android

package device

import (
	"github.com/windscribe/wireguard-go/conn"
	"github.com/windscribe/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
