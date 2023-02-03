//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func validateIpAddresses_AwsIpamAllocationParameters(props *AwsIpamProps) error {
	return nil
}

func validateIpAddresses_CidrParameters(cidrBlock *string) error {
	return nil
}

