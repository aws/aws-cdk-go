//go:build !no_runtime_type_checking

package customresources

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CustomResourceConfig) validateAddLambdaRuntimeParameters(lambdaRuntime awslambda.Runtime) error {
	if lambdaRuntime == nil {
		return fmt.Errorf("parameter lambdaRuntime is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CustomResourceConfig) validateAddLogRetentionLifetimeParameters(retention awslogs.RetentionDays) error {
	if retention == "" {
		return fmt.Errorf("parameter retention is required, but nil was provided")
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

