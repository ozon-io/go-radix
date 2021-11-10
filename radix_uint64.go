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

import "encoding/binary"

const length = 64

func uint64_to_key(value uint64)([]byte) {
	var bytes [8]byte

	binary.BigEndian.PutUint64(bytes[:], value)
	return bytes[:]
}

func (r *Radix)UInt64Get(value uint64)(*Node) {
	var key []byte

	/* Get the network width. width of 0 id prohibited */
	key = uint64_to_key(value)

	/* Perform lookup */
	return r.Get(&key, length)
}

func (r *Radix)UInt64Insert(value uint64, data interface{})(*Node, bool) {
	var key []byte

	/* Get the network width. width of 0 id prohibited */
	key = uint64_to_key(value)

	/* Perform insert */
	return r.Insert(&key, length, data)
}

func (r *Radix)UInt64Delete(value uint64)() {
	var node *Node
	var key []byte

	/* Get the network width. width of 0 id prohibited */
	key = uint64_to_key(value)

	/* Perform lookup */
	node = r.Get(&key, length)
	if node == nil {
		return
	}

	/* Delete entry */
	r.Delete(node)
}

func (n *Node)UInt64GetValue()(uint64) {
	if len(n.Bytes) != 8 {
		return 0
	}
	return binary.BigEndian.Uint64(n.Bytes)
}

func (r *Radix)UInt64NewIter(value uint64)(*Iter) {
	var key []byte

	key = uint64_to_key(value)
	return r.NewIter(&key, length)
}
