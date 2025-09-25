package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceDefinition.
// Experimental.
type IResourceDefinitionRef interface {
	constructs.IConstruct
	// A reference to a ResourceDefinition resource.
	// Experimental.
	ResourceDefinitionRef() *ResourceDefinitionReference
}

// The jsii proxy for IResourceDefinitionRef
type jsiiProxy_IResourceDefinitionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResourceDefinitionRef) ResourceDefinitionRef() *ResourceDefinitionReference {
	var returns *ResourceDefinitionReference
	_jsii_.Get(
		j,
		"resourceDefinitionRef",
		&returns,
	)
	return returns
}

