//go:build !no_runtime_type_checking

package awscdkpipessourcesalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

func (s *jsiiProxy_SqsSource) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	if pipe == nil {
		return fmt.Errorf("parameter pipe is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SqsSource) validateGrantReadParameters(grantee awsiam.IRole) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateNewSqsSourceParameters(queue awssqs.IQueue, parameters *SqsSourceParameters) error {
	if queue == nil {
		return fmt.Errorf("parameter queue is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(parameters, func() string { return "parameter parameters" }); err != nil {
		return err
	}

	return nil
}

