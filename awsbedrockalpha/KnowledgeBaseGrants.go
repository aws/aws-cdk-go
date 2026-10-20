package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrock"
)

// Collection of grant methods for a IKnowledgeBaseRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_alpha "github.com/aws/aws-cdk-go/awsbedrockalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var knowledgeBaseRef IKnowledgeBaseRef
//
//   knowledgeBaseGrants := bedrock_alpha.KnowledgeBaseGrants_FromKnowledgeBase(knowledgeBaseRef)
//
// Experimental.
type KnowledgeBaseGrants interface {
	// Experimental.
	Resource() interfacesawsbedrock.IKnowledgeBaseRef
	// Grant the given identity custom permissions.
	// Experimental.
	Actions(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) awsiam.Grant
	// Grant permissions to retrieve the original or extracted content of documents ingested into this knowledge base.
	//
	// `bedrock:GetDocumentContent` requires `bedrock:Retrieve` on the same knowledge base, so both are granted.
	//
	// This will grant the following permissions on the knowledge base:
	//
	//   - bedrock:Retrieve
	// - bedrock:GetDocumentContent.
	// Experimental.
	GetDocumentContent(grantee awsiam.IGrantable) awsiam.Grant
	// Grant permissions to query this knowledge base and retrieve relevant chunks.
	//
	// This will grant the following permissions on the knowledge base:
	//
	// - bedrock:Retrieve.
	// Experimental.
	Retrieve(grantee awsiam.IGrantable) awsiam.Grant
	// Grant permissions to query a knowledge base and generate a response with a foundation model.
	//
	// `bedrock:RetrieveAndGenerate` does not support resource-level permissions, so it is granted on all resources (`*`). Access to a specific knowledge base is controlled by `bedrock:Retrieve`, so also call `retrieve()` on this knowledge base, and grant `bedrock:InvokeModel` on the generation model with `BedrockFoundationModel.grantInvoke()`.
	//
	// This will grant the following permissions:
	//
	// - bedrock:RetrieveAndGenerate.
	// Experimental.
	RetrieveAndGenerate(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for KnowledgeBaseGrants
type jsiiProxy_KnowledgeBaseGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_KnowledgeBaseGrants) Resource() interfacesawsbedrock.IKnowledgeBaseRef {
	var returns interfacesawsbedrock.IKnowledgeBaseRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for KnowledgeBaseGrants.
// Experimental.
func KnowledgeBaseGrants_FromKnowledgeBase(resource interfacesawsbedrock.IKnowledgeBaseRef) KnowledgeBaseGrants {
	_init_.Initialize()

	if err := validateKnowledgeBaseGrants_FromKnowledgeBaseParameters(resource); err != nil {
		panic(err)
	}
	var returns KnowledgeBaseGrants

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.KnowledgeBaseGrants",
		"fromKnowledgeBase",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KnowledgeBaseGrants) Actions(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) awsiam.Grant {
	if err := k.validateActionsParameters(grantee, actions, options); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"actions",
		[]interface{}{grantee, actions, options},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KnowledgeBaseGrants) GetDocumentContent(grantee awsiam.IGrantable) awsiam.Grant {
	if err := k.validateGetDocumentContentParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"getDocumentContent",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KnowledgeBaseGrants) Retrieve(grantee awsiam.IGrantable) awsiam.Grant {
	if err := k.validateRetrieveParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"retrieve",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KnowledgeBaseGrants) RetrieveAndGenerate(grantee awsiam.IGrantable) awsiam.Grant {
	if err := k.validateRetrieveAndGenerateParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"retrieveAndGenerate",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

