//go:build !no_runtime_type_checking

package awscdkservicecatalogappregistryalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_AttributeGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AttributeGroup) validateAssociateWithParameters(application IApplication) error {
	if application == nil {
		return fmt.Errorf("parameter application is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AttributeGroup) validateGenerateUniqueHashParameters(resourceAddress *string) error {
	if resourceAddress == nil {
		return fmt.Errorf("parameter resourceAddress is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AttributeGroup) validateGetAttributeGroupSharePermissionARNParameters(shareOptions *ShareOptions) error {
	if shareOptions == nil {
		return fmt.Errorf("parameter shareOptions is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(shareOptions, func() string { return "parameter shareOptions" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AttributeGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
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

func (a *jsiiProxy_AttributeGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AttributeGroup) validateShareAttributeGroupParameters(id *string, shareOptions *ShareOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if shareOptions == nil {
		return fmt.Errorf("parameter shareOptions is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(shareOptions, func() string { return "parameter shareOptions" }); err != nil {
		return err
	}

	return nil
}

func validateAttributeGroup_FromAttributeGroupArnParameters(scope constructs.Construct, id *string, attributeGroupArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attributeGroupArn == nil {
		return fmt.Errorf("parameter attributeGroupArn is required, but nil was provided")
	}

	return nil
}

func validateAttributeGroup_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateAttributeGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateAttributeGroup_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewAttributeGroupParameters(scope constructs.Construct, id *string, props *AttributeGroupProps) error {
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

