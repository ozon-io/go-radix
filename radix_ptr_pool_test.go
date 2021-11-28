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

import "reflect"
import "testing"

func Test_ptr_range(t *testing.T) {
	var r *Radix
	var want []ptr_range

	r = &Radix{}

	r.add_range(12, 13, 1, 0)
	want = []ptr_range{
		ptr_range{12, 13, 1, 0},
	}
	if !reflect.DeepEqual(r.ptr_range, want) {
		t.Errorf("Unmatched:\n   got: %#v\nexpect: %#v\n", r.ptr_range, want)
	}

	r.add_range(14, 15, 2, 0)
	want = []ptr_range{
		ptr_range{12, 13, 1, 0},
		ptr_range{14, 15, 2, 0},
	}
	if !reflect.DeepEqual(r.ptr_range, want) {
		t.Errorf("Unmatched:\n   got: %#v\nexpect: %#v\n", r.ptr_range, want)
	}

	r.add_range(1, 2, 3, 0)
	want = []ptr_range{
		ptr_range{1, 2, 3, 0},
		ptr_range{12, 13, 1, 0},
		ptr_range{14, 15, 2, 0},
	}
	if !reflect.DeepEqual(r.ptr_range, want) {
		t.Errorf("Unmatched:\n   got: %#v\nexpect: %#v\n", r.ptr_range, want)
	}

	r.add_range(6, 9, 4, 0)
	want = []ptr_range{
		ptr_range{1, 2, 3, 0},
		ptr_range{6, 9, 4, 0},
		ptr_range{12, 13, 1, 0},
		ptr_range{14, 15, 2, 0},
	}
	if !reflect.DeepEqual(r.ptr_range, want) {
		t.Errorf("Unmatched:\n   got: %#v\nexpect: %#v\n", r.ptr_range, want)
	}

	r.add_range(7, 8, 5, 0)
	want = []ptr_range{
		ptr_range{1, 2, 3, 0},
		ptr_range{6, 9, 4, 0},
		ptr_range{12, 13, 1, 0},
		ptr_range{14, 15, 2, 0},
	}
	if !reflect.DeepEqual(r.ptr_range, want) {
		t.Errorf("Unmatched:\n   got: %#v\nexpect: %#v\n", r.ptr_range, want)
	}

	r.add_range(6, 7, 5, 0)
	want = []ptr_range{
		ptr_range{1, 2, 3, 0},
		ptr_range{6, 9, 4, 0},
		ptr_range{12, 13, 1, 0},
		ptr_range{14, 15, 2, 0},
	}
	if !reflect.DeepEqual(r.ptr_range, want) {
		t.Errorf("Unmatched:\n   got: %#v\nexpect: %#v\n", r.ptr_range, want)
	}

	r.add_range(7, 8, 5, 0)
	want = []ptr_range{
		ptr_range{1, 2, 3, 0},
		ptr_range{6, 9, 4, 0},
		ptr_range{12, 13, 1, 0},
		ptr_range{14, 15, 2, 0},
	}
	if !reflect.DeepEqual(r.ptr_range, want) {
		t.Errorf("Unmatched:\n   got: %#v\nexpect: %#v\n", r.ptr_range, want)
	}

	r.add_range(6, 9, 5, 0)
	want = []ptr_range{
		ptr_range{1, 2, 3, 0},
		ptr_range{6, 9, 4, 0},
		ptr_range{12, 13, 1, 0},
		ptr_range{14, 15, 2, 0},
	}
	if !reflect.DeepEqual(r.ptr_range, want) {
		t.Errorf("Unmatched:\n   got: %#v\nexpect: %#v\n", r.ptr_range, want)
	}

	r.add_range(5, 8, 5, 0)
	want = []ptr_range{
		ptr_range{1, 2, 3, 0},
		ptr_range{6, 9, 4, 0},
		ptr_range{12, 13, 1, 0},
		ptr_range{14, 15, 2, 0},
	}
	if !reflect.DeepEqual(r.ptr_range, want) {
		t.Errorf("Unmatched:\n   got: %#v\nexpect: %#v\n", r.ptr_range, want)
	}

	r.add_range(8, 10, 5, 0)
	want = []ptr_range{
		ptr_range{1, 2, 3, 0},
		ptr_range{6, 9, 4, 0},
		ptr_range{12, 13, 1, 0},
		ptr_range{14, 15, 2, 0},
	}
	if !reflect.DeepEqual(r.ptr_range, want) {
		t.Errorf("Unmatched:\n   got: %#v\nexpect: %#v\n", r.ptr_range, want)
	}

	r.add_range(5, 10, 5, 0)
	want = []ptr_range{
		ptr_range{1, 2, 3, 0},
		ptr_range{6, 9, 4, 0},
		ptr_range{12, 13, 1, 0},
		ptr_range{14, 15, 2, 0},
	}
	if !reflect.DeepEqual(r.ptr_range, want) {
		t.Errorf("Unmatched:\n   got: %#v\nexpect: %#v\n", r.ptr_range, want)
	}
}
