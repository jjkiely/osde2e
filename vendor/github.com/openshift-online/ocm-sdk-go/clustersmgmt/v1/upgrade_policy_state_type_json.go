/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalUpgradePolicyState writes a value of the 'upgrade_policy_state' type to the given writer.
func MarshalUpgradePolicyState(object *UpgradePolicyState, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeUpgradePolicyState(object, stream)
	stream.Flush()
	return stream.Error
}

// writeUpgradePolicyState writes a value of the 'upgrade_policy_state' type to the given stream.
func writeUpgradePolicyState(object *UpgradePolicyState, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	if count > 0 {
		stream.WriteMore()
	}
	stream.WriteObjectField("kind")
	if object.link {
		stream.WriteString(UpgradePolicyStateLinkKind)
	} else {
		stream.WriteString(UpgradePolicyStateKind)
	}
	count++
	if object.id != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("id")
		stream.WriteString(*object.id)
		count++
	}
	if object.href != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("href")
		stream.WriteString(*object.href)
		count++
	}
	if object.description != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("description")
		stream.WriteString(*object.description)
		count++
	}
	if object.value != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("value")
		stream.WriteString(*object.value)
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalUpgradePolicyState reads a value of the 'upgrade_policy_state' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalUpgradePolicyState(source interface{}) (object *UpgradePolicyState, err error) {
	if source == http.NoBody {
		return
	}
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readUpgradePolicyState(iterator)
	err = iterator.Error
	return
}

// readUpgradePolicyState reads a value of the 'upgrade_policy_state' type from the given iterator.
func readUpgradePolicyState(iterator *jsoniter.Iterator) *UpgradePolicyState {
	object := &UpgradePolicyState{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "kind":
			value := iterator.ReadString()
			object.link = value == UpgradePolicyStateLinkKind
		case "id":
			value := iterator.ReadString()
			object.id = &value
		case "href":
			value := iterator.ReadString()
			object.href = &value
		case "description":
			value := iterator.ReadString()
			object.description = &value
		case "value":
			value := iterator.ReadString()
			object.value = &value
		default:
			iterator.ReadAny()
		}
	}
	return object
}
