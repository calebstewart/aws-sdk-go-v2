// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackagevod

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/DeletePackagingConfigurationRequest
type DeletePackagingConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Id is a required field
	Id *string `location:"uri" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePackagingConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePackagingConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePackagingConfigurationInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeletePackagingConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/DeletePackagingConfigurationResponse
type DeletePackagingConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePackagingConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeletePackagingConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeletePackagingConfiguration = "DeletePackagingConfiguration"

// DeletePackagingConfigurationRequest returns a request value for making API operation for
// AWS Elemental MediaPackage VOD.
//
// Deletes a MediaPackage VOD PackagingConfiguration resource.
//
//    // Example sending a request using DeletePackagingConfigurationRequest.
//    req := client.DeletePackagingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/DeletePackagingConfiguration
func (c *Client) DeletePackagingConfigurationRequest(input *DeletePackagingConfigurationInput) DeletePackagingConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeletePackagingConfiguration,
		HTTPMethod: "DELETE",
		HTTPPath:   "/packaging_configurations/{id}",
	}

	if input == nil {
		input = &DeletePackagingConfigurationInput{}
	}

	req := c.newRequest(op, input, &DeletePackagingConfigurationOutput{})
	return DeletePackagingConfigurationRequest{Request: req, Input: input, Copy: c.DeletePackagingConfigurationRequest}
}

// DeletePackagingConfigurationRequest is the request type for the
// DeletePackagingConfiguration API operation.
type DeletePackagingConfigurationRequest struct {
	*aws.Request
	Input *DeletePackagingConfigurationInput
	Copy  func(*DeletePackagingConfigurationInput) DeletePackagingConfigurationRequest
}

// Send marshals and sends the DeletePackagingConfiguration API request.
func (r DeletePackagingConfigurationRequest) Send(ctx context.Context) (*DeletePackagingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePackagingConfigurationResponse{
		DeletePackagingConfigurationOutput: r.Request.Data.(*DeletePackagingConfigurationOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePackagingConfigurationResponse is the response type for the
// DeletePackagingConfiguration API operation.
type DeletePackagingConfigurationResponse struct {
	*DeletePackagingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePackagingConfiguration request.
func (r *DeletePackagingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
