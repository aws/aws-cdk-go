package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
)

// A resource with a resource policy that can be added to.
type IResourceWithPolicyV2 interface {
	awscdk.IEnvironmentAware
	// Add a statement to the resource's resource policy.
	AddToResourcePolicy(statement PolicyStatement) *AddToResourcePolicyResult
}

// The jsii proxy for IResourceWithPolicyV2
type jsiiProxy_IResourceWithPolicyV2 struct {
	internal.Type__awscdkIEnvironmentAware
}

func (i *jsiiProxy_IResourceWithPolicyV2) AddToResourcePolicy(statement PolicyStatement) *AddToResourcePolicyResult {
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

