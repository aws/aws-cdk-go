//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (m *jsiiProxy_McpServerTargetConfiguration) validateBindParameters(scope constructs.Construct, gateway IGateway) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if gateway == nil {
		return fmt.Errorf("parameter gateway is required, but nil was provided")
	}

	return nil
}

func validateMcpServerTargetConfiguration_CreateParameters(endpoint *string) error {
	if endpoint == nil {
		return fmt.Errorf("parameter endpoint is required, but nil was provided")
	}

	return nil
}

func validateNewMcpServerTargetConfigurationParameters(endpoint *string) error {
	if endpoint == nil {
		return fmt.Errorf("parameter endpoint is required, but nil was provided")
	}

	return nil
}

