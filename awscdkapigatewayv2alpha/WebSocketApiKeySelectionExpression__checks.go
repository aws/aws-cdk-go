//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha

import (
	"fmt"
)

func validateNewWebSocketApiKeySelectionExpressionParameters(customApiKeySelector *string) error {
	if customApiKeySelector == nil {
		return fmt.Errorf("parameter customApiKeySelector is required, but nil was provided")
	}

	return nil
}

