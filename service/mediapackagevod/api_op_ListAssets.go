// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackagevod

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/ListAssetsRequest
type ListAssetsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	PackagingGroupId *string `location:"querystring" locationName:"packagingGroupId" type:"string"`
}

// String returns the string representation
func (s ListAssetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAssetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAssetsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAssetsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PackagingGroupId != nil {
		v := *s.PackagingGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "packagingGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/ListAssetsResponse
type ListAssetsOutput struct {
	_ struct{} `type:"structure"`

	Assets []AssetShallow `locationName:"assets" type:"list"`

	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListAssetsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAssetsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Assets) > 0 {
		v := s.Assets

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "assets", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListAssets = "ListAssets"

// ListAssetsRequest returns a request value for making API operation for
// AWS Elemental MediaPackage VOD.
//
// Returns a collection of MediaPackage VOD Asset resources.
//
//    // Example sending a request using ListAssetsRequest.
//    req := client.ListAssetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/ListAssets
func (c *Client) ListAssetsRequest(input *ListAssetsInput) ListAssetsRequest {
	op := &aws.Operation{
		Name:       opListAssets,
		HTTPMethod: "GET",
		HTTPPath:   "/assets",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListAssetsInput{}
	}

	req := c.newRequest(op, input, &ListAssetsOutput{})
	return ListAssetsRequest{Request: req, Input: input, Copy: c.ListAssetsRequest}
}

// ListAssetsRequest is the request type for the
// ListAssets API operation.
type ListAssetsRequest struct {
	*aws.Request
	Input *ListAssetsInput
	Copy  func(*ListAssetsInput) ListAssetsRequest
}

// Send marshals and sends the ListAssets API request.
func (r ListAssetsRequest) Send(ctx context.Context) (*ListAssetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAssetsResponse{
		ListAssetsOutput: r.Request.Data.(*ListAssetsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAssetsRequestPaginator returns a paginator for ListAssets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAssetsRequest(input)
//   p := mediapackagevod.NewListAssetsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAssetsPaginator(req ListAssetsRequest) ListAssetsPaginator {
	return ListAssetsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListAssetsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListAssetsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAssetsPaginator struct {
	aws.Pager
}

func (p *ListAssetsPaginator) CurrentPage() *ListAssetsOutput {
	return p.Pager.CurrentPage().(*ListAssetsOutput)
}

// ListAssetsResponse is the response type for the
// ListAssets API operation.
type ListAssetsResponse struct {
	*ListAssetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAssets request.
func (r *ListAssetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
