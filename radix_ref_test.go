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


func Test_ref(t *testing.T) {
	var r *Radix
	var ref uint32
	var ref_back uint32
	var nref *Node
	var nref_back *Node

	r = NewRadix()
	r.growth()
	r.growth()
	r.growth()
	r.growth()
	r.growth()

	if r.capacity != (5 * 65536) - 1 {
		t.Fatalf("Expect capacity of %d, got %d", (5 * 65536) - 1, r.capacity)
	}

	ref = uint32(3 << 16 | 4343)
	ref_back = r.n2r(r.r2n(ref))
	if ref != ref_back {
		t.Fatalf("Expect reference %x, got %x", ref, ref_back)
	}

	nref = &r.pool[3].nodes[16332]
	nref_back = r.r2n(r.n2r(nref))
	if nref != nref_back {
		t.Fatalf("Expect reference %p, got %p", nref, nref_back)
	}
}
