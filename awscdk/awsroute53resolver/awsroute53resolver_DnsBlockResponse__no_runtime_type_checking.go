//go:build no_runtime_type_checking

package awsroute53resolver

// Building without runtime type checking enabled, so all the below just return nil

func validateDnsBlockResponse_OverrideParameters(domain *string) error {
	return nil
}

