package interfacesawsdeadline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QueueFleetAssociation.
// Experimental.
type IQueueFleetAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a QueueFleetAssociation resource.
	// Experimental.
	QueueFleetAssociationRef() *QueueFleetAssociationReference
}

// The jsii proxy for IQueueFleetAssociationRef
type jsiiProxy_IQueueFleetAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IQueueFleetAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IQueueFleetAssociationRef) QueueFleetAssociationRef() *QueueFleetAssociationReference {
	var returns *QueueFleetAssociationReference
	_jsii_.Get(
		j,
		"queueFleetAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueueFleetAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueueFleetAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

