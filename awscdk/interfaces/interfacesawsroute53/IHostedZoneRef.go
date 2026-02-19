package interfacesawsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HostedZone.
// Experimental.
type IHostedZoneRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a HostedZone resource.
	// Experimental.
	HostedZoneRef() *HostedZoneReference
}

// The jsii proxy for IHostedZoneRef
type jsiiProxy_IHostedZoneRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IHostedZoneRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IHostedZoneRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

