// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package xray

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/xray-2016-04-12/GetGroupsRequest
type GetGroupsInput struct {
	_ struct{} `type:"structure"`

	// Pagination token. Not used.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetGroupsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetGroupsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/xray-2016-04-12/GetGroupsResult
type GetGroupsOutput struct {
	_ struct{} `type:"structure"`

	// The collection of all active groups.
	Groups []GroupSummary `type:"list"`

	// Pagination token. Not used.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetGroupsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Groups) > 0 {
		v := s.Groups

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Groups", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetGroups = "GetGroups"

// GetGroupsRequest returns a request value for making API operation for
// AWS X-Ray.
//
// Retrieves all active group details.
//
//    // Example sending a request using GetGroupsRequest.
//    req := client.GetGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/xray-2016-04-12/GetGroups
func (c *Client) GetGroupsRequest(input *GetGroupsInput) GetGroupsRequest {
	op := &aws.Operation{
		Name:       opGetGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/Groups",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetGroupsInput{}
	}

	req := c.newRequest(op, input, &GetGroupsOutput{})
	return GetGroupsRequest{Request: req, Input: input, Copy: c.GetGroupsRequest}
}

// GetGroupsRequest is the request type for the
// GetGroups API operation.
type GetGroupsRequest struct {
	*aws.Request
	Input *GetGroupsInput
	Copy  func(*GetGroupsInput) GetGroupsRequest
}

// Send marshals and sends the GetGroups API request.
func (r GetGroupsRequest) Send(ctx context.Context) (*GetGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetGroupsResponse{
		GetGroupsOutput: r.Request.Data.(*GetGroupsOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetGroupsRequestPaginator returns a paginator for GetGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetGroupsRequest(input)
//   p := xray.NewGetGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetGroupsPaginator(req GetGroupsRequest) GetGroupsPaginator {
	return GetGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetGroupsInput
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

// GetGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetGroupsPaginator struct {
	aws.Pager
}

func (p *GetGroupsPaginator) CurrentPage() *GetGroupsOutput {
	return p.Pager.CurrentPage().(*GetGroupsOutput)
}

// GetGroupsResponse is the response type for the
// GetGroups API operation.
type GetGroupsResponse struct {
	*GetGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetGroups request.
func (r *GetGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
