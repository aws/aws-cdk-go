package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/constructs-go/constructs/v10"
)

// Utility class for discovering and registering encrypted resource traits.
//
// This class provides methods to retrieve IEncryptedResource instances from constructs,
// enabling automatic KMS key permission grants during IAM grant operations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptedResources := awscdk.Aws_iam.NewEncryptedResources()
//
type EncryptedResources interface {
}

// The jsii proxy struct for EncryptedResources
type jsiiProxy_EncryptedResources struct {
	_ byte // padding
}

func NewEncryptedResources() EncryptedResources {
	_init_.Initialize()

	j := jsiiProxy_EncryptedResources{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.EncryptedResources",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewEncryptedResources_Override(e EncryptedResources) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.EncryptedResources",
		nil, // no parameters
		e,
	)
}

// Retrieve the IEncryptedResource associated with a construct, if available.
func EncryptedResources_Of(resource interfaces.IEnvironmentAware) IEncryptedResource {
	_init_.Initialize()

	if err := validateEncryptedResources_OfParameters(resource); err != nil {
		panic(err)
	}
	var returns IEncryptedResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.EncryptedResources",
		"of",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Register a factory for a specific CloudFormation resource type and scope.
func EncryptedResources_Register(scope constructs.IConstruct, cfnType *string, factory IEncryptedResourceFactory) {
	_init_.Initialize()

	if err := validateEncryptedResources_RegisterParameters(scope, cfnType, factory); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_iam.EncryptedResources",
		"register",
		[]interface{}{scope, cfnType, factory},
	)
}

