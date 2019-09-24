/*
 * Copyright (c) 2016 Yanko Bolanos
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 */
package main // import "github.com/rem7/gotail"

// Sample program using tail library

import (
	"context"
	"fmt"
	"os"
	"sync"
	"unicode/utf8"

	"github.com/rem7/tail/v2"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("need one arg")
		os.Exit(-1)
	}

	// fls := regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}-\d{2}:\d{2}\s`)

	var wg sync.WaitGroup
	files := os.Args[1:]
	for _, file := range files {
		wg.Add(1)
		go func(path string) {

			t := tail.NewTailWithCtx(context.Background(), path, true, true, nil, false, false)
			n := 1

			for {
				line, open := <-t.LineChan
				if !open {
					break
				}
				if ok := utf8.ValidString(line); !ok {
					fmt.Print("line %d not UTF-8: ", n)
				}
				fmt.Println(line)
				n += 1

			}
			wg.Done()
		}(file)
	}
	wg.Wait()
}
