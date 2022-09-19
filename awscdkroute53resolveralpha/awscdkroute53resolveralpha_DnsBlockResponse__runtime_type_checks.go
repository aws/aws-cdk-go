//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// The CDK Construct Library for AWS::Route53Resolver
package awscdkroute53resolveralpha

import (
	"fmt"
)

func validateDnsBlockResponse_OverrideParameters(domain *string) error {
	if domain == nil {
		return fmt.Errorf("parameter domain is required, but nil was provided")
	}

	return nil
}

