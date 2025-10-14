package awsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceGateway.
// Experimental.
type IResourceGatewayRef interface {
	constructs.IConstruct
	// A reference to a ResourceGateway resource.
	// Experimental.
	ResourceGatewayRef() *ResourceGatewayReference
}

// The jsii proxy for IResourceGatewayRef
type jsiiProxy_IResourceGatewayRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResourceGatewayRef) ResourceGatewayRef() *ResourceGatewayReference {
	var returns *ResourceGatewayReference
	_jsii_.Get(
		j,
		"resourceGatewayRef",
		&returns,
	)
	return returns
}

