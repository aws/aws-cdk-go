package awsram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsram/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceShare.
// Experimental.
type IResourceShareRef interface {
	constructs.IConstruct
	// A reference to a ResourceShare resource.
	// Experimental.
	ResourceShareRef() *ResourceShareReference
}

// The jsii proxy for IResourceShareRef
type jsiiProxy_IResourceShareRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResourceShareRef) ResourceShareRef() *ResourceShareReference {
	var returns *ResourceShareReference
	_jsii_.Get(
		j,
		"resourceShareRef",
		&returns,
	)
	return returns
}

