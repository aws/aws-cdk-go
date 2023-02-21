//go:build !no_runtime_type_checking

package awsappmesh

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (h *jsiiProxy_HttpGatewayRoutePathMatch) validateBindParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateHttpGatewayRoutePathMatch_ExactlyParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateHttpGatewayRoutePathMatch_RegexParameters(regex *string) error {
	if regex == nil {
		return fmt.Errorf("parameter regex is required, but nil was provided")
	}

	return nil
}

func validateHttpGatewayRoutePathMatch_StartsWithParameters(prefix *string) error {
	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

