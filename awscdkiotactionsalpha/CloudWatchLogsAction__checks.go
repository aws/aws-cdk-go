//go:build !no_runtime_type_checking

// Receipt rule actions for AWS IoT
package awscdkiotactionsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

func validateNewCloudWatchLogsActionParameters(logGroup awslogs.ILogGroup, props *CloudWatchLogsActionProps) error {
	if logGroup == nil {
		return fmt.Errorf("parameter logGroup is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

