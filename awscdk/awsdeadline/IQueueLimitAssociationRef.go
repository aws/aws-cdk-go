package awsdeadline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QueueLimitAssociation.
// Experimental.
type IQueueLimitAssociationRef interface {
	constructs.IConstruct
	// A reference to a QueueLimitAssociation resource.
	// Experimental.
	QueueLimitAssociationRef() *QueueLimitAssociationReference
}

// The jsii proxy for IQueueLimitAssociationRef
type jsiiProxy_IQueueLimitAssociationRef struct {
	internal.Type__constructsIConstruct
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

