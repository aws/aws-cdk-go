//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::ServiceCatalogAppRegistry
package awscdkservicecatalogappregistryalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (i *jsiiProxy_IApplication) validateAddAttributeGroupParameters(id *string, attributeGroupProps *AttributeGroupAssociationProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attributeGroupProps == nil {
		return fmt.Errorf("parameter attributeGroupProps is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attributeGroupProps, func() string { return "parameter attributeGroupProps" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IApplication) validateAssociateAllStacksInScopeParameters(construct constructs.Construct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IApplication) validateAssociateApplicationWithStackParameters(stack awscdk.Stack) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IApplication) validateShareApplicationParameters(id *string, shareOptions *ShareOptions) error {
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

