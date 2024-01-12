//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (p *jsiiProxy_PublicSubnet) validateAddDefaultInternetRouteParameters(gatewayId *string, gatewayAttachment constructs.IDependable) error {
	if gatewayId == nil {
		return fmt.Errorf("parameter gatewayId is required, but nil was provided")
	}

	if gatewayAttachment == nil {
		return fmt.Errorf("parameter gatewayAttachment is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PublicSubnet) validateAddDefaultNatRouteParameters(natGatewayId *string) error {
	if natGatewayId == nil {
		return fmt.Errorf("parameter natGatewayId is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PublicSubnet) validateAddIpv6DefaultEgressOnlyInternetRouteParameters(gatewayId *string) error {
	if gatewayId == nil {
		return fmt.Errorf("parameter gatewayId is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PublicSubnet) validateAddIpv6DefaultInternetRouteParameters(gatewayId *string) error {
	if gatewayId == nil {
		return fmt.Errorf("parameter gatewayId is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PublicSubnet) validateAddIpv6Nat64RouteParameters(natGatewayId *string) error {
	if natGatewayId == nil {
		return fmt.Errorf("parameter natGatewayId is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PublicSubnet) validateAddRouteParameters(id *string, options *AddRouteOptions) error {
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

func (p *jsiiProxy_PublicSubnet) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PublicSubnet) validateAssociateNetworkAclParameters(id *string, networkAcl INetworkAcl) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if networkAcl == nil {
		return fmt.Errorf("parameter networkAcl is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PublicSubnet) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
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

func (p *jsiiProxy_PublicSubnet) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func validatePublicSubnet_FromPublicSubnetAttributesParameters(scope constructs.Construct, id *string, attrs *PublicSubnetAttributes) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attrs == nil {
		return fmt.Errorf("parameter attrs is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attrs, func() string { return "parameter attrs" }); err != nil {
		return err
	}

	return nil
}

func validatePublicSubnet_FromSubnetAttributesParameters(scope constructs.Construct, id *string, attrs *SubnetAttributes) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attrs == nil {
		return fmt.Errorf("parameter attrs is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attrs, func() string { return "parameter attrs" }); err != nil {
		return err
	}

	return nil
}

func validatePublicSubnet_FromSubnetIdParameters(scope constructs.Construct, id *string, subnetId *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if subnetId == nil {
		return fmt.Errorf("parameter subnetId is required, but nil was provided")
	}

	return nil
}

func validatePublicSubnet_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validatePublicSubnet_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validatePublicSubnet_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validatePublicSubnet_IsVpcSubnetParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewPublicSubnetParameters(scope constructs.Construct, id *string, props *PublicSubnetProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

