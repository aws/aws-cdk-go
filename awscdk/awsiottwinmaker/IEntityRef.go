package awsiottwinmaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiottwinmaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Entity.
// Experimental.
type IEntityRef interface {
	constructs.IConstruct
	// A reference to a Entity resource.
	// Experimental.
	EntityRef() *EntityReference
}

// The jsii proxy for IEntityRef
type jsiiProxy_IEntityRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEntityRef) EntityRef() *EntityReference {
	var returns *EntityReference
	_jsii_.Get(
		j,
		"entityRef",
		&returns,
	)
	return returns
}

