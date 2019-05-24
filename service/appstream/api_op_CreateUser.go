// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/CreateUserRequest
type CreateUserInput struct {
	_ struct{} `type:"structure"`

	// The authentication type for the user. You must specify USERPOOL.
	//
	// AuthenticationType is a required field
	AuthenticationType AuthenticationType `type:"string" required:"true" enum:"true"`

	// The first name, or given name, of the user.
	FirstName *string `type:"string"`

	// The last name, or surname, of the user.
	LastName *string `type:"string"`

	// The action to take for the welcome email that is sent to a user after the
	// user is created in the user pool. If you specify SUPPRESS, no email is sent.
	// If you specify RESEND, do not specify the first name or last name of the
	// user. If the value is null, the email is sent.
	//
	// The temporary password in the welcome email is valid for only 7 days. If
	// users don’t set their passwords within 7 days, you must send them a new
	// welcome email.
	MessageAction MessageAction `type:"string" enum:"true"`

	// The email address of the user.
	//
	// Users' email addresses are case-sensitive. During login, if they specify
	// an email address that doesn't use the same capitalization as the email address
	// specified when their user pool account was created, a "user does not exist"
	// error message displays.
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateUserInput"}
	if len(s.AuthenticationType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("AuthenticationType"))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/CreateUserResult
type CreateUserOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateUser = "CreateUser"

// CreateUserRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Creates a new user in the user pool.
//
//    // Example sending a request using CreateUserRequest.
//    req := client.CreateUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/CreateUser
func (c *Client) CreateUserRequest(input *CreateUserInput) CreateUserRequest {
	op := &aws.Operation{
		Name:       opCreateUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateUserInput{}
	}

	req := c.newRequest(op, input, &CreateUserOutput{})
	return CreateUserRequest{Request: req, Input: input, Copy: c.CreateUserRequest}
}

// CreateUserRequest is the request type for the
// CreateUser API operation.
type CreateUserRequest struct {
	*aws.Request
	Input *CreateUserInput
	Copy  func(*CreateUserInput) CreateUserRequest
}

// Send marshals and sends the CreateUser API request.
func (r CreateUserRequest) Send(ctx context.Context) (*CreateUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateUserResponse{
		CreateUserOutput: r.Request.Data.(*CreateUserOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateUserResponse is the response type for the
// CreateUser API operation.
type CreateUserResponse struct {
	*CreateUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateUser request.
func (r *CreateUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
