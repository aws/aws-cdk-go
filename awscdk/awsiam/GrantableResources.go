package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
)

// Utility methods to check for specific types of grantable resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grantableResources := awscdk.Aws_iam.NewGrantableResources()
//
type GrantableResources interface {
}

// The jsii proxy struct for GrantableResources
type jsiiProxy_GrantableResources struct {
	_ byte // padding
}

func NewGrantableResources() GrantableResources {
	_init_.Initialize()

	j := jsiiProxy_GrantableResources{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.GrantableResources",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewGrantableResources_Override(g GrantableResources) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.GrantableResources",
		nil, // no parameters
		g,
	)
}

// Whether this resource holds data that can be encrypted using a KMS key.
func GrantableResources_IsEncryptedResource(resource interfaces.IEnvironmentAware) *bool {
	_init_.Initialize()

	if err := validateGrantableResources_IsEncryptedResourceParameters(resource); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.GrantableResources",
		"isEncryptedResource",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Whether this resource admits a resource policy.
func GrantableResources_IsResourceWithPolicy(resource interfaces.IEnvironmentAware) *bool {
	_init_.Initialize()

	if err := validateGrantableResources_IsResourceWithPolicyParameters(resource); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.GrantableResources",
		"isResourceWithPolicy",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

