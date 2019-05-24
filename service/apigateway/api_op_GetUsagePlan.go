// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The GET request to get a usage plan of a given plan identifier.
type GetUsagePlanInput struct {
	_ struct{} `type:"structure"`

	// [Required] The identifier of the UsagePlan resource to be retrieved.
	//
	// UsagePlanId is a required field
	UsagePlanId *string `location:"uri" locationName:"usageplanId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetUsagePlanInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetUsagePlanInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetUsagePlanInput"}

	if s.UsagePlanId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UsagePlanId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetUsagePlanInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.UsagePlanId != nil {
		v := *s.UsagePlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "usageplanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents a usage plan than can specify who can assess associated API stages
// with specified request limits and quotas.
//
// In a usage plan, you associate an API by specifying the API's Id and a stage
// name of the specified API. You add plan customers by adding API keys to the
// plan.
//
// Create and Use Usage Plans (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html)
type GetUsagePlanOutput struct {
	_ struct{} `type:"structure"`

	// The associated API stages of a usage plan.
	ApiStages []ApiStage `locationName:"apiStages" type:"list"`

	// The description of a usage plan.
	Description *string `locationName:"description" type:"string"`

	// The identifier of a UsagePlan resource.
	Id *string `locationName:"id" type:"string"`

	// The name of a usage plan.
	Name *string `locationName:"name" type:"string"`

	// The AWS Markeplace product identifier to associate with the usage plan as
	// a SaaS product on AWS Marketplace.
	ProductCode *string `locationName:"productCode" type:"string"`

	// The maximum number of permitted requests per a given unit time interval.
	Quota *QuotaSettings `locationName:"quota" type:"structure"`

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string `locationName:"tags" type:"map"`

	// The request throttle limits of a usage plan.
	Throttle *ThrottleSettings `locationName:"throttle" type:"structure"`
}

// String returns the string representation
func (s GetUsagePlanOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetUsagePlanOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.ApiStages) > 0 {
		v := s.ApiStages

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "apiStages", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProductCode != nil {
		v := *s.ProductCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "productCode", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Quota != nil {
		v := s.Quota

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "quota", v, metadata)
	}
	if len(s.Tags) > 0 {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Throttle != nil {
		v := s.Throttle

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "throttle", v, metadata)
	}
	return nil
}

const opGetUsagePlan = "GetUsagePlan"

// GetUsagePlanRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Gets a usage plan of a given plan identifier.
//
//    // Example sending a request using GetUsagePlanRequest.
//    req := client.GetUsagePlanRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetUsagePlanRequest(input *GetUsagePlanInput) GetUsagePlanRequest {
	op := &aws.Operation{
		Name:       opGetUsagePlan,
		HTTPMethod: "GET",
		HTTPPath:   "/usageplans/{usageplanId}",
	}

	if input == nil {
		input = &GetUsagePlanInput{}
	}

	req := c.newRequest(op, input, &GetUsagePlanOutput{})
	return GetUsagePlanRequest{Request: req, Input: input, Copy: c.GetUsagePlanRequest}
}

// GetUsagePlanRequest is the request type for the
// GetUsagePlan API operation.
type GetUsagePlanRequest struct {
	*aws.Request
	Input *GetUsagePlanInput
	Copy  func(*GetUsagePlanInput) GetUsagePlanRequest
}

// Send marshals and sends the GetUsagePlan API request.
func (r GetUsagePlanRequest) Send(ctx context.Context) (*GetUsagePlanResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetUsagePlanResponse{
		GetUsagePlanOutput: r.Request.Data.(*GetUsagePlanOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetUsagePlanResponse is the response type for the
// GetUsagePlan API operation.
type GetUsagePlanResponse struct {
	*GetUsagePlanOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetUsagePlan request.
func (r *GetUsagePlanResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
