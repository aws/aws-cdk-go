package awskinesis

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StreamConsumer.
// Experimental.
type IStreamConsumerRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStreamConsumerRef
type jsiiProxy_IStreamConsumerRef struct {
	internal.Type__constructsIConstruct
}

