// Copyright 2014 BertWednesdays Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Simple command line parser test script
//

package args

import (
	"testing"
)

type basictest struct {
	in_setup  arguments_map
	in_params []string
	answer    arguments_map
}

var testSet = []interface{}{
	basictest{arguments_map{"key": "desc"}, []string{"--key", "value"}, arguments_map{"key": "value"}},
	basictest{arguments_map{"key": "desc1", "key1": "desc1", "key2": "desc2"}, []string{"--key", "value"}, arguments_map{"key": "value"}},
	basictest{arguments_map{"key": "desc1", "key1": "desc1", "key2": "desc2"}, []string{"--key1", "value"}, arguments_map{"key1": "value"}},
	basictest{arguments_map{"key": "desc1", "key1": "desc1", "key2": "desc2"}, []string{"--key2", "value"}, arguments_map{"key2": "value"}},
	basictest{arguments_map{"key": "desc1", "key1": "desc1", "key2": "desc2"}, []string{"--key1", "value1", "--key2", "value2"}, arguments_map{"key1": "value1", "key2": "value2"}},
	basictest{arguments_map{"key": "desc1", "key1": "desc1", "key2": "desc2"}, []string{"--key2", "value2", "--key1", "value1"}, arguments_map{"key1": "value1", "key2": "value2"}},
	basictest{arguments_map{"key": "desc1", "key1": "desc1", "key2": "desc2"}, []string{"--key2", "value2", "--key2", "value1"}, arguments_map{"key2": "value1"}},
}

func TestBuild(t *testing.T) {
	for _, element := range testSet {
		switch element.(type) {
		case basictest:
			e := element.(basictest)
			ans := New(e.in_setup, e.in_params)

			if ans.HasError() {
				t.Error("New returned error: " + ans.GetError())
			} else {
				if ans.Count() != len(e.answer) {
					t.Error("New returned incorrect number of elements ")
				} else {
					for key, val := range e.answer {
						if !ans.HasParam(key) {
							t.Error("key " + key + " is missing")
						} else if ans.GetParam(key) != val {
							t.Error("Value Missmatch : " + val + "!=" + ans.GetParam(key))
						}
					}
				}
			}
		}
	}
}
