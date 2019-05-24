// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/ListMembersInput
type ListMembersInput struct {
	_ struct{} `type:"structure"`

	// An optional Boolean value. If provided, the request is limited either to
	// members that the current AWS account owns (true) or that other AWS accounts
	// own (false). If omitted, all members are listed.
	IsOwned *bool `location:"querystring" locationName:"isOwned" type:"boolean"`

	// The maximum number of members to return in the request.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The optional name of the member to list.
	Name *string `location:"querystring" locationName:"name" type:"string"`

	// The unique identifier of the network for which to list members.
	//
	// NetworkId is a required field
	NetworkId *string `location:"uri" locationName:"networkId" min:"1" type:"string" required:"true"`

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// An optional status specifier. If provided, only members currently in this
	// status are listed.
	Status MemberStatus `location:"querystring" locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListMembersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMembersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMembersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.NetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkId"))
	}
	if s.NetworkId != nil && len(*s.NetworkId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NetworkId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMembersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.NetworkId != nil {
		v := *s.NetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "networkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IsOwned != nil {
		v := *s.IsOwned

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "isOwned", protocol.BoolValue(v), metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/ListMembersOutput
type ListMembersOutput struct {
	_ struct{} `type:"structure"`

	// An array of MemberSummary objects. Each object contains details about a network
	// member.
	Members []MemberSummary `type:"list"`

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListMembersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMembersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Members) > 0 {
		v := s.Members

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Members", metadata)
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

const opListMembers = "ListMembers"

// ListMembersRequest returns a request value for making API operation for
// Amazon Managed Blockchain.
//
// Returns a listing of the members in a network and properties of their configurations.
//
//    // Example sending a request using ListMembersRequest.
//    req := client.ListMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/ListMembers
func (c *Client) ListMembersRequest(input *ListMembersInput) ListMembersRequest {
	op := &aws.Operation{
		Name:       opListMembers,
		HTTPMethod: "GET",
		HTTPPath:   "/networks/{networkId}/members",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListMembersInput{}
	}

	req := c.newRequest(op, input, &ListMembersOutput{})
	return ListMembersRequest{Request: req, Input: input, Copy: c.ListMembersRequest}
}

// ListMembersRequest is the request type for the
// ListMembers API operation.
type ListMembersRequest struct {
	*aws.Request
	Input *ListMembersInput
	Copy  func(*ListMembersInput) ListMembersRequest
}

// Send marshals and sends the ListMembers API request.
func (r ListMembersRequest) Send(ctx context.Context) (*ListMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMembersResponse{
		ListMembersOutput: r.Request.Data.(*ListMembersOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListMembersRequestPaginator returns a paginator for ListMembers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListMembersRequest(input)
//   p := managedblockchain.NewListMembersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListMembersPaginator(req ListMembersRequest) ListMembersPaginator {
	return ListMembersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListMembersInput
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

// ListMembersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListMembersPaginator struct {
	aws.Pager
}

func (p *ListMembersPaginator) CurrentPage() *ListMembersOutput {
	return p.Pager.CurrentPage().(*ListMembersOutput)
}

// ListMembersResponse is the response type for the
// ListMembers API operation.
type ListMembersResponse struct {
	*ListMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMembers request.
func (r *ListMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
