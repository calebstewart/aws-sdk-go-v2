// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	protocoltesthttp "github.com/aws/aws-sdk-go-v2/internal/protocoltest"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/jsonrpc/types"
	smithydocument "github.com/aws/smithy-go/document"
	"github.com/aws/smithy-go/middleware"
	smithyprivateprotocol "github.com/aws/smithy-go/private/protocol"
	"github.com/aws/smithy-go/ptr"
	smithytesting "github.com/aws/smithy-go/testing"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"testing"
)

func TestClient_KitchenSinkOperation_awsAwsjson11Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *KitchenSinkOperationInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		Host          *url.URL
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Serializes string shapes
		"serializes_string_shapes": {
			Params: &KitchenSinkOperationInput{
				String_: ptr.String("abc xyz"),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"String":"abc xyz"}`))
			},
		},
		// Serializes string shapes with jsonvalue trait
		"serializes_string_shapes_with_jsonvalue_trait": {
			Params: &KitchenSinkOperationInput{
				JsonValue: ptr.String("{\"string\":\"value\",\"number\":1234.5,\"boolTrue\":true,\"boolFalse\":false,\"array\":[1,2,3,4],\"object\":{\"key\":\"value\"},\"null\":null}"),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"JsonValue":"{\"string\":\"value\",\"number\":1234.5,\"boolTrue\":true,\"boolFalse\":false,\"array\":[1,2,3,4],\"object\":{\"key\":\"value\"},\"null\":null}"}`))
			},
		},
		// Serializes integer shapes
		"serializes_integer_shapes": {
			Params: &KitchenSinkOperationInput{
				Integer: ptr.Int32(1234),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"Integer":1234}`))
			},
		},
		// Serializes long shapes
		"serializes_long_shapes": {
			Params: &KitchenSinkOperationInput{
				Long: ptr.Int64(999999999999),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"Long":999999999999}`))
			},
		},
		// Serializes float shapes
		"serializes_float_shapes": {
			Params: &KitchenSinkOperationInput{
				Float: ptr.Float32(1234.5),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"Float":1234.5}`))
			},
		},
		// Serializes double shapes
		"serializes_double_shapes": {
			Params: &KitchenSinkOperationInput{
				Double: ptr.Float64(1234.5),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"Double":1234.5}`))
			},
		},
		// Serializes blob shapes
		"serializes_blob_shapes": {
			Params: &KitchenSinkOperationInput{
				Blob: []byte("binary-value"),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"Blob":"YmluYXJ5LXZhbHVl"}`))
			},
		},
		// Serializes boolean shapes (true)
		"serializes_boolean_shapes_true": {
			Params: &KitchenSinkOperationInput{
				Boolean: ptr.Bool(true),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"Boolean":true}`))
			},
		},
		// Serializes boolean shapes (false)
		"serializes_boolean_shapes_false": {
			Params: &KitchenSinkOperationInput{
				Boolean: ptr.Bool(false),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"Boolean":false}`))
			},
		},
		// Serializes timestamp shapes
		"serializes_timestamp_shapes": {
			Params: &KitchenSinkOperationInput{
				Timestamp: ptr.Time(smithytime.ParseEpochSeconds(946845296)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"Timestamp":946845296}`))
			},
		},
		// Serializes timestamp shapes with iso8601 timestampFormat
		"serializes_timestamp_shapes_with_iso8601_timestampformat": {
			Params: &KitchenSinkOperationInput{
				Iso8601Timestamp: ptr.Time(smithytime.ParseEpochSeconds(946845296)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"Iso8601Timestamp":"2000-01-02T20:34:56Z"}`))
			},
		},
		// Serializes timestamp shapes with httpdate timestampFormat
		"serializes_timestamp_shapes_with_httpdate_timestampformat": {
			Params: &KitchenSinkOperationInput{
				HttpdateTimestamp: ptr.Time(smithytime.ParseEpochSeconds(946845296)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"HttpdateTimestamp":"Sun, 02 Jan 2000 20:34:56 GMT"}`))
			},
		},
		// Serializes timestamp shapes with unixTimestamp timestampFormat
		"serializes_timestamp_shapes_with_unixtimestamp_timestampformat": {
			Params: &KitchenSinkOperationInput{
				UnixTimestamp: ptr.Time(smithytime.ParseEpochSeconds(946845296)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"UnixTimestamp":946845296}`))
			},
		},
		// Serializes list shapes
		"serializes_list_shapes": {
			Params: &KitchenSinkOperationInput{
				ListOfStrings: []string{
					"abc",
					"mno",
					"xyz",
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"ListOfStrings":["abc","mno","xyz"]}`))
			},
		},
		// Serializes empty list shapes
		"serializes_empty_list_shapes": {
			Params: &KitchenSinkOperationInput{
				ListOfStrings: []string{},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"ListOfStrings":[]}`))
			},
		},
		// Serializes list of map shapes
		"serializes_list_of_map_shapes": {
			Params: &KitchenSinkOperationInput{
				ListOfMapsOfStrings: []map[string]string{
					{
						"foo": "bar",
					},
					{
						"abc": "xyz",
					},
					{
						"red": "blue",
					},
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"ListOfMapsOfStrings":[{"foo":"bar"},{"abc":"xyz"},{"red":"blue"}]}`))
			},
		},
		// Serializes list of structure shapes
		"serializes_list_of_structure_shapes": {
			Params: &KitchenSinkOperationInput{
				ListOfStructs: []types.SimpleStruct{
					{
						Value: ptr.String("abc"),
					},
					{
						Value: ptr.String("mno"),
					},
					{
						Value: ptr.String("xyz"),
					},
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"ListOfStructs":[{"Value":"abc"},{"Value":"mno"},{"Value":"xyz"}]}`))
			},
		},
		// Serializes list of recursive structure shapes
		"serializes_list_of_recursive_structure_shapes": {
			Params: &KitchenSinkOperationInput{
				RecursiveList: []types.KitchenSink{
					{
						RecursiveList: []types.KitchenSink{
							{
								RecursiveList: []types.KitchenSink{
									{
										Integer: ptr.Int32(123),
									},
								},
							},
						},
					},
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"RecursiveList":[{"RecursiveList":[{"RecursiveList":[{"Integer":123}]}]}]}`))
			},
		},
		// Serializes map shapes
		"serializes_map_shapes": {
			Params: &KitchenSinkOperationInput{
				MapOfStrings: map[string]string{
					"abc": "xyz",
					"mno": "hjk",
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"MapOfStrings":{"abc":"xyz","mno":"hjk"}}`))
			},
		},
		// Serializes empty map shapes
		"serializes_empty_map_shapes": {
			Params: &KitchenSinkOperationInput{
				MapOfStrings: map[string]string{},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"MapOfStrings":{}}`))
			},
		},
		// Serializes map of list shapes
		"serializes_map_of_list_shapes": {
			Params: &KitchenSinkOperationInput{
				MapOfListsOfStrings: map[string][]string{
					"abc": {
						"abc",
						"xyz",
					},
					"mno": {
						"xyz",
						"abc",
					},
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"MapOfListsOfStrings":{"abc":["abc","xyz"],"mno":["xyz","abc"]}}`))
			},
		},
		// Serializes map of structure shapes
		"serializes_map_of_structure_shapes": {
			Params: &KitchenSinkOperationInput{
				MapOfStructs: map[string]types.SimpleStruct{
					"key1": {
						Value: ptr.String("value-1"),
					},
					"key2": {
						Value: ptr.String("value-2"),
					},
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"MapOfStructs":{"key1":{"Value":"value-1"},"key2":{"Value":"value-2"}}}`))
			},
		},
		// Serializes map of recursive structure shapes
		"serializes_map_of_recursive_structure_shapes": {
			Params: &KitchenSinkOperationInput{
				RecursiveMap: map[string]types.KitchenSink{
					"key1": {
						RecursiveMap: map[string]types.KitchenSink{
							"key2": {
								RecursiveMap: map[string]types.KitchenSink{
									"key3": {
										Boolean: ptr.Bool(false),
									},
								},
							},
						},
					},
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"RecursiveMap":{"key1":{"RecursiveMap":{"key2":{"RecursiveMap":{"key3":{"Boolean":false}}}}}}}`))
			},
		},
		// Serializes structure shapes
		"serializes_structure_shapes": {
			Params: &KitchenSinkOperationInput{
				SimpleStruct: &types.SimpleStruct{
					Value: ptr.String("abc"),
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"SimpleStruct":{"Value":"abc"}}`))
			},
		},
		// Serializes structure members with locationName traits
		"serializes_structure_members_with_locationname_traits": {
			Params: &KitchenSinkOperationInput{
				StructWithJsonName: &types.StructWithJsonName{
					Value: ptr.String("some-value"),
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"StructWithJsonName":{"Value":"some-value"}}`))
			},
		},
		// Serializes empty structure shapes
		"serializes_empty_structure_shapes": {
			Params: &KitchenSinkOperationInput{
				SimpleStruct: &types.SimpleStruct{},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"SimpleStruct":{}}`))
			},
		},
		// Serializes structure which have no members
		"serializes_structure_which_have_no_members": {
			Params: &KitchenSinkOperationInput{
				EmptyStruct: &types.EmptyStruct{},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"EmptyStruct":{}}`))
			},
		},
		// Serializes recursive structure shapes
		"serializes_recursive_structure_shapes": {
			Params: &KitchenSinkOperationInput{
				String_: ptr.String("top-value"),
				Boolean: ptr.Bool(false),
				RecursiveStruct: &types.KitchenSink{
					String_: ptr.String("nested-value"),
					Boolean: ptr.Bool(true),
					RecursiveList: []types.KitchenSink{
						{
							String_: ptr.String("string-only"),
						},
						{
							RecursiveStruct: &types.KitchenSink{
								MapOfStrings: map[string]string{
									"color": "red",
									"size":  "large",
								},
							},
						},
					},
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.KitchenSinkOperation"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{"String":"top-value","Boolean":false,"RecursiveStruct":{"String":"nested-value","Boolean":true,"RecursiveList":[{"String":"string-only"},{"RecursiveStruct":{"MapOfStrings":{"color":"red","size":"large"}}}]}}`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			actualReq := &http.Request{}
			serverURL := "http://localhost:8888/"
			if c.Host != nil {
				u, err := url.Parse(serverURL)
				if err != nil {
					t.Fatalf("expect no error, got %v", err)
				}
				u.Path = c.Host.Path
				u.RawPath = c.Host.RawPath
				u.RawQuery = c.Host.RawQuery
				serverURL = u.String()
			}
			client := New(Options{
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient: protocoltesthttp.NewClient(),
				Region:     "us-west-2",
			})
			result, err := client.KitchenSinkOperation(context.Background(), c.Params, func(options *Options) {
				options.APIOptions = append(options.APIOptions, func(stack *middleware.Stack) error {
					return smithyprivateprotocol.AddCaptureRequestMiddleware(stack, actualReq)
				})
			})
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}

func TestClient_KitchenSinkOperation_awsAwsjson11Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *KitchenSinkOperationOutput
	}{
		// Parses operations with empty JSON bodies
		"parses_operations_with_empty_json_bodies": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{}`),
			ExpectResult:  &KitchenSinkOperationOutput{},
		},
		// Parses string shapes
		"parses_string_shapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"String":"string-value"}`),
			ExpectResult: &KitchenSinkOperationOutput{
				String_: ptr.String("string-value"),
			},
		},
		// Parses integer shapes
		"parses_integer_shapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"Integer":1234}`),
			ExpectResult: &KitchenSinkOperationOutput{
				Integer: ptr.Int32(1234),
			},
		},
		// Parses long shapes
		"parses_long_shapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"Long":1234567890123456789}`),
			ExpectResult: &KitchenSinkOperationOutput{
				Long: ptr.Int64(1234567890123456789),
			},
		},
		// Parses float shapes
		"parses_float_shapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"Float":1234.5}`),
			ExpectResult: &KitchenSinkOperationOutput{
				Float: ptr.Float32(1234.5),
			},
		},
		// Parses double shapes
		"parses_double_shapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"Double":123456789.12345679}`),
			ExpectResult: &KitchenSinkOperationOutput{
				Double: ptr.Float64(1.2345678912345679e8),
			},
		},
		// Parses boolean shapes (true)
		"parses_boolean_shapes_true": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"Boolean":true}`),
			ExpectResult: &KitchenSinkOperationOutput{
				Boolean: ptr.Bool(true),
			},
		},
		// Parses boolean (false)
		"parses_boolean_false": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"Boolean":false}`),
			ExpectResult: &KitchenSinkOperationOutput{
				Boolean: ptr.Bool(false),
			},
		},
		// Parses blob shapes
		"parses_blob_shapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"Blob":"YmluYXJ5LXZhbHVl"}`),
			ExpectResult: &KitchenSinkOperationOutput{
				Blob: []byte("binary-value"),
			},
		},
		// Parses timestamp shapes
		"parses_timestamp_shapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"Timestamp":946845296}`),
			ExpectResult: &KitchenSinkOperationOutput{
				Timestamp: ptr.Time(smithytime.ParseEpochSeconds(946845296)),
			},
		},
		// Parses iso8601 timestamps
		"parses_iso8601_timestamps": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"Iso8601Timestamp":"2000-01-02T20:34:56Z"}`),
			ExpectResult: &KitchenSinkOperationOutput{
				Iso8601Timestamp: ptr.Time(smithytime.ParseEpochSeconds(946845296)),
			},
		},
		// Parses httpdate timestamps
		"parses_httpdate_timestamps": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"HttpdateTimestamp":"Sun, 02 Jan 2000 20:34:56 GMT"}`),
			ExpectResult: &KitchenSinkOperationOutput{
				HttpdateTimestamp: ptr.Time(smithytime.ParseEpochSeconds(946845296)),
			},
		},
		// Parses list shapes
		"parses_list_shapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"ListOfStrings":["abc","mno","xyz"]}`),
			ExpectResult: &KitchenSinkOperationOutput{
				ListOfStrings: []string{
					"abc",
					"mno",
					"xyz",
				},
			},
		},
		// Parses list of map shapes
		"parses_list_of_map_shapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"ListOfMapsOfStrings":[{"size":"large"},{"color":"red"}]}`),
			ExpectResult: &KitchenSinkOperationOutput{
				ListOfMapsOfStrings: []map[string]string{
					{
						"size": "large",
					},
					{
						"color": "red",
					},
				},
			},
		},
		// Parses list of list shapes
		"parses_list_of_list_shapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"ListOfLists":[["abc","mno","xyz"],["hjk","qrs","tuv"]]}`),
			ExpectResult: &KitchenSinkOperationOutput{
				ListOfLists: [][]string{
					{
						"abc",
						"mno",
						"xyz",
					},
					{
						"hjk",
						"qrs",
						"tuv",
					},
				},
			},
		},
		// Parses list of structure shapes
		"parses_list_of_structure_shapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"ListOfStructs":[{"Value":"value-1"},{"Value":"value-2"}]}`),
			ExpectResult: &KitchenSinkOperationOutput{
				ListOfStructs: []types.SimpleStruct{
					{
						Value: ptr.String("value-1"),
					},
					{
						Value: ptr.String("value-2"),
					},
				},
			},
		},
		// Parses list of recursive structure shapes
		"parses_list_of_recursive_structure_shapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"RecursiveList":[{"RecursiveList":[{"RecursiveList":[{"String":"value"}]}]}]}`),
			ExpectResult: &KitchenSinkOperationOutput{
				RecursiveList: []types.KitchenSink{
					{
						RecursiveList: []types.KitchenSink{
							{
								RecursiveList: []types.KitchenSink{
									{
										String_: ptr.String("value"),
									},
								},
							},
						},
					},
				},
			},
		},
		// Parses map shapes
		"parses_map_shapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"MapOfStrings":{"size":"large","color":"red"}}`),
			ExpectResult: &KitchenSinkOperationOutput{
				MapOfStrings: map[string]string{
					"size":  "large",
					"color": "red",
				},
			},
		},
		// Parses map of list shapes
		"parses_map_of_list_shapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"MapOfListsOfStrings":{"sizes":["large","small"],"colors":["red","green"]}}`),
			ExpectResult: &KitchenSinkOperationOutput{
				MapOfListsOfStrings: map[string][]string{
					"sizes": {
						"large",
						"small",
					},
					"colors": {
						"red",
						"green",
					},
				},
			},
		},
		// Parses map of map shapes
		"parses_map_of_map_shapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"MapOfMaps":{"sizes":{"large":"L","medium":"M"},"colors":{"red":"R","blue":"B"}}}`),
			ExpectResult: &KitchenSinkOperationOutput{
				MapOfMaps: map[string]map[string]string{
					"sizes": {
						"large":  "L",
						"medium": "M",
					},
					"colors": {
						"red":  "R",
						"blue": "B",
					},
				},
			},
		},
		// Parses map of structure shapes
		"parses_map_of_structure_shapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"MapOfStructs":{"size":{"Value":"small"},"color":{"Value":"red"}}}`),
			ExpectResult: &KitchenSinkOperationOutput{
				MapOfStructs: map[string]types.SimpleStruct{
					"size": {
						Value: ptr.String("small"),
					},
					"color": {
						Value: ptr.String("red"),
					},
				},
			},
		},
		// Parses map of recursive structure shapes
		"parses_map_of_recursive_structure_shapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{"RecursiveMap":{"key-1":{"RecursiveMap":{"key-2":{"RecursiveMap":{"key-3":{"String":"value"}}}}}}}`),
			ExpectResult: &KitchenSinkOperationOutput{
				RecursiveMap: map[string]types.KitchenSink{
					"key-1": {
						RecursiveMap: map[string]types.KitchenSink{
							"key-2": {
								RecursiveMap: map[string]types.KitchenSink{
									"key-3": {
										String_: ptr.String("value"),
									},
								},
							},
						},
					},
				},
			},
		},
		// Parses the request id from the response
		"parses_the_request_id_from_the_response": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type":     []string{"application/x-amz-json-1.1"},
				"X-Amzn-Requestid": []string{"amazon-uniq-request-id"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{}`),
			ExpectResult:  &KitchenSinkOperationOutput{},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			serverURL := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				Region: "us-west-2",
			})
			var params KitchenSinkOperationInput
			result, err := client.KitchenSinkOperation(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			opts := cmp.Options{
				cmpopts.IgnoreUnexported(
					middleware.Metadata{},
				),
				cmp.FilterValues(func(x, y float64) bool {
					return math.IsNaN(x) && math.IsNaN(y)
				}, cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmp.FilterValues(func(x, y float32) bool {
					return math.IsNaN(float64(x)) && math.IsNaN(float64(y))
				}, cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmpopts.IgnoreTypes(smithydocument.NoSerde{}),
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, opts...); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
