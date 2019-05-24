// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribePatchPropertiesRequest
type DescribePatchPropertiesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`

	// The operating system type for which to list patches.
	//
	// OperatingSystem is a required field
	OperatingSystem OperatingSystem `type:"string" required:"true" enum:"true"`

	// Indicates whether to list patches for the Windows operating system or for
	// Microsoft applications. Not applicable for Linux operating systems.
	PatchSet PatchSet `type:"string" enum:"true"`

	// The patch property for which you want to view patch details.
	//
	// Property is a required field
	Property PatchProperty `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s DescribePatchPropertiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePatchPropertiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePatchPropertiesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if len(s.OperatingSystem) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("OperatingSystem"))
	}
	if len(s.Property) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Property"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribePatchPropertiesResult
type DescribePatchPropertiesOutput struct {
	_ struct{} `type:"structure"`

	// The token for the next set of items to return. (You use this token in the
	// next call.)
	NextToken *string `type:"string"`

	// A list of the properties for patches matching the filter request parameters.
	Properties []map[string]string `type:"list"`
}

// String returns the string representation
func (s DescribePatchPropertiesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribePatchProperties = "DescribePatchProperties"

// DescribePatchPropertiesRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Lists the properties of available patches organized by product, product family,
// classification, severity, and other properties of available patches. You
// can use the reported properties in the filters you specify in requests for
// actions such as CreatePatchBaseline, UpdatePatchBaseline, DescribeAvailablePatches,
// and DescribePatchBaselines.
//
// The following section lists the properties that can be used in filters for
// each major operating system type:
//
// WINDOWS
//
// Valid properties: PRODUCT, PRODUCT_FAMILY, CLASSIFICATION, MSRC_SEVERITY
//
// AMAZON_LINUX
//
// Valid properties: PRODUCT, CLASSIFICATION, SEVERITY
//
// AMAZON_LINUX_2
//
// Valid properties: PRODUCT, CLASSIFICATION, SEVERITY
//
// UBUNTU
//
// Valid properties: PRODUCT, PRIORITY
//
// REDHAT_ENTERPRISE_LINUX
//
// Valid properties: PRODUCT, CLASSIFICATION, SEVERITY
//
// SUSE
//
// Valid properties: PRODUCT, CLASSIFICATION, SEVERITY
//
// CENTOS
//
// Valid properties: PRODUCT, CLASSIFICATION, SEVERITY
//
//    // Example sending a request using DescribePatchPropertiesRequest.
//    req := client.DescribePatchPropertiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribePatchProperties
func (c *Client) DescribePatchPropertiesRequest(input *DescribePatchPropertiesInput) DescribePatchPropertiesRequest {
	op := &aws.Operation{
		Name:       opDescribePatchProperties,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribePatchPropertiesInput{}
	}

	req := c.newRequest(op, input, &DescribePatchPropertiesOutput{})
	return DescribePatchPropertiesRequest{Request: req, Input: input, Copy: c.DescribePatchPropertiesRequest}
}

// DescribePatchPropertiesRequest is the request type for the
// DescribePatchProperties API operation.
type DescribePatchPropertiesRequest struct {
	*aws.Request
	Input *DescribePatchPropertiesInput
	Copy  func(*DescribePatchPropertiesInput) DescribePatchPropertiesRequest
}

// Send marshals and sends the DescribePatchProperties API request.
func (r DescribePatchPropertiesRequest) Send(ctx context.Context) (*DescribePatchPropertiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePatchPropertiesResponse{
		DescribePatchPropertiesOutput: r.Request.Data.(*DescribePatchPropertiesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribePatchPropertiesResponse is the response type for the
// DescribePatchProperties API operation.
type DescribePatchPropertiesResponse struct {
	*DescribePatchPropertiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePatchProperties request.
func (r *DescribePatchPropertiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
