package interfacesawsmediaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BridgeSource.
// Experimental.
type IBridgeSourceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a BridgeSource resource.
	// Experimental.
	BridgeSourceRef() *BridgeSourceReference
}

// The jsii proxy for IBridgeSourceRef
type jsiiProxy_IBridgeSourceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IBridgeSourceRef) BridgeSourceRef() *BridgeSourceReference {
	var returns *BridgeSourceReference
	_jsii_.Get(
		j,
		"bridgeSourceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridgeSourceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridgeSourceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

