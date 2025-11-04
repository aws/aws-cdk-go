//go:build !no_runtime_type_checking

package awscdkpipessourcesalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

func (k *jsiiProxy_KinesisSource) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	if pipe == nil {
		return fmt.Errorf("parameter pipe is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisSource) validateGetDeadLetterTargetArnParameters(deadLetterTarget interface{}) error {
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

func (k *jsiiProxy_KinesisSource) validateGrantPushParameters(grantee awsiam.IRole, deadLetterTarget interface{}) error {
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

func (k *jsiiProxy_KinesisSource) validateGrantReadParameters(grantee awsiam.IRole) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateKinesisSource_IsSourceWithDeadLetterTargetParameters(source awscdkpipesalpha.ISource) error {
	if source == nil {
		return fmt.Errorf("parameter source is required, but nil was provided")
	}

	return nil
}

func validateNewKinesisSourceParameters(stream awskinesis.IStream, parameters *KinesisSourceParameters) error {
	if stream == nil {
		return fmt.Errorf("parameter stream is required, but nil was provided")
	}

	if parameters == nil {
		return fmt.Errorf("parameter parameters is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(parameters, func() string { return "parameter parameters" }); err != nil {
		return err
	}

	return nil
}

