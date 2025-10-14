package awsdeadline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QueueEnvironment.
// Experimental.
type IQueueEnvironmentRef interface {
	constructs.IConstruct
	// A reference to a QueueEnvironment resource.
	// Experimental.
	QueueEnvironmentRef() *QueueEnvironmentReference
}

// The jsii proxy for IQueueEnvironmentRef
type jsiiProxy_IQueueEnvironmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IQueueEnvironmentRef) QueueEnvironmentRef() *QueueEnvironmentReference {
	var returns *QueueEnvironmentReference
	_jsii_.Get(
		j,
		"queueEnvironmentRef",
		&returns,
	)
	return returns
}

