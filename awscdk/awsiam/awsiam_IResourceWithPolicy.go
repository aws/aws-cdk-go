package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam/internal"
)

// A resource with a resource policy that can be added to.
// Experimental.
type IResourceWithPolicy interface {
	awscdk.IResource
	// Add a statement to the resource's resource policy.
	// Experimental.
	AddToResourcePolicy(statement PolicyStatement) *AddToResourcePolicyResult
}

// The jsii proxy for IResourceWithPolicy
type jsiiProxy_IResourceWithPolicy struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IResourceWithPolicy) AddToResourcePolicy(statement PolicyStatement) *AddToResourcePolicyResult {
	if err := i.validateAddToResourcePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *AddToResourcePolicyResult

	_jsii_.Invoke(
		i,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

