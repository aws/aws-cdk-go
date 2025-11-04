package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Trigger.
// Experimental.
type ITriggerRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Trigger resource.
	// Experimental.
	TriggerRef() *TriggerReference
}

// The jsii proxy for ITriggerRef
type jsiiProxy_ITriggerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ITriggerRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITriggerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

