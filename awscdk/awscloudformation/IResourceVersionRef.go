package awscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceVersion.
// Experimental.
type IResourceVersionRef interface {
	constructs.IConstruct
	// A reference to a ResourceVersion resource.
	// Experimental.
	ResourceVersionRef() *ResourceVersionReference
}

// The jsii proxy for IResourceVersionRef
type jsiiProxy_IResourceVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResourceVersionRef) ResourceVersionRef() *ResourceVersionReference {
	var returns *ResourceVersionReference
	_jsii_.Get(
		j,
		"resourceVersionRef",
		&returns,
	)
	return returns
}

