package awsiottwinmaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiottwinmaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ComponentType.
// Experimental.
type IComponentTypeRef interface {
	constructs.IConstruct
	// A reference to a ComponentType resource.
	// Experimental.
	ComponentTypeRef() *ComponentTypeReference
}

// The jsii proxy for IComponentTypeRef
type jsiiProxy_IComponentTypeRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IComponentTypeRef) ComponentTypeRef() *ComponentTypeReference {
	var returns *ComponentTypeReference
	_jsii_.Get(
		j,
		"componentTypeRef",
		&returns,
	)
	return returns
}

