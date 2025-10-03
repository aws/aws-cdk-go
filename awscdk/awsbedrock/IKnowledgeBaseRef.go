package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a KnowledgeBase.
// Experimental.
type IKnowledgeBaseRef interface {
	constructs.IConstruct
}

// The jsii proxy for IKnowledgeBaseRef
type jsiiProxy_IKnowledgeBaseRef struct {
	internal.Type__constructsIConstruct
}

