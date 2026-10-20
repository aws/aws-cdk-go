package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrock"
)

// Provides read-only reflection on the configuration of a Bedrock knowledge base underlying `CfnKnowledgeBase` resource.
//
// Use `KnowledgeBaseReflection.of()` to obtain an instance from any KnowledgeBase reference.
// All getters read directly from the L1 resource, providing a single source of
// truth regardless of whether the knowledge base was created as an L2, imported, or
// constructed directly as a CfnKnowledgeBase.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_alpha "github.com/aws/aws-cdk-go/awsbedrockalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var knowledgeBaseRef IKnowledgeBaseRef
//
//   knowledgeBaseReflection := bedrock_alpha.KnowledgeBaseReflection_Of(knowledgeBaseRef)
//
// Experimental.
type KnowledgeBaseReflection interface {
	// The underlying L1 knowledge base resource.
	// Experimental.
	KnowledgeBase() awsbedrock.CfnKnowledgeBase
}

// The jsii proxy struct for KnowledgeBaseReflection
type jsiiProxy_KnowledgeBaseReflection struct {
	_ byte // padding
}

func (j *jsiiProxy_KnowledgeBaseReflection) KnowledgeBase() awsbedrock.CfnKnowledgeBase {
	var returns awsbedrock.CfnKnowledgeBase
	_jsii_.Get(
		j,
		"knowledgeBase",
		&returns,
	)
	return returns
}


// Creates a reflection facade for a knowledge base reference.
// Experimental.
func KnowledgeBaseReflection_Of(knowledgeBaseRef interfacesawsbedrock.IKnowledgeBaseRef) KnowledgeBaseReflection {
	_init_.Initialize()

	if err := validateKnowledgeBaseReflection_OfParameters(knowledgeBaseRef); err != nil {
		panic(err)
	}
	var returns KnowledgeBaseReflection

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.KnowledgeBaseReflection",
		"of",
		[]interface{}{knowledgeBaseRef},
		&returns,
	)

	return returns
}

