// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/UpdateConstraintInput
type UpdateConstraintInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The updated description of the constraint.
	Description *string `type:"string"`

	// The identifier of the constraint.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// The constraint parameters, in JSON format. The syntax depends on the constraint
	// type as follows:
	//
	// LAUNCH
	//
	// Specify the RoleArn property as follows:
	//
	// {"RoleArn" : "arn:aws:iam::123456789012:role/LaunchRole"}
	//
	// You cannot have both a LAUNCH and a STACKSET constraint.
	//
	// You also cannot have more than one LAUNCH constraint on a product and portfolio.
	//
	// NOTIFICATION
	//
	// Specify the NotificationArns property as follows:
	//
	// {"NotificationArns" : ["arn:aws:sns:us-east-1:123456789012:Topic"]}
	//
	// RESOURCE_UPDATE
	//
	// Specify the TagUpdatesOnProvisionedProduct property as follows:
	//
	// {"Version":"2.0","Properties":{"TagUpdateOnProvisionedProduct":"String"}}
	//
	// The TagUpdatesOnProvisionedProduct property accepts a string value of ALLOWED
	// or NOT_ALLOWED.
	//
	// STACKSET
	//
	// Specify the Parameters property as follows:
	//
	// {"Version": "String", "Properties": {"AccountList": [ "String" ], "RegionList":
	// [ "String" ], "AdminRole": "String", "ExecutionRole": "String"}}
	//
	// You cannot have both a LAUNCH and a STACKSET constraint.
	//
	// You also cannot have more than one STACKSET constraint on a product and portfolio.
	//
	// Products with a STACKSET constraint will launch an AWS CloudFormation stack
	// set.
	//
	// TEMPLATE
	//
	// Specify the Rules property. For more information, see Template Constraint
	// Rules (http://docs.aws.amazon.com/servicecatalog/latest/adminguide/reference-template_constraint_rules.html).
	Parameters *string `type:"string"`
}

// String returns the string representation
func (s UpdateConstraintInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateConstraintInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateConstraintInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/UpdateConstraintOutput
type UpdateConstraintOutput struct {
	_ struct{} `type:"structure"`

	// Information about the constraint.
	ConstraintDetail *ConstraintDetail `type:"structure"`

	// The constraint parameters.
	ConstraintParameters *string `type:"string"`

	// The status of the current request.
	Status Status `type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateConstraintOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateConstraint = "UpdateConstraint"

// UpdateConstraintRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Updates the specified constraint.
//
//    // Example sending a request using UpdateConstraintRequest.
//    req := client.UpdateConstraintRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/UpdateConstraint
func (c *Client) UpdateConstraintRequest(input *UpdateConstraintInput) UpdateConstraintRequest {
	op := &aws.Operation{
		Name:       opUpdateConstraint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateConstraintInput{}
	}

	req := c.newRequest(op, input, &UpdateConstraintOutput{})
	return UpdateConstraintRequest{Request: req, Input: input, Copy: c.UpdateConstraintRequest}
}

// UpdateConstraintRequest is the request type for the
// UpdateConstraint API operation.
type UpdateConstraintRequest struct {
	*aws.Request
	Input *UpdateConstraintInput
	Copy  func(*UpdateConstraintInput) UpdateConstraintRequest
}

// Send marshals and sends the UpdateConstraint API request.
func (r UpdateConstraintRequest) Send(ctx context.Context) (*UpdateConstraintResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateConstraintResponse{
		UpdateConstraintOutput: r.Request.Data.(*UpdateConstraintOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateConstraintResponse is the response type for the
// UpdateConstraint API operation.
type UpdateConstraintResponse struct {
	*UpdateConstraintOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateConstraint request.
func (r *UpdateConstraintResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
