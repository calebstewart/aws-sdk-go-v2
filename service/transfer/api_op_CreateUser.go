// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transfer

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/CreateUserRequest
type CreateUserInput struct {
	_ struct{} `type:"structure"`

	// The landing directory (folder) for a user when they log in to the server
	// using their SFTP client. An example is /home/username .
	HomeDirectory *string `type:"string"`

	// A scope-down policy for your user so you can use the same IAM role across
	// multiple users. This policy scopes down user access to portions of their
	// Amazon S3 bucket. Variables you can use inside this policy include ${Transfer:UserName},
	// ${Transfer:HomeDirectory}, and ${Transfer:HomeBucket}.
	Policy *string `type:"string"`

	// The IAM role that controls your user's access to your Amazon S3 bucket. The
	// policies attached to this role will determine the level of access you want
	// to provide your users when transferring files into and out of your Amazon
	// S3 bucket or buckets. The IAM role should also contain a trust relationship
	// that allows the SFTP server to access your resources when servicing your
	// SFTP user's transfer requests.
	//
	// Role is a required field
	Role *string `type:"string" required:"true"`

	// A system-assigned unique identifier for an SFTP server instance. This is
	// the specific SFTP server that you added your user to.
	//
	// ServerId is a required field
	ServerId *string `type:"string" required:"true"`

	// The public portion of the Secure Shall (SSH) key used to authenticate the
	// user to the SFTP server.
	SshPublicKeyBody *string `type:"string"`

	// Key-value pairs that can be used to group and search for users. Tags are
	// metadata attached to users for any purpose.
	Tags []Tag `min:"1" type:"list"`

	// A unique string that identifies a user and is associated with a server as
	// specified by the ServerId. This user name must be a minimum of 3 and a maximum
	// of 32 characters long. The following are valid characters: a-z, A-Z, 0-9,
	// underscore, and hyphen. The user name can't start with a hyphen.
	//
	// UserName is a required field
	UserName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateUserInput"}

	if s.Role == nil {
		invalidParams.Add(aws.NewErrParamRequired("Role"))
	}

	if s.ServerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerId"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/CreateUserResponse
type CreateUserOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the SFTP server that the user is attached to.
	//
	// ServerId is a required field
	ServerId *string `type:"string" required:"true"`

	// A unique string that identifies a user account associated with an SFTP server.
	//
	// UserName is a required field
	UserName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateUser = "CreateUser"

// CreateUserRequest returns a request value for making API operation for
// AWS Transfer for SFTP.
//
// Adds a user and associate them with an existing Secure File Transfer Protocol
// (SFTP) server. Using parameters for CreateUser, you can specify the user
// name, set the home directory, store the user's public key, and assign the
// user's AWS Identity and Access Management (IAM) role. You can also optionally
// add a scope-down policy, and assign metadata with tags that can be used to
// group and search for users.
//
// The response returns the UserName and ServerId values of the new user for
// that server.
//
//    // Example sending a request using CreateUserRequest.
//    req := client.CreateUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/CreateUser
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
