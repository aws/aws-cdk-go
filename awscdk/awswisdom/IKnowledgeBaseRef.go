package awswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a KnowledgeBase.
// Experimental.
type IKnowledgeBaseRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a KnowledgeBase resource.
	// Experimental.
	KnowledgeBaseRef() *KnowledgeBaseReference
}

// The jsii proxy for IKnowledgeBaseRef
type jsiiProxy_IKnowledgeBaseRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IKnowledgeBaseRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBaseRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

