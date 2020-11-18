// Copyright (C) 2020 OZON SAS, Thierry Fournier <thierry.fournier@ozon.io>
//
// This library is free software; you can redistribute it and/or
// modify it under the terms of the GNU Lesser General Public
// License as published by the Free Software Foundation, version 3
// exclusively.
// 
// This library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Lesser General Public License for more details.
// 
// You should have received a copy of the GNU Lesser General Public
// License along with this library; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301  USA

package radix

import "net"
import "testing"

func TestRadixIPv4(t *testing.T) {
	var nw1 *net.IPNet
	var nw2 *net.IPNet
	var nw3 *net.IPNet
	var r *Radix
	var n *Node
	var ns []*Node
	var s string

	/* Init DB */

	r = NewRadix(true)

	/* Insert network */

	nw1 = &net.IPNet{}
	nw1.IP = net.ParseIP("10.4.0.0")
	nw1.Mask = net.CIDRMask(16, 32)
	r.IPv4Insert(nw1, "test - 10.4.0.0/16")

	/* Lookup network */

	n = r.IPv4Get(nw1)
	if n == nil {
		t.Errorf("Should match")
	}
	nw2 = n.IPv4GetNet()
	if !nw2.IP.Equal(nw1.IP) {
		t.Errorf("Should match")
	}
	s = n.Data[0].(string)
	if s != "test - 10.4.0.0/16" {
		t.Errorf("Should match")
	}

	/* Lookup longest */

	nw2 = &net.IPNet{}
	nw2.IP = net.ParseIP("10.4.0.0")
	nw2.Mask = net.CIDRMask(32, 32)

	n = r.IPv4LookupLonguest(nw2)
	if n == nil {
		t.Errorf("Should match")
	}
	nw2 = n.IPv4GetNet()
	if !nw2.IP.Equal(nw1.IP) {
		t.Errorf("Should match")
	}
	s = n.Data[0].(string)
	if s != "test - 10.4.0.0/16" {
		t.Errorf("Should match")
	}

	/* Insert parent network */

	nw3 = &net.IPNet{}
	nw3.IP = net.ParseIP("10.0.0.0")
	nw3.Mask = net.CIDRMask(8, 32)
	r.IPv4Insert(nw3, "test - 10.0.0.0/8")

	/* Lookup longest */

	nw2 = &net.IPNet{}
	nw2.IP = net.ParseIP("10.4.0.0")
	nw2.Mask = net.CIDRMask(32, 32)

	n = r.IPv4LookupLonguest(nw2)
	if n == nil {
		t.Errorf("Should match")
	}
	nw2 = n.IPv4GetNet()
	if !nw2.IP.Equal(nw1.IP) {
		t.Errorf("Should match")
	}
	s = n.Data[0].(string)
	if s != "test - 10.4.0.0/16" {
		t.Errorf("Should match")
	}

	/* Lookup longest path */

	nw2 = &net.IPNet{}
	nw2.IP = net.ParseIP("10.4.0.0")
	nw2.Mask = net.CIDRMask(32, 32)

	ns = r.IPv4LookupLonguestPath(nw2)
	if len(ns) != 2 {
		t.Errorf("Should have length of 2, got %d", len(ns))
	}

	nw2 = ns[0].IPv4GetNet()
	if !nw2.IP.Equal(nw2.IP) {
		t.Errorf("Should match")
	}
	s = ns[0].Data[0].(string)
	if s != "test - 10.0.0.0/8" {
		t.Errorf("Should match")
	}

	nw2 = ns[1].IPv4GetNet()
	if !nw2.IP.Equal(nw2.IP) {
		t.Errorf("Should match")
	}
	s = ns[1].Data[0].(string)
	if s != "test - 10.4.0.0/16" {
		t.Errorf("Should match")
	}

	/* Delete */

	r.IPv4DeleteNetwork(nw1)

	/* Lookup deleted */

	n = r.IPv4Get(nw1)
	if n != nil {
		t.Errorf("Should not match")
	}
}