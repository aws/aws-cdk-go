//go:build !no_runtime_type_checking

package awsec2alpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IVpcV2) validateAddEgressOnlyInternetGatewayParameters(options *EgressOnlyInternetGatewayOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IVpcV2) validateAddInternetGatewayParameters(options *InternetGatewayOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IVpcV2) validateAddNatGatewayParameters(options *NatGatewayOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IVpcV2) validateCreateAcceptorVpcRoleParameters(requestorAccountId *string) error {
	if requestorAccountId == nil {
		return fmt.Errorf("parameter requestorAccountId is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IVpcV2) validateCreatePeeringConnectionParameters(id *string, options *VPCPeeringConnectionOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IVpcV2) validateEnableVpnGatewayV2Parameters(options *VPNGatewayV2Options) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

