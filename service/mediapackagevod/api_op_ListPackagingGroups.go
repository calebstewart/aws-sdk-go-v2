// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackagevod

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/ListPackagingGroupsRequest
type ListPackagingGroupsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListPackagingGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPackagingGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPackagingGroupsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPackagingGroupsInput) MarshalFields(e protocol.FieldEncoder) error {
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
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/ListPackagingGroupsResponse
type ListPackagingGroupsOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `locationName:"nextToken" type:"string"`

	PackagingGroups []PackagingGroup `locationName:"packagingGroups" type:"list"`
}

// String returns the string representation
func (s ListPackagingGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPackagingGroupsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.PackagingGroups) > 0 {
		v := s.PackagingGroups

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "packagingGroups", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListPackagingGroups = "ListPackagingGroups"

// ListPackagingGroupsRequest returns a request value for making API operation for
// AWS Elemental MediaPackage VOD.
//
// Returns a collection of MediaPackage VOD PackagingGroup resources.
//
//    // Example sending a request using ListPackagingGroupsRequest.
//    req := client.ListPackagingGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/ListPackagingGroups
func (c *Client) ListPackagingGroupsRequest(input *ListPackagingGroupsInput) ListPackagingGroupsRequest {
	op := &aws.Operation{
		Name:       opListPackagingGroups,
		HTTPMethod: "GET",
		HTTPPath:   "/packaging_groups",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListPackagingGroupsInput{}
	}

	req := c.newRequest(op, input, &ListPackagingGroupsOutput{})
	return ListPackagingGroupsRequest{Request: req, Input: input, Copy: c.ListPackagingGroupsRequest}
}

// ListPackagingGroupsRequest is the request type for the
// ListPackagingGroups API operation.
type ListPackagingGroupsRequest struct {
	*aws.Request
	Input *ListPackagingGroupsInput
	Copy  func(*ListPackagingGroupsInput) ListPackagingGroupsRequest
}

// Send marshals and sends the ListPackagingGroups API request.
func (r ListPackagingGroupsRequest) Send(ctx context.Context) (*ListPackagingGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPackagingGroupsResponse{
		ListPackagingGroupsOutput: r.Request.Data.(*ListPackagingGroupsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPackagingGroupsRequestPaginator returns a paginator for ListPackagingGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPackagingGroupsRequest(input)
//   p := mediapackagevod.NewListPackagingGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPackagingGroupsPaginator(req ListPackagingGroupsRequest) ListPackagingGroupsPaginator {
	return ListPackagingGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListPackagingGroupsInput
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

// ListPackagingGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPackagingGroupsPaginator struct {
	aws.Pager
}

func (p *ListPackagingGroupsPaginator) CurrentPage() *ListPackagingGroupsOutput {
	return p.Pager.CurrentPage().(*ListPackagingGroupsOutput)
}

// ListPackagingGroupsResponse is the response type for the
// ListPackagingGroups API operation.
type ListPackagingGroupsResponse struct {
	*ListPackagingGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPackagingGroups request.
func (r *ListPackagingGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
