//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (n *jsiiProxy_NatInstanceImage) validateGetImageParameters(scope awscdk.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

