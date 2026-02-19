package interfacesawslocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslocation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Tracker.
// Experimental.
type ITrackerRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Tracker resource.
	// Experimental.
	TrackerRef() *TrackerReference
}

// The jsii proxy for ITrackerRef
type jsiiProxy_ITrackerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITrackerRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ITrackerRef) TrackerRef() *TrackerReference {
	var returns *TrackerReference
	_jsii_.Get(
		j,
		"trackerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrackerRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrackerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

