// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package docdb

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// WaitUntilDBInstanceAvailable uses the Amazon DocDB API operation
// DescribeDBInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DocDB) WaitUntilDBInstanceAvailable(ctx context.Context, input *DescribeDBInstancesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBInstanceAvailable",
		MaxAttempts: 60,
		Delay:       aws.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "available",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "deleted",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "deleting",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "failed",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "incompatible-restore",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "incompatible-parameters",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDBInstanceDeleted uses the Amazon DocDB API operation
// DescribeDBInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DocDB) WaitUntilDBInstanceDeleted(ctx context.Context, input *DescribeDBInstancesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBInstanceDeleted",
		MaxAttempts: 60,
		Delay:       aws.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "deleted",
			},
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "DBInstanceNotFound",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "creating",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "modifying",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "rebooting",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "resetting-master-credentials",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}
