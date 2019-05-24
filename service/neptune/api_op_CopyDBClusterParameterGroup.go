// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/CopyDBClusterParameterGroupMessage
type CopyDBClusterParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The identifier or Amazon Resource Name (ARN) for the source DB cluster parameter
	// group. For information about creating an ARN, see Constructing an Amazon
	// Resource Name (ARN) (https://docs.aws.amazon.com/neptune/latest/UserGuide/tagging.ARN.html#tagging.ARN.Constructing).
	//
	// Constraints:
	//
	//    * Must specify a valid DB cluster parameter group.
	//
	//    * If the source DB cluster parameter group is in the same AWS Region as
	//    the copy, specify a valid DB parameter group identifier, for example my-db-cluster-param-group,
	//    or a valid ARN.
	//
	//    * If the source DB parameter group is in a different AWS Region than the
	//    copy, specify a valid DB cluster parameter group ARN, for example arn:aws:rds:us-east-1:123456789012:cluster-pg:custom-cluster-group1.
	//
	// SourceDBClusterParameterGroupIdentifier is a required field
	SourceDBClusterParameterGroupIdentifier *string `type:"string" required:"true"`

	// The tags to be assigned to the copied DB cluster parameter group.
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// A description for the copied DB cluster parameter group.
	//
	// TargetDBClusterParameterGroupDescription is a required field
	TargetDBClusterParameterGroupDescription *string `type:"string" required:"true"`

	// The identifier for the copied DB cluster parameter group.
	//
	// Constraints:
	//
	//    * Cannot be null, empty, or blank
	//
	//    * Must contain from 1 to 255 letters, numbers, or hyphens
	//
	//    * First character must be a letter
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens
	//
	// Example: my-cluster-param-group1
	//
	// TargetDBClusterParameterGroupIdentifier is a required field
	TargetDBClusterParameterGroupIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CopyDBClusterParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CopyDBClusterParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CopyDBClusterParameterGroupInput"}

	if s.SourceDBClusterParameterGroupIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceDBClusterParameterGroupIdentifier"))
	}

	if s.TargetDBClusterParameterGroupDescription == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetDBClusterParameterGroupDescription"))
	}

	if s.TargetDBClusterParameterGroupIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetDBClusterParameterGroupIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/CopyDBClusterParameterGroupResult
type CopyDBClusterParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon Neptune DB cluster parameter group.
	//
	// This data type is used as a response element in the DescribeDBClusterParameterGroups
	// action.
	DBClusterParameterGroup *DBClusterParameterGroup `type:"structure"`
}

// String returns the string representation
func (s CopyDBClusterParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCopyDBClusterParameterGroup = "CopyDBClusterParameterGroup"

// CopyDBClusterParameterGroupRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Copies the specified DB cluster parameter group.
//
//    // Example sending a request using CopyDBClusterParameterGroupRequest.
//    req := client.CopyDBClusterParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/CopyDBClusterParameterGroup
func (c *Client) CopyDBClusterParameterGroupRequest(input *CopyDBClusterParameterGroupInput) CopyDBClusterParameterGroupRequest {
	op := &aws.Operation{
		Name:       opCopyDBClusterParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CopyDBClusterParameterGroupInput{}
	}

	req := c.newRequest(op, input, &CopyDBClusterParameterGroupOutput{})
	return CopyDBClusterParameterGroupRequest{Request: req, Input: input, Copy: c.CopyDBClusterParameterGroupRequest}
}

// CopyDBClusterParameterGroupRequest is the request type for the
// CopyDBClusterParameterGroup API operation.
type CopyDBClusterParameterGroupRequest struct {
	*aws.Request
	Input *CopyDBClusterParameterGroupInput
	Copy  func(*CopyDBClusterParameterGroupInput) CopyDBClusterParameterGroupRequest
}

// Send marshals and sends the CopyDBClusterParameterGroup API request.
func (r CopyDBClusterParameterGroupRequest) Send(ctx context.Context) (*CopyDBClusterParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CopyDBClusterParameterGroupResponse{
		CopyDBClusterParameterGroupOutput: r.Request.Data.(*CopyDBClusterParameterGroupOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CopyDBClusterParameterGroupResponse is the response type for the
// CopyDBClusterParameterGroup API operation.
type CopyDBClusterParameterGroupResponse struct {
	*CopyDBClusterParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CopyDBClusterParameterGroup request.
func (r *CopyDBClusterParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
