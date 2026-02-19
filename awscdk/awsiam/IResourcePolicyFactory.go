package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Factory interface for creating IResourceWithPolicyV2 instances from constructs.
//
// Implementations of this interface are registered in the DefaultPolicyFactories registry
// and enable automatic resource policy support for CloudFormation resources. When a grant
// operation is performed, the factory converts L1 constructs into resources that support
// resource-based policies.
//
// Factories are typically registered during static initialization and associated with
// specific CloudFormation resource types (e.g., 'AWS::DynamoDB::Table'). The CDK's grant
// system uses these factories to determine whether a resource supports resource policies
// and to create the appropriate wrapper when needed.
type IResourcePolicyFactory interface {
	// Create an IResourceWithPolicyV2 from a construct.
	ForResource(resource awscdk.CfnResource) IResourceWithPolicyV2
}

// The jsii proxy for IResourcePolicyFactory
type jsiiProxy_IResourcePolicyFactory struct {
	_ byte // padding
}

func (i *jsiiProxy_IResourcePolicyFactory) ForResource(resource awscdk.CfnResource) IResourceWithPolicyV2 {
	if err := i.validateForResourceParameters(resource); err != nil {
		panic(err)
	}
	var returns IResourceWithPolicyV2

	_jsii_.Invoke(
		i,
		"forResource",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

