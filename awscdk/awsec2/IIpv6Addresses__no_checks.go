//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IIpv6Addresses) validateAllocateSubnetsIpv6CidrParameters(input *AllocateIpv6CidrRequest) error {
	return nil
}

func (i *jsiiProxy_IIpv6Addresses) validateAllocateVpcIpv6CidrParameters(input *AllocateVpcIpv6CidrRequest) error {
	return nil
}

func (i *jsiiProxy_IIpv6Addresses) validateCreateIpv6CidrBlocksParameters(input *CreateIpv6CidrBlocksRequest) error {
	return nil
}

func (j *jsiiProxy_IIpv6Addresses) validateSetAmazonProvidedParameters(val *bool) error {
	return nil
}

