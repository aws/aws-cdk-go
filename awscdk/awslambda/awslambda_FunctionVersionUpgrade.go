package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda/internal"
)

// Aspect for upgrading function versions when the feature flag provided feature flag present.
//
// This can be necessary when the feature flag
// changes the function hash, as such changes must be associated with a new
// version. This aspect will change the function description in these cases,
// which "validates" the new function hash.
//
// Example:
//   stack := awscdk.Newstack()
//   awscdk.Aspects.of(stack).add(lambda.NewFunctionVersionUpgrade(monocdkcxapi.LAMBDA_RECOGNIZE_VERSION_PROPS))
//
// Experimental.
type FunctionVersionUpgrade interface {
	awscdk.IAspect
	// All aspects can visit an IConstruct.
	// Experimental.
	Visit(node awscdk.IConstruct)
}

// The jsii proxy struct for FunctionVersionUpgrade
type jsiiProxy_FunctionVersionUpgrade struct {
	internal.Type__awscdkIAspect
}

// Experimental.
func NewFunctionVersionUpgrade(featureFlag *string, enabled *bool) FunctionVersionUpgrade {
	_init_.Initialize()

	j := jsiiProxy_FunctionVersionUpgrade{}

	_jsii_.Create(
		"monocdk.aws_lambda.FunctionVersionUpgrade",
		[]interface{}{featureFlag, enabled},
		&j,
	)

	return &j
}

// Experimental.
func NewFunctionVersionUpgrade_Override(f FunctionVersionUpgrade, featureFlag *string, enabled *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lambda.FunctionVersionUpgrade",
		[]interface{}{featureFlag, enabled},
		f,
	)
}

func (f *jsiiProxy_FunctionVersionUpgrade) Visit(node awscdk.IConstruct) {
	_jsii_.InvokeVoid(
		f,
		"visit",
		[]interface{}{node},
	)
}

