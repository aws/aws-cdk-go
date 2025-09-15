package awsroute53recoverycontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53recoverycontrol/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RoutingControl.
// Experimental.
type IRoutingControlRef interface {
	constructs.IConstruct
	// A reference to a RoutingControl resource.
	// Experimental.
	RoutingControlRef() *RoutingControlReference
}

// The jsii proxy for IRoutingControlRef
type jsiiProxy_IRoutingControlRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRoutingControlRef) RoutingControlRef() *RoutingControlReference {
	var returns *RoutingControlReference
	_jsii_.Get(
		j,
		"routingControlRef",
		&returns,
	)
	return returns
}

