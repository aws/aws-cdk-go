package awsamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsamazonmq/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Broker.
// Experimental.
type IBrokerRef interface {
	constructs.IConstruct
	// A reference to a Broker resource.
	// Experimental.
	BrokerRef() *BrokerReference
}

// The jsii proxy for IBrokerRef
type jsiiProxy_IBrokerRef struct {
	internal.Type__constructsIConstruct
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

