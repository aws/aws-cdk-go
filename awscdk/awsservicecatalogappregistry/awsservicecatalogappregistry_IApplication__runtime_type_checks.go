//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsservicecatalogappregistry

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (i *jsiiProxy_IApplication) validateAssociateAttributeGroupParameters(attributeGroup IAttributeGroup) error {
	if attributeGroup == nil {
		return fmt.Errorf("parameter attributeGroup is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IApplication) validateAssociateStackParameters(stack awscdk.Stack) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	return nil
}

