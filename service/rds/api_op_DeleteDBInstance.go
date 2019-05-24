// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DeleteDBInstanceMessage
type DeleteDBInstanceInput struct {
	_ struct{} `type:"structure"`

	// The DB instance identifier for the DB instance to be deleted. This parameter
	// isn't case-sensitive.
	//
	// Constraints:
	//
	//    * Must match the name of an existing DB instance.
	//
	// DBInstanceIdentifier is a required field
	DBInstanceIdentifier *string `type:"string" required:"true"`

	// A value that indicates whether to remove automated backups immediately after
	// the DB instance is deleted. This parameter isn't case-sensitive. The default
	// is to remove automated backups immediately after the DB instance is deleted.
	DeleteAutomatedBackups *bool `type:"boolean"`

	// The DBSnapshotIdentifier of the new DBSnapshot created when the SkipFinalSnapshot
	// parameter is disabled.
	//
	// Specifying this parameter and also specifying to skip final DB snapshot creation
	// in SkipFinalShapshot results in an error.
	//
	// Constraints:
	//
	//    * Must be 1 to 255 letters or numbers.
	//
	//    * First character must be a letter.
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens.
	//
	//    * Can't be specified when deleting a Read Replica.
	FinalDBSnapshotIdentifier *string `type:"string"`

	// A value that indicates whether to skip the creation of a final DB snapshot
	// before the DB instance is deleted. If skip is specified, no DB snapshot is
	// created. If skip is not specified, a DB snapshot is created before the DB
	// instance is deleted. By default, skip is not specified, and the DB snapshot
	// is created.
	//
	// Note that when a DB instance is in a failure state and has a status of 'failed',
	// 'incompatible-restore', or 'incompatible-network', it can only be deleted
	// when skip is specified.
	//
	// Specify skip when deleting a Read Replica.
	//
	// The FinalDBSnapshotIdentifier parameter must be specified if skip is not
	// specified.
	SkipFinalSnapshot *bool `type:"boolean"`
}

// String returns the string representation
func (s DeleteDBInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDBInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDBInstanceInput"}

	if s.DBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBInstanceIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DeleteDBInstanceResult
type DeleteDBInstanceOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon RDS DB instance.
	//
	// This data type is used as a response element in the DescribeDBInstances action.
	DBInstance *DBInstance `type:"structure"`
}

// String returns the string representation
func (s DeleteDBInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDBInstance = "DeleteDBInstance"

// DeleteDBInstanceRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// The DeleteDBInstance action deletes a previously provisioned DB instance.
// When you delete a DB instance, all automated backups for that instance are
// deleted and can't be recovered. Manual DB snapshots of the DB instance to
// be deleted by DeleteDBInstance are not deleted.
//
// If you request a final DB snapshot the status of the Amazon RDS DB instance
// is deleting until the DB snapshot is created. The API action DescribeDBInstance
// is used to monitor the status of this operation. The action can't be canceled
// or reverted once submitted.
//
// Note that when a DB instance is in a failure state and has a status of failed,
// incompatible-restore, or incompatible-network, you can only delete it when
// you skip creation of the final snapshot with the SkipFinalSnapshot parameter.
//
// If the specified DB instance is part of an Amazon Aurora DB cluster, you
// can't delete the DB instance if both of the following conditions are true:
//
//    * The DB cluster is a Read Replica of another Amazon Aurora DB cluster.
//
//    * The DB instance is the only instance in the DB cluster.
//
// To delete a DB instance in this case, first call the PromoteReadReplicaDBCluster
// API action to promote the DB cluster so it's no longer a Read Replica. After
// the promotion completes, then call the DeleteDBInstance API action to delete
// the final instance in the DB cluster.
//
//    // Example sending a request using DeleteDBInstanceRequest.
//    req := client.DeleteDBInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DeleteDBInstance
func (c *Client) DeleteDBInstanceRequest(input *DeleteDBInstanceInput) DeleteDBInstanceRequest {
	op := &aws.Operation{
		Name:       opDeleteDBInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDBInstanceInput{}
	}

	req := c.newRequest(op, input, &DeleteDBInstanceOutput{})
	return DeleteDBInstanceRequest{Request: req, Input: input, Copy: c.DeleteDBInstanceRequest}
}

// DeleteDBInstanceRequest is the request type for the
// DeleteDBInstance API operation.
type DeleteDBInstanceRequest struct {
	*aws.Request
	Input *DeleteDBInstanceInput
	Copy  func(*DeleteDBInstanceInput) DeleteDBInstanceRequest
}

// Send marshals and sends the DeleteDBInstance API request.
func (r DeleteDBInstanceRequest) Send(ctx context.Context) (*DeleteDBInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDBInstanceResponse{
		DeleteDBInstanceOutput: r.Request.Data.(*DeleteDBInstanceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDBInstanceResponse is the response type for the
// DeleteDBInstance API operation.
type DeleteDBInstanceResponse struct {
	*DeleteDBInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDBInstance request.
func (r *DeleteDBInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
