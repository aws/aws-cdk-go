//go:build !no_runtime_type_checking

package awscdkpipessourcesalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

func (d *jsiiProxy_DynamoDBSource) validateBindParameters(_pipe awscdkpipesalpha.IPipe) error {
	if _pipe == nil {
		return fmt.Errorf("parameter _pipe is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DynamoDBSource) validateGetDeadLetterTargetArnParameters(deadLetterTarget interface{}) error {
	switch deadLetterTarget.(type) {
	case awssqs.IQueue:
		// ok
	case awssns.ITopic:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(deadLetterTarget) {
			return fmt.Errorf("parameter deadLetterTarget must be one of the allowed types: awssqs.IQueue, awssns.ITopic; received %#v (a %T)", deadLetterTarget, deadLetterTarget)
		}
	}

	return nil
}

func (d *jsiiProxy_DynamoDBSource) validateGrantPushParameters(grantee awsiam.IRole, deadLetterTarget interface{}) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	switch deadLetterTarget.(type) {
	case awssqs.IQueue:
		// ok
	case awssns.ITopic:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(deadLetterTarget) {
			return fmt.Errorf("parameter deadLetterTarget must be one of the allowed types: awssqs.IQueue, awssns.ITopic; received %#v (a %T)", deadLetterTarget, deadLetterTarget)
		}
	}

	return nil
}

func (d *jsiiProxy_DynamoDBSource) validateGrantReadParameters(grantee awsiam.IRole) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateDynamoDBSource_IsSourceWithDeadLetterTargetParameters(source awscdkpipesalpha.ISource) error {
	if source == nil {
		return fmt.Errorf("parameter source is required, but nil was provided")
	}

	return nil
}

func validateNewDynamoDBSourceParameters(table awsdynamodb.ITableV2, parameters *DynamoDBSourceParameters) error {
	if table == nil {
		return fmt.Errorf("parameter table is required, but nil was provided")
	}

	if parameters == nil {
		return fmt.Errorf("parameter parameters is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(parameters, func() string { return "parameter parameters" }); err != nil {
		return err
	}

	return nil
}

