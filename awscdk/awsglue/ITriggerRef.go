package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Trigger.
// Experimental.
type ITriggerRef interface {
	constructs.IConstruct
	// A reference to a Trigger resource.
	// Experimental.
	TriggerRef() *TriggerReference
}

// The jsii proxy for ITriggerRef
type jsiiProxy_ITriggerRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITriggerRef) TriggerRef() *TriggerReference {
	var returns *TriggerReference
	_jsii_.Get(
		j,
		"triggerRef",
		&returns,
	)
	return returns
}

