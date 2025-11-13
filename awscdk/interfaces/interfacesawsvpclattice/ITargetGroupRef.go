package interfacesawsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TargetGroup.
// Experimental.
type ITargetGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TargetGroup resource.
	// Experimental.
	TargetGroupRef() *TargetGroupReference
}

// The jsii proxy for ITargetGroupRef
type jsiiProxy_ITargetGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITargetGroupRef) TargetGroupRef() *TargetGroupReference {
	var returns *TargetGroupReference
	_jsii_.Get(
		j,
		"targetGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITargetGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITargetGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

