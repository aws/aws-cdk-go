package awsdeadline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QueueEnvironment.
// Experimental.
type IQueueEnvironmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IQueueEnvironmentRef
type jsiiProxy_IQueueEnvironmentRef struct {
	internal.Type__constructsIConstruct
}

