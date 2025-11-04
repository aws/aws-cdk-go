package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NatGateway.
// Experimental.
type INatGatewayRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a NatGateway resource.
	// Experimental.
	NatGatewayRef() *NatGatewayReference
}

// The jsii proxy for INatGatewayRef
type jsiiProxy_INatGatewayRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_INatGatewayRef) NatGatewayRef() *NatGatewayReference {
	var returns *NatGatewayReference
	_jsii_.Get(
		j,
		"natGatewayRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INatGatewayRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INatGatewayRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

