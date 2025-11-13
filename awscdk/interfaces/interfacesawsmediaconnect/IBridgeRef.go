package interfacesawsmediaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Bridge.
// Experimental.
type IBridgeRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Bridge resource.
	// Experimental.
	BridgeRef() *BridgeReference
}

// The jsii proxy for IBridgeRef
type jsiiProxy_IBridgeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IBridgeRef) BridgeRef() *BridgeReference {
	var returns *BridgeReference
	_jsii_.Get(
		j,
		"bridgeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridgeRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridgeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

