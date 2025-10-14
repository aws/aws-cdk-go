package awsroute53recoveryreadiness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53recoveryreadiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceSet.
// Experimental.
type IResourceSetRef interface {
	constructs.IConstruct
	// A reference to a ResourceSet resource.
	// Experimental.
	ResourceSetRef() *ResourceSetReference
}

// The jsii proxy for IResourceSetRef
type jsiiProxy_IResourceSetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResourceSetRef) ResourceSetRef() *ResourceSetReference {
	var returns *ResourceSetReference
	_jsii_.Get(
		j,
		"resourceSetRef",
		&returns,
	)
	return returns
}

