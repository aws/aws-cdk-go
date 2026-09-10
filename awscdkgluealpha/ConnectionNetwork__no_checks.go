//go:build no_runtime_type_checking

package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateConnectionNetwork_SubnetParameters(subnet awsec2.ISubnet) error {
	return nil
}

func validateConnectionNetwork_VpcParameters(vpc awsec2.IVpc, vpcSubnets *awsec2.SubnetSelection) error {
	return nil
}

