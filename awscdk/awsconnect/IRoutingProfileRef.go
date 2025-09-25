package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RoutingProfile.
// Experimental.
type IRoutingProfileRef interface {
	constructs.IConstruct
	// A reference to a RoutingProfile resource.
	// Experimental.
	RoutingProfileRef() *RoutingProfileReference
}

// The jsii proxy for IRoutingProfileRef
type jsiiProxy_IRoutingProfileRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRoutingProfileRef) RoutingProfileRef() *RoutingProfileReference {
	var returns *RoutingProfileReference
	_jsii_.Get(
		j,
		"routingProfileRef",
		&returns,
	)
	return returns
}

