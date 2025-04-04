//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IIpamPool) validateProvisionCidrParameters(id *string, options *IpamPoolCidrProvisioningOptions) error {
	return nil
}

