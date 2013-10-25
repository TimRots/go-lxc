/*
 * concurrent_create.go
 *
 * Copyright © 2013, S.Çağlar Onur
 *
 * Authors:
 * S.Çağlar Onur <caglar@10ur.org>
 *
 * This library is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 2, as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along
 * with this program; if not, write to the Free Software Foundation, Inc.,
 * 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

package main

import (
	"fmt"
	"github.com/caglar10ur/lxc"
	"runtime"
	"strconv"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			c := lxc.NewContainer(strconv.Itoa(i))
			defer lxc.PutContainer(c)

			fmt.Printf("Creating the container (%d)...\n", i)
			if err := c.Create("busybox"); err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
