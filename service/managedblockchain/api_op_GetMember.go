// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/GetMemberInput
type GetMemberInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the member.
	//
	// MemberId is a required field
	MemberId *string `location:"uri" locationName:"memberId" min:"1" type:"string" required:"true"`

	// The unique identifier of the network to which the member belongs.
	//
	// NetworkId is a required field
	NetworkId *string `location:"uri" locationName:"networkId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMemberInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMemberInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMemberInput"}

	if s.MemberId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MemberId"))
	}
	if s.MemberId != nil && len(*s.MemberId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MemberId", 1))
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
func (s GetMemberInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MemberId != nil {
		v := *s.MemberId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "memberId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NetworkId != nil {
		v := *s.NetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "networkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/GetMemberOutput
type GetMemberOutput struct {
	_ struct{} `type:"structure"`

	// The properties of a member.
	Member *Member `type:"structure"`
}

// String returns the string representation
func (s GetMemberOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetMemberOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Member != nil {
		v := s.Member

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Member", v, metadata)
	}
	return nil
}

const opGetMember = "GetMember"

// GetMemberRequest returns a request value for making API operation for
// Amazon Managed Blockchain.
//
// Returns detailed information about a member.
//
//    // Example sending a request using GetMemberRequest.
//    req := client.GetMemberRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/GetMember
func (c *Client) GetMemberRequest(input *GetMemberInput) GetMemberRequest {
	op := &aws.Operation{
		Name:       opGetMember,
		HTTPMethod: "GET",
		HTTPPath:   "/networks/{networkId}/members/{memberId}",
	}

	if input == nil {
		input = &GetMemberInput{}
	}

	req := c.newRequest(op, input, &GetMemberOutput{})
	return GetMemberRequest{Request: req, Input: input, Copy: c.GetMemberRequest}
}

// GetMemberRequest is the request type for the
// GetMember API operation.
type GetMemberRequest struct {
	*aws.Request
	Input *GetMemberInput
	Copy  func(*GetMemberInput) GetMemberRequest
}

// Send marshals and sends the GetMember API request.
func (r GetMemberRequest) Send(ctx context.Context) (*GetMemberResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMemberResponse{
		GetMemberOutput: r.Request.Data.(*GetMemberOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMemberResponse is the response type for the
// GetMember API operation.
type GetMemberResponse struct {
	*GetMemberOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMember request.
func (r *GetMemberResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
