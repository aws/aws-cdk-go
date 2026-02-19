package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/constructs-go/constructs/v10"
)

// Utility class for discovering and managing resource policy traits.
//
// This class provides methods to retrieve IResourceWithPolicyV2 instances from constructs,
// enabling resource-based policy management during IAM grant operations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceWithPolicies := awscdk.Aws_iam.NewResourceWithPolicies()
//
type ResourceWithPolicies interface {
}

// The jsii proxy struct for ResourceWithPolicies
type jsiiProxy_ResourceWithPolicies struct {
	_ byte // padding
}

func NewResourceWithPolicies() ResourceWithPolicies {
	_init_.Initialize()

	j := jsiiProxy_ResourceWithPolicies{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.ResourceWithPolicies",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewResourceWithPolicies_Override(r ResourceWithPolicies) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.ResourceWithPolicies",
		nil, // no parameters
		r,
	)
}

// Retrieve the IResourceWithPolicyV2 associated with a construct, if available.
func ResourceWithPolicies_Of(resource interfaces.IEnvironmentAware) IResourceWithPolicyV2 {
	_init_.Initialize()

	if err := validateResourceWithPolicies_OfParameters(resource); err != nil {
		panic(err)
	}
	var returns IResourceWithPolicyV2

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.ResourceWithPolicies",
		"of",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Register a factory for a specific CloudFormation resource type and scope.
func ResourceWithPolicies_Register(scope constructs.IConstruct, cfnType *string, factory IResourcePolicyFactory) {
	_init_.Initialize()

	if err := validateResourceWithPolicies_RegisterParameters(scope, cfnType, factory); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_iam.ResourceWithPolicies",
		"register",
		[]interface{}{scope, cfnType, factory},
	)
}

