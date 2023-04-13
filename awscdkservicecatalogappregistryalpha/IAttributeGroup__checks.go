//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::ServiceCatalogAppRegistry
package awscdkservicecatalogappregistryalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IAttributeGroup) validateAssociateWithParameters(application IApplication) error {
	if application == nil {
		return fmt.Errorf("parameter application is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IAttributeGroup) validateShareAttributeGroupParameters(id *string, shareOptions *ShareOptions) error {
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

