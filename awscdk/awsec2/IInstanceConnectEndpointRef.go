package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceConnectEndpoint.
// Experimental.
type IInstanceConnectEndpointRef interface {
	constructs.IConstruct
	// A reference to a InstanceConnectEndpoint resource.
	// Experimental.
	InstanceConnectEndpointRef() *InstanceConnectEndpointReference
}

// The jsii proxy for IInstanceConnectEndpointRef
type jsiiProxy_IInstanceConnectEndpointRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IInstanceConnectEndpointRef) InstanceConnectEndpointRef() *InstanceConnectEndpointReference {
	var returns *InstanceConnectEndpointReference
	_jsii_.Get(
		j,
		"instanceConnectEndpointRef",
		&returns,
	)
	return returns
}

