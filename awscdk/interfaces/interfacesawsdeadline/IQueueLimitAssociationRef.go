package interfacesawsdeadline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QueueLimitAssociation.
// Experimental.
type IQueueLimitAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a QueueLimitAssociation resource.
	// Experimental.
	QueueLimitAssociationRef() *QueueLimitAssociationReference
}

// The jsii proxy for IQueueLimitAssociationRef
type jsiiProxy_IQueueLimitAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IQueueLimitAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IQueueLimitAssociationRef) QueueLimitAssociationRef() *QueueLimitAssociationReference {
	var returns *QueueLimitAssociationReference
	_jsii_.Get(
		j,
		"queueLimitAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueueLimitAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueueLimitAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

