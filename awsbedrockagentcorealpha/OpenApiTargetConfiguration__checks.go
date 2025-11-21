//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (o *jsiiProxy_OpenApiTargetConfiguration) validateBindParameters(scope constructs.Construct, gateway IGateway) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if gateway == nil {
		return fmt.Errorf("parameter gateway is required, but nil was provided")
	}

	return nil
}

func validateOpenApiTargetConfiguration_CreateParameters(apiSchema ApiSchema) error {
	if apiSchema == nil {
		return fmt.Errorf("parameter apiSchema is required, but nil was provided")
	}

	return nil
}

func validateNewOpenApiTargetConfigurationParameters(apiSchema ApiSchema) error {
	if apiSchema == nil {
		return fmt.Errorf("parameter apiSchema is required, but nil was provided")
	}

	return nil
}

