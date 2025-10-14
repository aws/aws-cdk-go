package awsmediapackagev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediapackagev2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OriginEndpoint.
// Experimental.
type IOriginEndpointRef interface {
	constructs.IConstruct
	// A reference to a OriginEndpoint resource.
	// Experimental.
	OriginEndpointRef() *OriginEndpointReference
}

// The jsii proxy for IOriginEndpointRef
type jsiiProxy_IOriginEndpointRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IOriginEndpointRef) OriginEndpointRef() *OriginEndpointReference {
	var returns *OriginEndpointReference
	_jsii_.Get(
		j,
		"originEndpointRef",
		&returns,
	)
	return returns
}

