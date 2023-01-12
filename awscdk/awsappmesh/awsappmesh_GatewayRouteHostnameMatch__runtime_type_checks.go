//go:build !no_runtime_type_checking

package awsappmesh

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (g *jsiiProxy_GatewayRouteHostnameMatch) validateBindParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateGatewayRouteHostnameMatch_EndsWithParameters(suffix *string) error {
	if suffix == nil {
		return fmt.Errorf("parameter suffix is required, but nil was provided")
	}

	return nil
}

func validateGatewayRouteHostnameMatch_ExactlyParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

