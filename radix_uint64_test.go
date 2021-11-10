// Copyright (C) 2021 OZON SAS, Thierry Fournier <thierry.fournier@ozon.io>
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

import "testing"

func TestRadixUInt64(t *testing.T) {
	var nw1 uint64
	var nw2 uint64
	var r *Radix
	var n *Node
	var s string

	/* Init DB */
	r = NewRadix()

	/* Insert value */
	nw1 = 432343254252
	r.UInt64Insert(nw1, "test - nw1")

	/* Lookup network */
	n = r.UInt64Get(nw1)
	if n == nil {
		t.Errorf("Should match")
	}
	nw2 = n.UInt64GetValue()
	if nw2 != nw1 {
		t.Errorf("Should match")
	}
	s = n.Data.(string)
	if s != "test - nw1" {
		t.Errorf("Should match")
	}
}
