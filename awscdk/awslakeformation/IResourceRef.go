package awslakeformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Resource.
// Experimental.
type IResourceRef interface {
	constructs.IConstruct
	// A reference to a Resource resource.
	// Experimental.
	ResourceRef() *ResourceReference
}

// The jsii proxy for IResourceRef
type jsiiProxy_IResourceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResourceRef) ResourceRef() *ResourceReference {
	var returns *ResourceReference
	_jsii_.Get(
		j,
		"resourceRef",
		&returns,
	)
	return returns
}

