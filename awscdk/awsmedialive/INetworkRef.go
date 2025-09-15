package awsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Network.
// Experimental.
type INetworkRef interface {
	constructs.IConstruct
	// A reference to a Network resource.
	// Experimental.
	NetworkRef() *NetworkReference
}

// The jsii proxy for INetworkRef
type jsiiProxy_INetworkRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_INetworkRef) NetworkRef() *NetworkReference {
	var returns *NetworkReference
	_jsii_.Get(
		j,
		"networkRef",
		&returns,
	)
	return returns
}

