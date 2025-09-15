package awscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceDefaultVersion.
// Experimental.
type IResourceDefaultVersionRef interface {
	constructs.IConstruct
	// A reference to a ResourceDefaultVersion resource.
	// Experimental.
	ResourceDefaultVersionRef() *ResourceDefaultVersionReference
}

// The jsii proxy for IResourceDefaultVersionRef
type jsiiProxy_IResourceDefaultVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResourceDefaultVersionRef) ResourceDefaultVersionRef() *ResourceDefaultVersionReference {
	var returns *ResourceDefaultVersionReference
	_jsii_.Get(
		j,
		"resourceDefaultVersionRef",
		&returns,
	)
	return returns
}

