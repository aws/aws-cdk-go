//go:build !no_runtime_type_checking

package awsservicecatalog

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (i *jsiiProxy_IProduct) validateAssociateTagOptionsParameters(tagOptions TagOptions) error {
	if tagOptions == nil {
		return fmt.Errorf("parameter tagOptions is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IProduct) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

