// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackagevod

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/DescribePackagingConfigurationRequest
type DescribePackagingConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Id is a required field
	Id *string `location:"uri" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribePackagingConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePackagingConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePackagingConfigurationInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribePackagingConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/DescribePackagingConfigurationResponse
type DescribePackagingConfigurationOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `locationName:"arn" type:"string"`

	// A CMAF packaging configuration.
	CmafPackage *CmafPackage `locationName:"cmafPackage" type:"structure"`

	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *DashPackage `locationName:"dashPackage" type:"structure"`

	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *HlsPackage `locationName:"hlsPackage" type:"structure"`

	Id *string `locationName:"id" type:"string"`

	// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
	MssPackage *MssPackage `locationName:"mssPackage" type:"structure"`

	PackagingGroupId *string `locationName:"packagingGroupId" type:"string"`
}

// String returns the string representation
func (s DescribePackagingConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribePackagingConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CmafPackage != nil {
		v := s.CmafPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "cmafPackage", v, metadata)
	}
	if s.DashPackage != nil {
		v := s.DashPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "dashPackage", v, metadata)
	}
	if s.HlsPackage != nil {
		v := s.HlsPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "hlsPackage", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MssPackage != nil {
		v := s.MssPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "mssPackage", v, metadata)
	}
	if s.PackagingGroupId != nil {
		v := *s.PackagingGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "packagingGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribePackagingConfiguration = "DescribePackagingConfiguration"

// DescribePackagingConfigurationRequest returns a request value for making API operation for
// AWS Elemental MediaPackage VOD.
//
// Returns a description of a MediaPackage VOD PackagingConfiguration resource.
//
//    // Example sending a request using DescribePackagingConfigurationRequest.
//    req := client.DescribePackagingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/DescribePackagingConfiguration
func (c *Client) DescribePackagingConfigurationRequest(input *DescribePackagingConfigurationInput) DescribePackagingConfigurationRequest {
	op := &aws.Operation{
		Name:       opDescribePackagingConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/packaging_configurations/{id}",
	}

	if input == nil {
		input = &DescribePackagingConfigurationInput{}
	}

	req := c.newRequest(op, input, &DescribePackagingConfigurationOutput{})
	return DescribePackagingConfigurationRequest{Request: req, Input: input, Copy: c.DescribePackagingConfigurationRequest}
}

// DescribePackagingConfigurationRequest is the request type for the
// DescribePackagingConfiguration API operation.
type DescribePackagingConfigurationRequest struct {
	*aws.Request
	Input *DescribePackagingConfigurationInput
	Copy  func(*DescribePackagingConfigurationInput) DescribePackagingConfigurationRequest
}

// Send marshals and sends the DescribePackagingConfiguration API request.
func (r DescribePackagingConfigurationRequest) Send(ctx context.Context) (*DescribePackagingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePackagingConfigurationResponse{
		DescribePackagingConfigurationOutput: r.Request.Data.(*DescribePackagingConfigurationOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribePackagingConfigurationResponse is the response type for the
// DescribePackagingConfiguration API operation.
type DescribePackagingConfigurationResponse struct {
	*DescribePackagingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePackagingConfiguration request.
func (r *DescribePackagingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
