// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Input for SetEndpointAttributes action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/SetEndpointAttributesInput
type SetEndpointAttributesInput struct {
	_ struct{} `type:"structure"`

	// A map of the endpoint attributes. Attributes in this map include the following:
	//
	//    * CustomUserData – arbitrary user data to associate with the endpoint.
	//    Amazon SNS does not use this data. The data must be in UTF-8 format and
	//    less than 2KB.
	//
	//    * Enabled – flag that enables/disables delivery to the endpoint. Amazon
	//    SNS will set this to false when a notification service indicates to Amazon
	//    SNS that the endpoint is invalid. Users can set it back to true, typically
	//    after updating Token.
	//
	//    * Token – device token, also referred to as a registration id, for an
	//    app and mobile device. This is returned from the notification service
	//    when an app and mobile device are registered with the notification service.
	//
	// Attributes is a required field
	Attributes map[string]string `type:"map" required:"true"`

	// EndpointArn used for SetEndpointAttributes action.
	//
	// EndpointArn is a required field
	EndpointArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SetEndpointAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetEndpointAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetEndpointAttributesInput"}

	if s.Attributes == nil {
		invalidParams.Add(aws.NewErrParamRequired("Attributes"))
	}

	if s.EndpointArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/SetEndpointAttributesOutput
type SetEndpointAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetEndpointAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetEndpointAttributes = "SetEndpointAttributes"

// SetEndpointAttributesRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Sets the attributes for an endpoint for a device on one of the supported
// push notification services, such as GCM and APNS. For more information, see
// Using Amazon SNS Mobile Push Notifications (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html).
//
//    // Example sending a request using SetEndpointAttributesRequest.
//    req := client.SetEndpointAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/SetEndpointAttributes
func (c *Client) SetEndpointAttributesRequest(input *SetEndpointAttributesInput) SetEndpointAttributesRequest {
	op := &aws.Operation{
		Name:       opSetEndpointAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetEndpointAttributesInput{}
	}

	req := c.newRequest(op, input, &SetEndpointAttributesOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetEndpointAttributesRequest{Request: req, Input: input, Copy: c.SetEndpointAttributesRequest}
}

// SetEndpointAttributesRequest is the request type for the
// SetEndpointAttributes API operation.
type SetEndpointAttributesRequest struct {
	*aws.Request
	Input *SetEndpointAttributesInput
	Copy  func(*SetEndpointAttributesInput) SetEndpointAttributesRequest
}

// Send marshals and sends the SetEndpointAttributes API request.
func (r SetEndpointAttributesRequest) Send(ctx context.Context) (*SetEndpointAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetEndpointAttributesResponse{
		SetEndpointAttributesOutput: r.Request.Data.(*SetEndpointAttributesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetEndpointAttributesResponse is the response type for the
// SetEndpointAttributes API operation.
type SetEndpointAttributesResponse struct {
	*SetEndpointAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetEndpointAttributes request.
func (r *SetEndpointAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
