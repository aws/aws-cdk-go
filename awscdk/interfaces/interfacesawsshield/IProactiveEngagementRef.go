package interfacesawsshield

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsshield/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProactiveEngagement.
// Experimental.
type IProactiveEngagementRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ProactiveEngagement resource.
	// Experimental.
	ProactiveEngagementRef() *ProactiveEngagementReference
}

// The jsii proxy for IProactiveEngagementRef
type jsiiProxy_IProactiveEngagementRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IProactiveEngagementRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IProactiveEngagementRef) ProactiveEngagementRef() *ProactiveEngagementReference {
	var returns *ProactiveEngagementReference
	_jsii_.Get(
		j,
		"proactiveEngagementRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProactiveEngagementRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProactiveEngagementRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

