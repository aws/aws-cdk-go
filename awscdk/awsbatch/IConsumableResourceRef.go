package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConsumableResource.
// Experimental.
type IConsumableResourceRef interface {
	constructs.IConstruct
	// A reference to a ConsumableResource resource.
	// Experimental.
	ConsumableResourceRef() *ConsumableResourceReference
}

// The jsii proxy for IConsumableResourceRef
type jsiiProxy_IConsumableResourceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IConsumableResourceRef) ConsumableResourceRef() *ConsumableResourceReference {
	var returns *ConsumableResourceReference
	_jsii_.Get(
		j,
		"consumableResourceRef",
		&returns,
	)
	return returns
}

