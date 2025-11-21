//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_SmithyTargetConfiguration) validateBindParameters(scope constructs.Construct, gateway IGateway) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if gateway == nil {
		return fmt.Errorf("parameter gateway is required, but nil was provided")
	}

	return nil
}

func validateSmithyTargetConfiguration_CreateParameters(smithyModel ApiSchema) error {
	if smithyModel == nil {
		return fmt.Errorf("parameter smithyModel is required, but nil was provided")
	}

	return nil
}

func validateNewSmithyTargetConfigurationParameters(smithyModel ApiSchema) error {
	if smithyModel == nil {
		return fmt.Errorf("parameter smithyModel is required, but nil was provided")
	}

	return nil
}

