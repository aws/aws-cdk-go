//go:build !no_runtime_type_checking

package awsapigateway

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker"
)

func (s *jsiiProxy_SagemakerIntegration) validateBindParameters(method Method) error {
	if method == nil {
		return fmt.Errorf("parameter method is required, but nil was provided")
	}

	return nil
}

func validateNewSagemakerIntegrationParameters(endpoint awssagemaker.IEndpoint, options *SagemakerIntegrationOptions) error {
	if endpoint == nil {
		return fmt.Errorf("parameter endpoint is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

