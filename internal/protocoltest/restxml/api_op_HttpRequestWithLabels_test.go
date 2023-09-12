// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	protocoltesthttp "github.com/aws/aws-sdk-go-v2/internal/protocoltest"
	"github.com/aws/smithy-go/middleware"
	smithyprivateprotocol "github.com/aws/smithy-go/private/protocol"
	"github.com/aws/smithy-go/ptr"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithytime "github.com/aws/smithy-go/time"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestClient_HttpRequestWithLabels_awsRestxmlSerialize(t *testing.T) {
	cases := map[string]struct {
		Params        *HttpRequestWithLabelsInput
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
		// Sends a GET request that uses URI label bindings
		"InputWithHeadersAndAllParams": {
			Params: &HttpRequestWithLabelsInput{
				String_:   ptr.String("string"),
				Short:     ptr.Int16(1),
				Integer:   ptr.Int32(2),
				Long:      ptr.Int64(3),
				Float:     ptr.Float32(4.1),
				Double:    ptr.Float64(5.1),
				Boolean:   ptr.Bool(true),
				Timestamp: ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
			},
			ExpectMethod:  "GET",
			ExpectURIPath: "/HttpRequestWithLabels/string/1/2/3/4.1/5.1/true/2019-12-16T23%3A48%3A18Z",
			ExpectQuery:   []smithytesting.QueryItem{},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Sends a GET request that uses URI label bindings
		"HttpRequestLabelEscaping": {
			Params: &HttpRequestWithLabelsInput{
				String_:   ptr.String(" %:/?#[]@!$&'()*+,;=😹"),
				Short:     ptr.Int16(1),
				Integer:   ptr.Int32(2),
				Long:      ptr.Int64(3),
				Float:     ptr.Float32(4.1),
				Double:    ptr.Float64(5.1),
				Boolean:   ptr.Bool(true),
				Timestamp: ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
			},
			ExpectMethod:  "GET",
			ExpectURIPath: "/HttpRequestWithLabels/%20%25%3A%2F%3F%23%5B%5D%40%21%24%26%27%28%29%2A%2B%2C%3B%3D%F0%9F%98%B9/1/2/3/4.1/5.1/true/2019-12-16T23%3A48%3A18Z",
			ExpectQuery:   []smithytesting.QueryItem{},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
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
				HTTPClient:               protocoltesthttp.NewClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			result, err := client.HttpRequestWithLabels(context.Background(), c.Params, func(options *Options) {
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
