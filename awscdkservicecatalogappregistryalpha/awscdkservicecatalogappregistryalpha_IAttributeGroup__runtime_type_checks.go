//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::ServiceCatalogAppRegistry
package awscdkservicecatalogappregistryalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IAttributeGroup) validateShareAttributeGroupParameters(shareOptions *ShareOptions) error {
	if shareOptions == nil {
		return fmt.Errorf("parameter shareOptions is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(shareOptions, func() string { return "parameter shareOptions" }); err != nil {
		return err
	}

	return nil
}

