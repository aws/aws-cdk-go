package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GlobalNetwork.
// Experimental.
type IGlobalNetworkRef interface {
	constructs.IConstruct
	// A reference to a GlobalNetwork resource.
	// Experimental.
	GlobalNetworkRef() *GlobalNetworkReference
}

// The jsii proxy for IGlobalNetworkRef
type jsiiProxy_IGlobalNetworkRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGlobalNetworkRef) GlobalNetworkRef() *GlobalNetworkReference {
	var returns *GlobalNetworkReference
	_jsii_.Get(
		j,
		"globalNetworkRef",
		&returns,
	)
	return returns
}

