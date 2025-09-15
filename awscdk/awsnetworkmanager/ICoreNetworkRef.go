package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CoreNetwork.
// Experimental.
type ICoreNetworkRef interface {
	constructs.IConstruct
	// A reference to a CoreNetwork resource.
	// Experimental.
	CoreNetworkRef() *CoreNetworkReference
}

// The jsii proxy for ICoreNetworkRef
type jsiiProxy_ICoreNetworkRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICoreNetworkRef) CoreNetworkRef() *CoreNetworkReference {
	var returns *CoreNetworkReference
	_jsii_.Get(
		j,
		"coreNetworkRef",
		&returns,
	)
	return returns
}

