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

import "net"
import "testing"

func Example_ipv4() {

	// Create new tree root
	r := NewRadix()

	// Insert first network
	_, n1, _ := net.ParseCIDR("10.0.0.0/16") 
	r.IPv4Insert(n1, "This is the first network inserted")

	// Lookup the network
	_, n2, _ := net.ParseCIDR("10.0.0.33/32") 
	node1 := r.IPv4LookupLonguest(n2)
	if node1 != nil {
		println("network", n2.String(), "is contained in network", node1.IPv4GetNet().String())
		println("network", node1.IPv4GetNet().String(), "is associated with data", node1.Data.(string))
	}

	// Lookup too large network
	_, n3, _ := net.ParseCIDR("10.0.0.0/8") 
	node2 := r.IPv4LookupLonguest(n3)
	if node2 == nil {
		println("network", n3.String(), "has no entries in the tree")
	}
}

func Example_string() {
	// Create new tree root
	r := NewRadix()

	// insert string
	r.StringInsert("home", "This is a prefix")

	// lookup word
	node1 := r.StringLookupLonguest("homemade")
	if node1 != nil {
		println("homemade has prefix", node1.StringGetKey(), "in the tree, with data", node1.Data.(string))
	}
}

func Test_examples(t *testing.T) {
	Example_ipv4()
	Example_string()
}
