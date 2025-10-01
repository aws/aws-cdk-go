package awsdeadline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QueueFleetAssociation.
// Experimental.
type IQueueFleetAssociationRef interface {
	constructs.IConstruct
	// A reference to a QueueFleetAssociation resource.
	// Experimental.
	QueueFleetAssociationRef() *QueueFleetAssociationReference
}

// The jsii proxy for IQueueFleetAssociationRef
type jsiiProxy_IQueueFleetAssociationRef struct {
	internal.Type__constructsIConstruct
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

