// Copyright 2018 Nuno Preguica, NOVA LINCS, FCT, Universidade NOVA de Lisboa.
// All rights reserved.
// Use of this source code is governed by Apache 2.0
// license that can be found in the LICENSE file.
package opcrdts

type CRDTOperation interface {
	GetCRDTType() byte
	GetType() byte
	Merge(CRDTOperation) (CRDTOperation, bool)
	Serialize() ([]byte, bool)
	Apply(CRDT) bool
}
