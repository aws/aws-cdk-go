package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceDefinitionVersion.
// Experimental.
type IResourceDefinitionVersionRef interface {
	constructs.IConstruct
	// A reference to a ResourceDefinitionVersion resource.
	// Experimental.
	ResourceDefinitionVersionRef() *ResourceDefinitionVersionReference
}

// The jsii proxy for IResourceDefinitionVersionRef
type jsiiProxy_IResourceDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResourceDefinitionVersionRef) ResourceDefinitionVersionRef() *ResourceDefinitionVersionReference {
	var returns *ResourceDefinitionVersionReference
	_jsii_.Get(
		j,
		"resourceDefinitionVersionRef",
		&returns,
	)
	return returns
}

