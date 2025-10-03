package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Topic.
// Experimental.
type ITopicRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITopicRef
type jsiiProxy_ITopicRef struct {
	internal.Type__constructsIConstruct
}

