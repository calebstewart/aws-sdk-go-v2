// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/ListAliasesRequest
type ListAliasesInput struct {
	_ struct{} `type:"structure"`

	// Lists only aliases that refer to the specified CMK. The value of this parameter
	// can be the ID or Amazon Resource Name (ARN) of a CMK in the caller's account
	// and region. You cannot use an alias name or alias ARN in this value.
	//
	// This parameter is optional. If you omit it, ListAliases returns all aliases
	// in the account and region.
	KeyId *string `min:"1" type:"string"`

	// Use this parameter to specify the maximum number of items to return. When
	// this value is present, AWS KMS does not return more than the specified number
	// of items, but it might return fewer.
	//
	// This value is optional. If you include a value, it must be between 1 and
	// 100, inclusive. If you do not include a value, it defaults to 50.
	Limit *int64 `min:"1" type:"integer"`

	// Use this parameter in a subsequent request after you receive a response with
	// truncated results. Set it to the value of NextMarker from the truncated response
	// you just received.
	Marker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListAliasesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAliasesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAliasesInput"}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/ListAliasesResponse
type ListAliasesOutput struct {
	_ struct{} `type:"structure"`

	// A list of aliases.
	Aliases []AliasListEntry `type:"list"`

	// When Truncated is true, this element is present and contains the value to
	// use for the Marker parameter in a subsequent request.
	NextMarker *string `min:"1" type:"string"`

	// A flag that indicates whether there are more items in the list. When this
	// value is true, the list in this response is truncated. To get more items,
	// pass the value of the NextMarker element in thisresponse to the Marker parameter
	// in a subsequent request.
	Truncated *bool `type:"boolean"`
}

// String returns the string representation
func (s ListAliasesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListAliases = "ListAliases"

// ListAliasesRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Gets a list of aliases in the caller's AWS account and region. You cannot
// list aliases in other accounts. For more information about aliases, see CreateAlias.
//
// By default, the ListAliases command returns all aliases in the account and
// region. To get only the aliases that point to a particular customer master
// key (CMK), use the KeyId parameter.
//
// The ListAliases response can include aliases that you created and associated
// with your customer managed CMKs, and aliases that AWS created and associated
// with AWS managed CMKs in your account. You can recognize AWS aliases because
// their names have the format aws/<service-name>, such as aws/dynamodb.
//
// The response might also include aliases that have no TargetKeyId field. These
// are predefined aliases that AWS has created but has not yet associated with
// a CMK. Aliases that AWS creates in your account, including predefined aliases,
// do not count against your AWS KMS aliases limit (https://docs.aws.amazon.com/kms/latest/developerguide/limits.html#aliases-limit).
//
//    // Example sending a request using ListAliasesRequest.
//    req := client.ListAliasesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/ListAliases
func (c *Client) ListAliasesRequest(input *ListAliasesInput) ListAliasesRequest {
	op := &aws.Operation{
		Name:       opListAliases,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"NextMarker"},
			LimitToken:      "Limit",
			TruncationToken: "Truncated",
		},
	}

	if input == nil {
		input = &ListAliasesInput{}
	}

	req := c.newRequest(op, input, &ListAliasesOutput{})
	return ListAliasesRequest{Request: req, Input: input, Copy: c.ListAliasesRequest}
}

// ListAliasesRequest is the request type for the
// ListAliases API operation.
type ListAliasesRequest struct {
	*aws.Request
	Input *ListAliasesInput
	Copy  func(*ListAliasesInput) ListAliasesRequest
}

// Send marshals and sends the ListAliases API request.
func (r ListAliasesRequest) Send(ctx context.Context) (*ListAliasesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAliasesResponse{
		ListAliasesOutput: r.Request.Data.(*ListAliasesOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAliasesRequestPaginator returns a paginator for ListAliases.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAliasesRequest(input)
//   p := kms.NewListAliasesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAliasesPaginator(req ListAliasesRequest) ListAliasesPaginator {
	return ListAliasesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListAliasesInput
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

// ListAliasesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAliasesPaginator struct {
	aws.Pager
}

func (p *ListAliasesPaginator) CurrentPage() *ListAliasesOutput {
	return p.Pager.CurrentPage().(*ListAliasesOutput)
}

// ListAliasesResponse is the response type for the
// ListAliases API operation.
type ListAliasesResponse struct {
	*ListAliasesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAliases request.
func (r *ListAliasesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
