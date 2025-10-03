package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TopicRuleDestination.
// Experimental.
type ITopicRuleDestinationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITopicRuleDestinationRef
type jsiiProxy_ITopicRuleDestinationRef struct {
	internal.Type__constructsIConstruct
}

