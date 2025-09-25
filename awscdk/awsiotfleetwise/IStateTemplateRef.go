package awsiotfleetwise

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StateTemplate.
// Experimental.
type IStateTemplateRef interface {
	constructs.IConstruct
	// A reference to a StateTemplate resource.
	// Experimental.
	StateTemplateRef() *StateTemplateReference
}

// The jsii proxy for IStateTemplateRef
type jsiiProxy_IStateTemplateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStateTemplateRef) StateTemplateRef() *StateTemplateReference {
	var returns *StateTemplateReference
	_jsii_.Get(
		j,
		"stateTemplateRef",
		&returns,
	)
	return returns
}

