package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a KnowledgeBase.
// Experimental.
type IKnowledgeBaseRef interface {
	constructs.IConstruct
	// A reference to a KnowledgeBase resource.
	// Experimental.
	KnowledgeBaseRef() *KnowledgeBaseReference
}

// The jsii proxy for IKnowledgeBaseRef
type jsiiProxy_IKnowledgeBaseRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IKnowledgeBaseRef) KnowledgeBaseRef() *KnowledgeBaseReference {
	var returns *KnowledgeBaseReference
	_jsii_.Get(
		j,
		"knowledgeBaseRef",
		&returns,
	)
	return returns
}

