package interfacesawspinpointemail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawspinpointemail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DedicatedIpPool.
// Experimental.
type IDedicatedIpPoolRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DedicatedIpPool resource.
	// Experimental.
	DedicatedIpPoolRef() *DedicatedIpPoolReference
}

// The jsii proxy for IDedicatedIpPoolRef
type jsiiProxy_IDedicatedIpPoolRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDedicatedIpPoolRef) DedicatedIpPoolRef() *DedicatedIpPoolReference {
	var returns *DedicatedIpPoolReference
	_jsii_.Get(
		j,
		"dedicatedIpPoolRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDedicatedIpPoolRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDedicatedIpPoolRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

