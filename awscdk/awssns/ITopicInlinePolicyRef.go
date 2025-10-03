package awssns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TopicInlinePolicy.
// Experimental.
type ITopicInlinePolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITopicInlinePolicyRef
type jsiiProxy_ITopicInlinePolicyRef struct {
	internal.Type__constructsIConstruct
}

