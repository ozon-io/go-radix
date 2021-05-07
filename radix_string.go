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

func string_to_key(str string)([]byte, int) {
	return []byte(str), len(str) * 8
}

func (r *Radix)StringLookupLonguest(str string)(*Node) {
	var length int
	var key []byte

	/* Get the network width. width of 0 id prohibited */
	key, length = string_to_key(str)
	if length == 0 {
		return nil
	}

	/* Perform lookup */
	return r.LookupLonguest(&key, length)
}

func (r *Radix)StringLookupLonguestPath(str string)([]*Node) {
	var length int
	var key []byte

	/* Get the network width. width of 0 id prohibited */
	key, length = string_to_key(str)
	if length == 0 {
		return make([]*Node, 0)
	}

	/* Perform lookup */
	return r.LookupLonguestPath(&key, length)
}

func (r *Radix)StringGet(str string)(*Node) {
	var length int
	var key []byte

	/* Get the network width. width of 0 id prohibited */
	key, length = string_to_key(str)
	if length == 0 {
		return nil
	}

	/* Perform lookup */
	return r.Get(&key, length)
}

func (r *Radix)StringInsert(str string, data interface{})(*Node, bool) {
	var length int
	var key []byte

	/* Get the network width. width of 0 id prohibited */
	key, length = string_to_key(str)
	if length == 0 {
		return nil, false
	}

	/* Perform insert */
	return r.Insert(&key, length, data)
}

func (r *Radix)StringDelete(str string)() {
	var node *Node
	var length int
	var key []byte

	/* Get the network width. width of 0 id prohibited */
	key, length = string_to_key(str)
	if length == 0 {
		return
	}

	/* Perform lookup */
	node = r.Get(&key, length)
	if node == nil {
		return
	}

	/* Delete entry */
	r.Delete(node)
}

func (r *Radix)StringNewIter(str string)(*Iter) {
	var length int
	var key []byte

	key, length = string_to_key(str)
	return r.NewIter(&key, length)
}

func (n *Node)StringGetKey()(string) {
	return string(n.Bytes)
}
