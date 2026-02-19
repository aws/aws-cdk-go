package interfacesawscloudtrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudtrail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventDataStore.
// Experimental.
type IEventDataStoreRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EventDataStore resource.
	// Experimental.
	EventDataStoreRef() *EventDataStoreReference
}

// The jsii proxy for IEventDataStoreRef
type jsiiProxy_IEventDataStoreRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IEventDataStoreRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IEventDataStoreRef) EventDataStoreRef() *EventDataStoreReference {
	var returns *EventDataStoreReference
	_jsii_.Get(
		j,
		"eventDataStoreRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventDataStoreRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventDataStoreRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

