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

func Test_string(t *testing.T) {
	var r *Radix
	var n *Node
	var ns []*Node
	var s string

	/* Init DB */
	r = NewRadix()

	/* Insert string */
	r.StringInsert("aaaa", "key aaaa")
	r.StringInsert("aaa", "key aaa")
	r.StringInsert("aa", "key aa")

	/* Lookup exact */
	n = r.StringGet("aaaa")
	if n == nil {
		t.Errorf("aaaa should be found")
	} else {
		if n.StringGetKey() != "aaaa" {
			t.Errorf("aaaa should be found")
		}
		s, _ = n.Data.(string)
		if s != "key aaaa" {
			t.Errorf("\"key aaaa\" should be found")
		}
	}

	/* Lookup exact */
	n = r.StringGet("aa")
	if n == nil {
		t.Errorf("aa should be found")
	} else {
		if n.StringGetKey() != "aa" {
			t.Errorf("aa should be found")
		}
		s, _ = n.Data.(string)
		if s != "key aa" {
			t.Errorf("\"key aa\" should be found")
		}
	}

	/* lookup longest prefix */
	n = r.StringLookupLonguest("aaaa stayin alive")
	if n == nil {
		t.Errorf("aaaa should be found")
	} else {
		if n.StringGetKey() != "aaaa" {
			t.Errorf("aaaa should be found")
		}
		s, _ = n.Data.(string)
		if s != "key aaaa" {
			t.Errorf("\"key aaaa\" should be found")
		}
	}

	/* lookup longest prefix */
	n = r.StringLookupLonguest("aa stayin alive")
	if n == nil {
		t.Errorf("aa should be found")
	} else {
		if n.StringGetKey() != "aa" {
			t.Errorf("aa should be found")
		}
		s, _ = n.Data.(string)
		if s != "key aa" {
			t.Errorf("\"key aa\" should be found")
		}
	}

	/* lookup longest prefix */
	n = r.StringLookupLonguest("ar stayin alive")
	if n != nil {
		t.Errorf("lookup should be nil")
	}

	/* Lookup longest path */
	ns = r.StringLookupLonguestPath("aaaa")
	if len(ns) != 3 {
		t.Errorf("Expect 3 entries")
	}
}
