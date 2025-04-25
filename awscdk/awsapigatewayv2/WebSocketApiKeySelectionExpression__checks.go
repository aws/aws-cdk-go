//go:build !no_runtime_type_checking

package awsapigatewayv2

import (
	"fmt"
)

func validateNewWebSocketApiKeySelectionExpressionParameters(customApiKeySelector *string) error {
	if customApiKeySelector == nil {
		return fmt.Errorf("parameter customApiKeySelector is required, but nil was provided")
	}

	return nil
}

