// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/UpdateChannelClassRequest
type UpdateChannelClassInput struct {
	_ struct{} `type:"structure"`

	// A standard channel has two encoding pipelines and a single pipeline channel
	// only has one.
	//
	// ChannelClass is a required field
	ChannelClass ChannelClass `locationName:"channelClass" type:"string" required:"true" enum:"true"`

	// ChannelId is a required field
	ChannelId *string `location:"uri" locationName:"channelId" type:"string" required:"true"`

	Destinations []OutputDestination `locationName:"destinations" type:"list"`
}

// String returns the string representation
func (s UpdateChannelClassInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateChannelClassInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateChannelClassInput"}
	if len(s.ChannelClass) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ChannelClass"))
	}

	if s.ChannelId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChannelId"))
	}
	if s.Destinations != nil {
		for i, v := range s.Destinations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Destinations", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateChannelClassInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if len(s.ChannelClass) > 0 {
		v := s.ChannelClass

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "channelClass", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.Destinations) > 0 {
		v := s.Destinations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "destinations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.ChannelId != nil {
		v := *s.ChannelId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "channelId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/UpdateChannelClassResponse
type UpdateChannelClassOutput struct {
	_ struct{} `type:"structure"`

	Channel *Channel `locationName:"channel" type:"structure"`
}

// String returns the string representation
func (s UpdateChannelClassOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateChannelClassOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Channel != nil {
		v := s.Channel

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "channel", v, metadata)
	}
	return nil
}

const opUpdateChannelClass = "UpdateChannelClass"

// UpdateChannelClassRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Changes the class of the channel.
//
//    // Example sending a request using UpdateChannelClassRequest.
//    req := client.UpdateChannelClassRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/UpdateChannelClass
func (c *Client) UpdateChannelClassRequest(input *UpdateChannelClassInput) UpdateChannelClassRequest {
	op := &aws.Operation{
		Name:       opUpdateChannelClass,
		HTTPMethod: "PUT",
		HTTPPath:   "/prod/channels/{channelId}/channelClass",
	}

	if input == nil {
		input = &UpdateChannelClassInput{}
	}

	req := c.newRequest(op, input, &UpdateChannelClassOutput{})
	return UpdateChannelClassRequest{Request: req, Input: input, Copy: c.UpdateChannelClassRequest}
}

// UpdateChannelClassRequest is the request type for the
// UpdateChannelClass API operation.
type UpdateChannelClassRequest struct {
	*aws.Request
	Input *UpdateChannelClassInput
	Copy  func(*UpdateChannelClassInput) UpdateChannelClassRequest
}

// Send marshals and sends the UpdateChannelClass API request.
func (r UpdateChannelClassRequest) Send(ctx context.Context) (*UpdateChannelClassResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateChannelClassResponse{
		UpdateChannelClassOutput: r.Request.Data.(*UpdateChannelClassOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateChannelClassResponse is the response type for the
// UpdateChannelClass API operation.
type UpdateChannelClassResponse struct {
	*UpdateChannelClassOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateChannelClass request.
func (r *UpdateChannelClassResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
