package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGateway.
// Experimental.
type ITransitGatewayRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TransitGateway resource.
	// Experimental.
	TransitGatewayRef() *TransitGatewayReference
}

// The jsii proxy for ITransitGatewayRef
type jsiiProxy_ITransitGatewayRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ITransitGatewayRef) TransitGatewayRef() *TransitGatewayReference {
	var returns *TransitGatewayReference
	_jsii_.Get(
		j,
		"transitGatewayRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

