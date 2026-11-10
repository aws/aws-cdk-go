package interfacesawsnetworkmonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnetworkmonitor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Probe.
// Experimental.
type IProbeRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Probe resource.
	// Experimental.
	ProbeRef() *ProbeReference
}

// The jsii proxy for IProbeRef
type jsiiProxy_IProbeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IProbeRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IProbeRef) ProbeRef() *ProbeReference {
	var returns *ProbeReference
	_jsii_.Get(
		j,
		"probeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProbeRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProbeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

