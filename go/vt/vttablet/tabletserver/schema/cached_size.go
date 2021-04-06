/*
Copyright 2021 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by Sizegen. DO NOT EDIT.

package schema

func (cached *MessageInfo) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(80)
	}
	// field Fields []*vitess.io/vitess/go/vt/proto/query.Field
	{
		size += int64(cap(cached.Fields)) * int64(8)
		for _, elem := range cached.Fields {
			size += elem.CachedSize(true)
		}
	}
	return size
}
func (cached *SequenceInfo) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(24)
	}
	return size
}
func (cached *Table) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(104)
	}
	// field Name vitess.io/vitess/go/vt/sqlparser.TableIdent
	size += cached.Name.CachedSize(false)
	// field Fields []*vitess.io/vitess/go/vt/proto/query.Field
	{
		size += int64(cap(cached.Fields)) * int64(8)
		for _, elem := range cached.Fields {
			size += elem.CachedSize(true)
		}
	}
	// field PKColumns []int
	{
		size += int64(cap(cached.PKColumns)) * int64(8)
	}
	// field SequenceInfo *vitess.io/vitess/go/vt/vttablet/tabletserver/schema.SequenceInfo
	size += cached.SequenceInfo.CachedSize(true)
	// field MessageInfo *vitess.io/vitess/go/vt/vttablet/tabletserver/schema.MessageInfo
	size += cached.MessageInfo.CachedSize(true)
	return size
}