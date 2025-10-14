package awsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceNetwork.
// Experimental.
type IServiceNetworkRef interface {
	constructs.IConstruct
	// A reference to a ServiceNetwork resource.
	// Experimental.
	ServiceNetworkRef() *ServiceNetworkReference
}

// The jsii proxy for IServiceNetworkRef
type jsiiProxy_IServiceNetworkRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IServiceNetworkRef) ServiceNetworkRef() *ServiceNetworkReference {
	var returns *ServiceNetworkReference
	_jsii_.Get(
		j,
		"serviceNetworkRef",
		&returns,
	)
	return returns
}

