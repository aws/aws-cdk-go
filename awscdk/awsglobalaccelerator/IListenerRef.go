package awsglobalaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Listener.
// Experimental.
type IListenerRef interface {
	constructs.IConstruct
	// A reference to a Listener resource.
	// Experimental.
	ListenerRef() *ListenerReference
}

// The jsii proxy for IListenerRef
type jsiiProxy_IListenerRef struct {
	internal.Type__constructsIConstruct
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

