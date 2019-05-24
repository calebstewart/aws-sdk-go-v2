// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackagevod

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/DeleteAssetRequest
type DeleteAssetInput struct {
	_ struct{} `type:"structure"`

	// Id is a required field
	Id *string `location:"uri" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAssetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAssetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAssetInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAssetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/DeleteAssetResponse
type DeleteAssetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAssetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAssetOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteAsset = "DeleteAsset"

// DeleteAssetRequest returns a request value for making API operation for
// AWS Elemental MediaPackage VOD.
//
// Deletes an existing MediaPackage VOD Asset resource.
//
//    // Example sending a request using DeleteAssetRequest.
//    req := client.DeleteAssetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/DeleteAsset
func (c *Client) DeleteAssetRequest(input *DeleteAssetInput) DeleteAssetRequest {
	op := &aws.Operation{
		Name:       opDeleteAsset,
		HTTPMethod: "DELETE",
		HTTPPath:   "/assets/{id}",
	}

	if input == nil {
		input = &DeleteAssetInput{}
	}

	req := c.newRequest(op, input, &DeleteAssetOutput{})
	return DeleteAssetRequest{Request: req, Input: input, Copy: c.DeleteAssetRequest}
}

// DeleteAssetRequest is the request type for the
// DeleteAsset API operation.
type DeleteAssetRequest struct {
	*aws.Request
	Input *DeleteAssetInput
	Copy  func(*DeleteAssetInput) DeleteAssetRequest
}

// Send marshals and sends the DeleteAsset API request.
func (r DeleteAssetRequest) Send(ctx context.Context) (*DeleteAssetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAssetResponse{
		DeleteAssetOutput: r.Request.Data.(*DeleteAssetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAssetResponse is the response type for the
// DeleteAsset API operation.
type DeleteAssetResponse struct {
	*DeleteAssetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAsset request.
func (r *DeleteAssetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
