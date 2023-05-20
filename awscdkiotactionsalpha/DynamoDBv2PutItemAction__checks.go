//go:build !no_runtime_type_checking

package awscdkiotactionsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
)

func validateNewDynamoDBv2PutItemActionParameters(table awsdynamodb.ITable, props *DynamoDBv2PutItemActionProps) error {
	if table == nil {
		return fmt.Errorf("parameter table is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

