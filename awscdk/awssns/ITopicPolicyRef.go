package awssns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TopicPolicy.
// Experimental.
type ITopicPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITopicPolicyRef
type jsiiProxy_ITopicPolicyRef struct {
	internal.Type__constructsIConstruct
}

