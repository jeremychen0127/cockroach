// Copyright 2016 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.
//
// Author: Radu Berinde (radu@cockroachlabs.com)

package distsql

import "github.com/pkg/errors"

func makeRouter(typ OutputRouterSpec_Type, streams []RowReceiver) (
	RowReceiver, error,
) {
	switch len(streams) {
	case 0:
		return nil, errors.Errorf("no streams in router")
	case 1:
		// Special passthrough case - no router.
		return streams[0], nil
	default:
		return nil, errors.Errorf("router type %s not supported", typ)
	}
}
