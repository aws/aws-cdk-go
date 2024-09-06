//go:build !no_runtime_type_checking

package awsec2alpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (v *jsiiProxy_VpcV2) validateAddClientVpnEndpointParameters(id *string, options *awsec2.ClientVpnEndpointOptions) error {
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

func (v *jsiiProxy_VpcV2) validateAddFlowLogParameters(id *string, options *awsec2.FlowLogOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (v *jsiiProxy_VpcV2) validateAddGatewayEndpointParameters(id *string, options *awsec2.GatewayVpcEndpointOptions) error {
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

func (v *jsiiProxy_VpcV2) validateAddInterfaceEndpointParameters(id *string, options *awsec2.InterfaceVpcEndpointOptions) error {
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

func (v *jsiiProxy_VpcV2) validateAddVpnConnectionParameters(id *string, options *awsec2.VpnConnectionOptions) error {
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

func (v *jsiiProxy_VpcV2) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (v *jsiiProxy_VpcV2) validateEnableVpnGatewayParameters(options *awsec2.EnableVpnGatewayOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (v *jsiiProxy_VpcV2) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	if arnAttr == nil {
		return fmt.Errorf("parameter arnAttr is required, but nil was provided")
	}

	if arnComponents == nil {
		return fmt.Errorf("parameter arnComponents is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(arnComponents, func() string { return "parameter arnComponents" }); err != nil {
		return err
	}

	return nil
}

func (v *jsiiProxy_VpcV2) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (v *jsiiProxy_VpcV2) validateSelectSubnetObjectsParameters(selection *awsec2.SubnetSelection) error {
	if err := _jsii_.ValidateStruct(selection, func() string { return "parameter selection" }); err != nil {
		return err
	}

	return nil
}

func (v *jsiiProxy_VpcV2) validateSelectSubnetsParameters(selection *awsec2.SubnetSelection) error {
	if err := _jsii_.ValidateStruct(selection, func() string { return "parameter selection" }); err != nil {
		return err
	}

	return nil
}

func validateVpcV2_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateVpcV2_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateVpcV2_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_VpcV2) validateSetIncompleteSubnetDefinitionParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewVpcV2Parameters(scope constructs.Construct, id *string, props *VpcV2Props) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

