package interfacesawsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ThingGroup.
// Experimental.
type IThingGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ThingGroup resource.
	// Experimental.
	ThingGroupRef() *ThingGroupReference
}

// The jsii proxy for IThingGroupRef
type jsiiProxy_IThingGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IThingGroupRef) ThingGroupRef() *ThingGroupReference {
	var returns *ThingGroupReference
	_jsii_.Get(
		j,
		"thingGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IThingGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IThingGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

