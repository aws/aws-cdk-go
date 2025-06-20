package customresources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Manages AWS-vended Custom Resources.
//
// This feature is currently experimental.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customResourceConfig := awscdk.Custom_resources.CustomResourceConfig_Of(this)
//
type CustomResourceConfig interface {
	// Set the runtime version on AWS-vended custom resources lambdas.
	//
	// This feature is currently experimental.
	AddLambdaRuntime(lambdaRuntime awslambda.Runtime)
	// Set the log retention of AWS-vended custom resource lambdas.
	//
	// This feature is currently experimental.
	AddLogRetentionLifetime(retention awslogs.RetentionDays)
	// Set the removal policy of AWS-vended custom resource logGroup.
	//
	// This feature is currently experimental.
	AddRemovalPolicy(removalPolicy awscdk.RemovalPolicy)
}

// The jsii proxy struct for CustomResourceConfig
type jsiiProxy_CustomResourceConfig struct {
	_ byte // padding
}

// Returns the CustomResourceConfig for this scope.
func CustomResourceConfig_Of(scope constructs.IConstruct) CustomResourceConfig {
	_init_.Initialize()

	if err := validateCustomResourceConfig_OfParameters(scope); err != nil {
		panic(err)
	}
	var returns CustomResourceConfig

	_jsii_.StaticInvoke(
		"aws-cdk-lib.custom_resources.CustomResourceConfig",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResourceConfig) AddLambdaRuntime(lambdaRuntime awslambda.Runtime) {
	if err := c.validateAddLambdaRuntimeParameters(lambdaRuntime); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addLambdaRuntime",
		[]interface{}{lambdaRuntime},
	)
}

func (c *jsiiProxy_CustomResourceConfig) AddLogRetentionLifetime(retention awslogs.RetentionDays) {
	if err := c.validateAddLogRetentionLifetimeParameters(retention); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addLogRetentionLifetime",
		[]interface{}{retention},
	)
}

func (c *jsiiProxy_CustomResourceConfig) AddRemovalPolicy(removalPolicy awscdk.RemovalPolicy) {
	if err := c.validateAddRemovalPolicyParameters(removalPolicy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addRemovalPolicy",
		[]interface{}{removalPolicy},
	)
}

