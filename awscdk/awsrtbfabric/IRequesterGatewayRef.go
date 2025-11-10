package awsrtbfabric

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrtbfabric/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RequesterGateway.
// Experimental.
type IRequesterGatewayRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RequesterGateway resource.
	// Experimental.
	RequesterGatewayRef() *RequesterGatewayReference
}

// The jsii proxy for IRequesterGatewayRef
type jsiiProxy_IRequesterGatewayRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IRequesterGatewayRef) RequesterGatewayRef() *RequesterGatewayReference {
	var returns *RequesterGatewayReference
	_jsii_.Get(
		j,
		"requesterGatewayRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRequesterGatewayRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRequesterGatewayRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

