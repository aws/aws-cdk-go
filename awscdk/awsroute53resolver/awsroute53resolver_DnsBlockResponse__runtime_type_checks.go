//go:build !no_runtime_type_checking

package awsroute53resolver

import (
	"fmt"
)

func validateDnsBlockResponse_OverrideParameters(domain *string) error {
	if domain == nil {
		return fmt.Errorf("parameter domain is required, but nil was provided")
	}

	return nil
}

