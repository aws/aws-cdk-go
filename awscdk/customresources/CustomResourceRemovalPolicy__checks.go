//go:build !no_runtime_type_checking

package customresources

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CustomResourceRemovalPolicy) validateVisitParameters(node constructs.IConstruct) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func validateNewCustomResourceRemovalPolicyParameters(removalPolicy awscdk.RemovalPolicy) error {
	if removalPolicy == "" {
		return fmt.Errorf("parameter removalPolicy is required, but nil was provided")
	}

	return nil
}

