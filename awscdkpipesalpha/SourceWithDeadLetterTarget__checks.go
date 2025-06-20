//go:build !no_runtime_type_checking

package awscdkpipesalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

func (s *jsiiProxy_SourceWithDeadLetterTarget) validateBindParameters(pipe IPipe) error {
	if pipe == nil {
		return fmt.Errorf("parameter pipe is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SourceWithDeadLetterTarget) validateGetDeadLetterTargetArnParameters(deadLetterTarget interface{}) error {
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

func (s *jsiiProxy_SourceWithDeadLetterTarget) validateGrantPushParameters(grantee awsiam.IRole, deadLetterTarget interface{}) error {
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

func (s *jsiiProxy_SourceWithDeadLetterTarget) validateGrantReadParameters(grantee awsiam.IRole) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateSourceWithDeadLetterTarget_IsSourceWithDeadLetterTargetParameters(source ISource) error {
	if source == nil {
		return fmt.Errorf("parameter source is required, but nil was provided")
	}

	return nil
}

func validateNewSourceWithDeadLetterTargetParameters(sourceArn *string, deadLetterTarget interface{}) error {
	if sourceArn == nil {
		return fmt.Errorf("parameter sourceArn is required, but nil was provided")
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

