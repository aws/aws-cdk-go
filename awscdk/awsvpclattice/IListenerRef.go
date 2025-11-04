package awsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Listener.
// Experimental.
type IListenerRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Listener resource.
	// Experimental.
	ListenerRef() *ListenerReference
}

// The jsii proxy for IListenerRef
type jsiiProxy_IListenerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IListenerRef) ListenerRef() *ListenerReference {
	var returns *ListenerReference
	_jsii_.Get(
		j,
		"listenerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IListenerRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IListenerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

