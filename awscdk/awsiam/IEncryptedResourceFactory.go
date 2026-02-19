package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Factory interface for creating IEncryptedResource instances from constructs.
//
// Implementations of this interface are registered in the DefaultEncryptedResourceFactories
// registry and enable automatic KMS key permission grants for encrypted CloudFormation resources.
// When a grant operation is performed on an encrypted resource, the factory converts L1 constructs
// into resources that can grant permissions on their associated KMS encryption keys.
//
// Factories are typically registered during static initialization and associated with specific
// CloudFormation resource types (e.g., 'AWS::DynamoDB::Table'). The CDK's grant system uses
// these factories to automatically add necessary KMS key permissions when granting access to
// encrypted resources.
type IEncryptedResourceFactory interface {
	// Create an IEncryptedResource from a construct.
	ForResource(resource awscdk.CfnResource) IEncryptedResource
}

// The jsii proxy for IEncryptedResourceFactory
type jsiiProxy_IEncryptedResourceFactory struct {
	_ byte // padding
}

func (i *jsiiProxy_IEncryptedResourceFactory) ForResource(resource awscdk.CfnResource) IEncryptedResource {
	if err := i.validateForResourceParameters(resource); err != nil {
		panic(err)
	}
	var returns IEncryptedResource

	_jsii_.Invoke(
		i,
		"forResource",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

