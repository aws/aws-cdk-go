package customresources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/customresources/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Manages lambda runtime for AWS-vended custom resources.
//
// This feature is currently experimental.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var runtime Runtime
//
//   customResourceLambdaRuntime := awscdk.Custom_resources.NewCustomResourceLambdaRuntime(runtime)
//
type CustomResourceLambdaRuntime interface {
	awscdk.IAspect
	// All aspects can visit an IConstruct.
	Visit(node constructs.IConstruct)
}

// The jsii proxy struct for CustomResourceLambdaRuntime
type jsiiProxy_CustomResourceLambdaRuntime struct {
	internal.Type__awscdkIAspect
}

func NewCustomResourceLambdaRuntime(lambdaRuntime awslambda.Runtime) CustomResourceLambdaRuntime {
	_init_.Initialize()

	if err := validateNewCustomResourceLambdaRuntimeParameters(lambdaRuntime); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomResourceLambdaRuntime{}

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.CustomResourceLambdaRuntime",
		[]interface{}{lambdaRuntime},
		&j,
	)

	return &j
}

func NewCustomResourceLambdaRuntime_Override(c CustomResourceLambdaRuntime, lambdaRuntime awslambda.Runtime) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.CustomResourceLambdaRuntime",
		[]interface{}{lambdaRuntime},
		c,
	)
}

func (c *jsiiProxy_CustomResourceLambdaRuntime) Visit(node constructs.IConstruct) {
	if err := c.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"visit",
		[]interface{}{node},
	)
}

