//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_SecurityGroup) validateAddEgressRuleParameters(peer IPeer, connection Port) error {
	if peer == nil {
		return fmt.Errorf("parameter peer is required, but nil was provided")
	}

	if connection == nil {
		return fmt.Errorf("parameter connection is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityGroup) validateAddIngressRuleParameters(peer IPeer, connection Port) error {
	if peer == nil {
		return fmt.Errorf("parameter peer is required, but nil was provided")
	}

	if connection == nil {
		return fmt.Errorf("parameter connection is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityGroup) validateDetermineRuleScopeParameters(peer IPeer, connection Port, fromTo *string) error {
	if peer == nil {
		return fmt.Errorf("parameter peer is required, but nil was provided")
	}

	if connection == nil {
		return fmt.Errorf("parameter connection is required, but nil was provided")
	}

	if fromTo == nil {
		return fmt.Errorf("parameter fromTo is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
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

func (s *jsiiProxy_SecurityGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func validateSecurityGroup_FromLookupByIdParameters(scope constructs.Construct, id *string, securityGroupId *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if securityGroupId == nil {
		return fmt.Errorf("parameter securityGroupId is required, but nil was provided")
	}

	return nil
}

func validateSecurityGroup_FromLookupByNameParameters(scope constructs.Construct, id *string, securityGroupName *string, vpc IVpc) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if securityGroupName == nil {
		return fmt.Errorf("parameter securityGroupName is required, but nil was provided")
	}

	if vpc == nil {
		return fmt.Errorf("parameter vpc is required, but nil was provided")
	}

	return nil
}

func validateSecurityGroup_FromSecurityGroupIdParameters(scope constructs.Construct, id *string, securityGroupId *string, options *SecurityGroupImportOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if securityGroupId == nil {
		return fmt.Errorf("parameter securityGroupId is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateSecurityGroup_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateSecurityGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateSecurityGroup_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateSecurityGroup_IsSecurityGroupParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewSecurityGroupParameters(scope constructs.Construct, id *string, props *SecurityGroupProps) error {
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

