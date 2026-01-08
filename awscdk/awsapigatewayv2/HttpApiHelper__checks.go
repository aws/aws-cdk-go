//go:build !no_runtime_type_checking

package awsapigatewayv2

import (
	"fmt"
)

func validateHttpApiHelper_FromHttpApiParameters(httpApi IHttpApiRef) error {
	if httpApi == nil {
		return fmt.Errorf("parameter httpApi is required, but nil was provided")
	}

	return nil
}

