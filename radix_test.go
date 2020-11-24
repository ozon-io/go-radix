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

import "encoding/binary"
import "bufio"
import "fmt"
import "math/rand"
import "net"
import "os"
import "strconv"
import "strings"
import "testing"
import "time"

func display_node(n *Node, level int, branch string) {
	var typ string
	var ip net.IPNet
	var b []byte

	if n.Data != nil {
		typ = "LEAF"
	} else {
		typ = "NODE"
	}

	b = n.Bytes
	for len(b) < 4 {
		b = append([]byte{0x00}, b...)
	}
	ip.IP = net.IP(b)
	ip.Mask = net.CIDRMask(n.End + 1, 32)

	fmt.Printf("%s%s: %p/%s start=%d end=%d ip=%s\n", strings.Repeat("   ", level), branch, n, typ, n.Start, n.End, ip.String())
	if n.Left != nil {
		display_node(n.Left, level+1, "L")
	}
	if n.Right != nil {
		display_node(n.Right, level+1, "R")
	}

}

func display_radix(r *Radix) {

	if r.Node == nil {
		fmt.Printf("root pointer nil\n")
		return
	}

	display_node(r.Node, 0, "-")
}

func TestRadix(t *testing.T) {
	/*
	var r *Radix
	var b []byte
	var n []*Node

	r = NewRadix(true)

	println("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")
	b = []byte{0x00, 0x00, 0x00, 0b00010101}
	r.Insert(b, 32, nil)
	display_radix(r)

	println("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")
	b = []byte{0x00, 0x00, 0x00, 0b00010011}
	r.Insert(b, 32, nil)
	display_radix(r)

	println("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")
	b = []byte{0x00, 0x00, 0x00, 0b00010000}
	r.Insert(b, 29, nil)
	display_radix(r)

	println("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")
	b = []byte{0x00, 0x00, 0x00, 0b00000101}
	r.Insert(b, 32, nil)
	display_radix(r)

	println("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")
	b = []byte{0x00, 0x00, 0x00, 0b00000101}
	r.Insert(b, 32, nil)
	display_radix(r)

	println("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")
	b = []byte{0x00, 0x00, 0x00, 0b00001100}
	r.Insert(b, 30, nil)
	display_radix(r)

	println("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")
	b = []byte{0x00, 0x00, 0x00, 0b00001100}
	r.Insert(b, 32, nil)
	display_radix(r)

	println("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")
	b = []byte{0x00, 0x00, 0x00, 0b00001000}
	r.Insert(b, 29, nil)
	display_radix(r)

	println("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")
	b = []byte{0x00, 0x00, 0x00, 0b00010101}
	n = lookup_longuest(r, b, 32, true)
	fmt.Printf("%+v\n", n)
	*/
}

type ent struct {
	b []byte
	l int
}

func Benchmark_Radix(t *testing.B) {
	var r *Radix
	var file *os.File
	var scanner *bufio.Scanner
	var err error
	var tokens []string
	var ip uint32
	var bytes []byte
	var int_dec int64
	var now time.Time
	var step time.Time
	var count int
	var list []ent
	var i int
	var ent ent
	var rounds = 1000000
	var node *Node
	var ip3 net.IPNet
	var ip2 net.IPNet
	var b []byte
	var hit int
	var miss int

	r = NewRadix()

	/* Load file data/ip.db */

	file, err = os.Open("data/ip.db")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	count = 0
	now = time.Now()
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {

		tokens = strings.Split(scanner.Text(), "|")

		int_dec, err = strconv.ParseInt(tokens[0], 10, 64)
		if err != nil {
			panic(err)
		}
		ip = uint32(int_dec)
		ent.b = make([]byte, 4)
		binary.BigEndian.PutUint32(ent.b, ip)

		int_dec, err = strconv.ParseInt(tokens[1], 10, 8)
		if err != nil {
			panic(err)
		}
		ent.l = int(int_dec)

		list = append(list, ent)

		count++
	}
	step = time.Now()
	fmt.Printf("Load %d entries in %fs\n", count, step.Sub((now)).Seconds())

	/* populate radix with file loaded */

	now = time.Now()
	for _, ent = range list {
		r.Insert(&ent.b, ent.l, nil)
	}
	step = time.Now()
	fmt.Printf("Index %d entries in %fs\n", count, step.Sub((now)).Seconds())

	/* Perform random generation to have a reference */

	bytes = make([]byte, 4)
	now = time.Now()
	for i = 0; i < rounds; i++ {
		binary.BigEndian.PutUint32(bytes, rand.Uint32())
	}
	step = time.Now()
	fmt.Printf("Generate %d random numbers in %fs\n", rounds, step.Sub((now)).Seconds())

	/* perform random lookup to bench algo */

	now = time.Now()
	hit = 0
	miss = 0
	for i = 0; i < rounds; i++ {
		binary.BigEndian.PutUint32(bytes, rand.Uint32())
		node = r.LookupLonguest(&bytes, 32)
		if node != nil {
			hit++
		} else {
			miss++
		}
	}
	step = time.Now()
	d := step.Sub((now)).Seconds()
	fmt.Printf("Generate %d lookup in %fs with %d hit, %d miss\n", rounds, d, hit, miss)
	fmt.Printf("Mean time %f / 1000000 = %fns\n", d, d * 1000.0)

	/* perform full scan */

	now = time.Now()
	node = r.Node
	for {
		node = node.Next()
		if node == nil {
			break
		}
		b = node.Bytes
		for len(b) < 4 {
			b = append([]byte{0x00}, b...)
		}
		ip2.IP = net.IP(b)
		ip2.Mask = net.CIDRMask(node.End + 1, 32)
//		fmt.Printf("%s\n", ip2.String())
	}
	step = time.Now()
	fmt.Printf("Dump all data in %fs\n", step.Sub((now)).Seconds())

	/* Return first entry */

	node = r.First()
	if node == nil {
		panic("first cannot be null")
	}
	b = node.Bytes
	for len(b) < 4 {
		b = append([]byte{0x00}, b...)
	}
	ip2.IP = net.IP(b)
	ip2.Mask = net.CIDRMask(node.End + 1, 32)
	fmt.Printf("first = %s\n", ip2.String())

	/* Return last entry */

	node = r.Last()
	if node == nil {
		panic("first cannot be null")
	}
	b = node.Bytes
	for len(b) < 4 {
		b = append([]byte{0x00}, b...)
	}
	ip2.IP = net.IP(b)
	ip2.Mask = net.CIDRMask(node.End + 1, 32)
	fmt.Printf("last = %s\n", ip2.String())

	/* Returrn all cildrens of key 255.255.224.0/20 */

	var it *Iter
	var key []byte
	var ml int

//	key = []byte{0xff, 0xff, 0xe0, 0x00}
	key = []byte{0xff, 0xff, 0x80, 0x00}
//	key = []byte{0xd9, 0x14, 0x74, 0x88}
	ml = 18

	ip3.IP = net.IP(key)
	ip3.Mask = net.CIDRMask(ml, 32)
	it = r.NewIter(&key, ml)
	for it.Next() {
		node = it.Get()
		b = node.Bytes
		for len(b) < 4 {
			b = append([]byte{0x00}, b...)
		}
		ip2.IP = net.IP(b)
		ip2.Mask = net.CIDRMask(node.End + 1, 32)
		fmt.Printf("%s contains %s\n", ip3.String(), ip2.String())
	}

	count = 0
	key = []byte{}
	ml = 0
	now = time.Now()
	it = r.NewIter(&key, ml)
	for it.Next() {
		node = it.Get()
		r.Delete(node)
		count++
	}
	step = time.Now()
	fmt.Printf("Delete %d data in %fs\n", count, step.Sub((now)).Seconds())
}

func Test_Radix(t *testing.T) {
	var r *Radix
	var ipn *net.IPNet
	var it *Iter
	var n *Node

	/* Check error case */

	r = NewRadix()

	ipn = &net.IPNet{}
	ipn.IP = net.ParseIP("192.168.0.0")
	ipn.Mask = net.CIDRMask(16, 32)
	r.IPv4Insert(ipn, "Network 192.168.0.0/16")

	ipn = &net.IPNet{}
	ipn.IP = net.ParseIP("10.0.0.0")
	ipn.Mask = net.CIDRMask(8, 32)
	it = r.IPv4NewIter(ipn)
	for it.Next() {
		n = it.Get()
		t.Errorf("Case: we have one entry in the tree, initiate browsing on unconcerned network, " +
		         "expect 0 iteration. Got one entry: %s", n.String())
	}

	ipn = &net.IPNet{}
	ipn.IP = net.ParseIP("10.0.0.0")
	ipn.Mask = net.CIDRMask(8, 32)
	r.IPv4Insert(ipn, "Network 10.0.0.0/8")

	ipn = &net.IPNet{}
	ipn.IP = net.ParseIP("10.0.0.0")
	ipn.Mask = net.CIDRMask(8, 32)
	it = r.IPv4NewIter(ipn)
	n = nil
	for it.Next() {
		if n != nil {
			t.Errorf("Case: we have two entries in the tree, initiate browsing on concerned " +
			         "network, expect 1 iteration. Got at least two")
		}
		n = it.Get()
	}
}

func Test_Equal(t *testing.T) {
	var n1 Node
	var n2 Node

	/* test equal */
	n1.Bytes = []byte{0,0,0,0}
	n1.Start = 0
	n1.End = 31

	n2.Bytes = []byte{0,0,0,0}
	n2.Start = 0
	n2.End = 31

	if !Equal(&n1, &n2) {
		t.Errorf("Should be equal")
	}

	n2.End = 30
	if Equal(&n1, &n2) {
		t.Errorf("Should be different")
	}

	n2.Bytes[2] = 1
	n2.End = 31
	if Equal(&n1, &n2) {
		t.Errorf("Should be different")
	}
}
