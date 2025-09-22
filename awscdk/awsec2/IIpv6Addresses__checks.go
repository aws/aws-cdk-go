//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IIpv6Addresses) validateAllocateSubnetsIpv6CidrParameters(input *AllocateIpv6CidrRequest) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(input, func() string { return "parameter input" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IIpv6Addresses) validateAllocateVpcIpv6CidrParameters(input *AllocateVpcIpv6CidrRequest) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(input, func() string { return "parameter input" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IIpv6Addresses) validateCreateIpv6CidrBlocksParameters(input *CreateIpv6CidrBlocksRequest) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(input, func() string { return "parameter input" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_IIpv6Addresses) validateSetAmazonProvidedParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

