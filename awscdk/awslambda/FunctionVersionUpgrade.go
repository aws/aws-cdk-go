package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Aspect for upgrading function versions when the provided feature flag is enabled.
//
// This can be necessary when the feature flag
// changes the function hash, as such changes must be associated with a new
// version. This aspect will change the function description in these cases,
// which "validates" the new function hash.
//
// Example:
//   stack := awscdk.Newstack()
//   awscdk.Aspects_Of(stack).Add(lambda.NewFunctionVersionUpgrade(awscdklibcxapi.LAMBDA_RECOGNIZE_VERSION_PROPS))
//
type FunctionVersionUpgrade interface {
	awscdk.IAspect
	// All aspects can visit an IConstruct.
	Visit(node constructs.IConstruct)
}

// The jsii proxy struct for FunctionVersionUpgrade
type jsiiProxy_FunctionVersionUpgrade struct {
	internal.Type__awscdkIAspect
}

func NewFunctionVersionUpgrade(featureFlag *string, enabled *bool) FunctionVersionUpgrade {
	_init_.Initialize()

	if err := validateNewFunctionVersionUpgradeParameters(featureFlag); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionVersionUpgrade{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.FunctionVersionUpgrade",
		[]interface{}{featureFlag, enabled},
		&j,
	)

	return &j
}

func NewFunctionVersionUpgrade_Override(f FunctionVersionUpgrade, featureFlag *string, enabled *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.FunctionVersionUpgrade",
		[]interface{}{featureFlag, enabled},
		f,
	)
}

func (f *jsiiProxy_FunctionVersionUpgrade) Visit(node constructs.IConstruct) {
	if err := f.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"visit",
		[]interface{}{node},
	)
}

