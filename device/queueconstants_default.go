// +build !android,!ios

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2020 WireGuard LLC. All Rights Reserved.
 */

package device

const (
	QueueOutboundSize          = 10240
	QueueInboundSize           = 10240
	QueueHandshakeSize         = 1024
	MaxSegmentSize             = (1 << 16) - 1 // largest possible UDP datagram
	PreallocatedBuffersPerPool = 0             // Disable and allow for infinite memory growth
)
