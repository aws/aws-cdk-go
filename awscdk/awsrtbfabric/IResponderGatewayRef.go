package awsrtbfabric

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrtbfabric/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResponderGateway.
// Experimental.
type IResponderGatewayRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ResponderGateway resource.
	// Experimental.
	ResponderGatewayRef() *ResponderGatewayReference
}

// The jsii proxy for IResponderGatewayRef
type jsiiProxy_IResponderGatewayRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IResponderGatewayRef) ResponderGatewayRef() *ResponderGatewayReference {
	var returns *ResponderGatewayReference
	_jsii_.Get(
		j,
		"responderGatewayRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResponderGatewayRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResponderGatewayRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

