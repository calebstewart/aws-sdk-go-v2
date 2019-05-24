// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/EnableKeyRotationRequest
type EnableKeyRotationInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the customer master key (CMK).
	//
	// Specify the key ID or the Amazon Resource Name (ARN) of the CMK.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// KeyId is a required field
	KeyId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s EnableKeyRotationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableKeyRotationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableKeyRotationInput"}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/EnableKeyRotationOutput
type EnableKeyRotationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EnableKeyRotationOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnableKeyRotation = "EnableKeyRotation"

// EnableKeyRotationRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Enables automatic rotation of the key material (https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html)
// for the specified customer master key (CMK). You cannot perform this operation
// on a CMK in a different AWS account.
//
// You cannot enable automatic rotation of CMKs with imported key material or
// CMKs in a custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html).
//
// The result of this operation varies with the key state of the CMK. For details,
// see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using EnableKeyRotationRequest.
//    req := client.EnableKeyRotationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/EnableKeyRotation
func (c *Client) EnableKeyRotationRequest(input *EnableKeyRotationInput) EnableKeyRotationRequest {
	op := &aws.Operation{
		Name:       opEnableKeyRotation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableKeyRotationInput{}
	}

	req := c.newRequest(op, input, &EnableKeyRotationOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return EnableKeyRotationRequest{Request: req, Input: input, Copy: c.EnableKeyRotationRequest}
}

// EnableKeyRotationRequest is the request type for the
// EnableKeyRotation API operation.
type EnableKeyRotationRequest struct {
	*aws.Request
	Input *EnableKeyRotationInput
	Copy  func(*EnableKeyRotationInput) EnableKeyRotationRequest
}

// Send marshals and sends the EnableKeyRotation API request.
func (r EnableKeyRotationRequest) Send(ctx context.Context) (*EnableKeyRotationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableKeyRotationResponse{
		EnableKeyRotationOutput: r.Request.Data.(*EnableKeyRotationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableKeyRotationResponse is the response type for the
// EnableKeyRotation API operation.
type EnableKeyRotationResponse struct {
	*EnableKeyRotationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableKeyRotation request.
func (r *EnableKeyRotationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
