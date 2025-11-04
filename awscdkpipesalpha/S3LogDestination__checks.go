//go:build !no_runtime_type_checking

package awscdkpipesalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

func (s *jsiiProxy_S3LogDestination) validateBindParameters(pipe IPipe) error {
	if pipe == nil {
		return fmt.Errorf("parameter pipe is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_S3LogDestination) validateGrantPushParameters(grantee awsiam.IRole) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateNewS3LogDestinationParameters(parameters *S3LogDestinationProps) error {
	if parameters == nil {
		return fmt.Errorf("parameter parameters is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(parameters, func() string { return "parameter parameters" }); err != nil {
		return err
	}

	return nil
}

