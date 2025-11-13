package interfacesawsamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsamazonmq/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Broker.
// Experimental.
type IBrokerRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Broker resource.
	// Experimental.
	BrokerRef() *BrokerReference
}

// The jsii proxy for IBrokerRef
type jsiiProxy_IBrokerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IBrokerRef) BrokerRef() *BrokerReference {
	var returns *BrokerReference
	_jsii_.Get(
		j,
		"brokerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrokerRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrokerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

