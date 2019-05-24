// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyOptionGroupMessage
type ModifyOptionGroupInput struct {
	_ struct{} `type:"structure"`

	// A value that indicates whether to apply the change immediately or during
	// the next maintenance window for each instance associated with the option
	// group.
	ApplyImmediately *bool `type:"boolean"`

	// The name of the option group to be modified.
	//
	// Permanent options, such as the TDE option for Oracle Advanced Security TDE,
	// can't be removed from an option group, and that option group can't be removed
	// from a DB instance once it is associated with a DB instance
	//
	// OptionGroupName is a required field
	OptionGroupName *string `type:"string" required:"true"`

	// Options in this list are added to the option group or, if already present,
	// the specified configuration is used to update the existing configuration.
	OptionsToInclude []OptionConfiguration `locationNameList:"OptionConfiguration" type:"list"`

	// Options in this list are removed from the option group.
	OptionsToRemove []string `type:"list"`
}

// String returns the string representation
func (s ModifyOptionGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyOptionGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyOptionGroupInput"}

	if s.OptionGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("OptionGroupName"))
	}
	if s.OptionsToInclude != nil {
		for i, v := range s.OptionsToInclude {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "OptionsToInclude", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyOptionGroupResult
type ModifyOptionGroupOutput struct {
	_ struct{} `type:"structure"`

	OptionGroup *OptionGroup `type:"structure"`
}

// String returns the string representation
func (s ModifyOptionGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyOptionGroup = "ModifyOptionGroup"

// ModifyOptionGroupRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Modifies an existing option group.
//
//    // Example sending a request using ModifyOptionGroupRequest.
//    req := client.ModifyOptionGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyOptionGroup
func (c *Client) ModifyOptionGroupRequest(input *ModifyOptionGroupInput) ModifyOptionGroupRequest {
	op := &aws.Operation{
		Name:       opModifyOptionGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyOptionGroupInput{}
	}

	req := c.newRequest(op, input, &ModifyOptionGroupOutput{})
	return ModifyOptionGroupRequest{Request: req, Input: input, Copy: c.ModifyOptionGroupRequest}
}

// ModifyOptionGroupRequest is the request type for the
// ModifyOptionGroup API operation.
type ModifyOptionGroupRequest struct {
	*aws.Request
	Input *ModifyOptionGroupInput
	Copy  func(*ModifyOptionGroupInput) ModifyOptionGroupRequest
}

// Send marshals and sends the ModifyOptionGroup API request.
func (r ModifyOptionGroupRequest) Send(ctx context.Context) (*ModifyOptionGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyOptionGroupResponse{
		ModifyOptionGroupOutput: r.Request.Data.(*ModifyOptionGroupOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyOptionGroupResponse is the response type for the
// ModifyOptionGroup API operation.
type ModifyOptionGroupResponse struct {
	*ModifyOptionGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyOptionGroup request.
func (r *ModifyOptionGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
