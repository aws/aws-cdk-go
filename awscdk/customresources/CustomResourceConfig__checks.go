//go:build !no_runtime_type_checking

package customresources

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CustomResourceConfig) validateAddLogRetentionLifetimeParameters(rentention awslogs.RetentionDays) error {
	if rentention == "" {
		return fmt.Errorf("parameter rentention is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CustomResourceConfig) validateAddRemovalPolicyParameters(removalPolicy awscdk.RemovalPolicy) error {
	if removalPolicy == "" {
		return fmt.Errorf("parameter removalPolicy is required, but nil was provided")
	}

	return nil
}

func validateCustomResourceConfig_OfParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

