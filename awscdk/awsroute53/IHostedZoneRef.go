package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HostedZone.
// Experimental.
type IHostedZoneRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a HostedZone resource.
	// Experimental.
	HostedZoneRef() *HostedZoneReference
}

// The jsii proxy for IHostedZoneRef
type jsiiProxy_IHostedZoneRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IHostedZoneRef) HostedZoneRef() *HostedZoneReference {
	var returns *HostedZoneReference
	_jsii_.Get(
		j,
		"hostedZoneRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHostedZoneRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHostedZoneRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

