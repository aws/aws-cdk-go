package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourcePolicy.
// Experimental.
type IResourcePolicyRef interface {
	constructs.IConstruct
	// A reference to a ResourcePolicy resource.
	// Experimental.
	ResourcePolicyRef() *ResourcePolicyReference
}

// The jsii proxy for IResourcePolicyRef
type jsiiProxy_IResourcePolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResourcePolicyRef) ResourcePolicyRef() *ResourcePolicyReference {
	var returns *ResourcePolicyReference
	_jsii_.Get(
		j,
		"resourcePolicyRef",
		&returns,
	)
	return returns
}

