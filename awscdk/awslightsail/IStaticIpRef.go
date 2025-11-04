package awslightsail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StaticIp.
// Experimental.
type IStaticIpRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a StaticIp resource.
	// Experimental.
	StaticIpRef() *StaticIpReference
}

// The jsii proxy for IStaticIpRef
type jsiiProxy_IStaticIpRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IStaticIpRef) StaticIpRef() *StaticIpReference {
	var returns *StaticIpReference
	_jsii_.Get(
		j,
		"staticIpRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStaticIpRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStaticIpRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

