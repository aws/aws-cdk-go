package awsbedrockalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsbedrockalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrock"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an Amazon Bedrock knowledge base of any type, either created with the CDK or imported.
// Experimental.
type IKnowledgeBase interface {
	awsiam.IGrantable
	interfacesawsbedrock.IKnowledgeBaseRef
	awscdk.IResource
	// Adds a statement to the IAM policy of the knowledge base service role.
	//
	// For a referenced knowledge base whose role was not provided, the statement
	// cannot be applied; a warning is emitted with the statement to add to the
	// role manually.
	// Experimental.
	AddToRolePolicy(statement awsiam.PolicyStatement)
	// Grant permissions on this knowledge base to IAM principals.
	// Experimental.
	Grants() KnowledgeBaseGrants
	// The ARN of the knowledge base.
	// Experimental.
	KnowledgeBaseArn() *string
	// The unique identifier of the knowledge base.
	// Experimental.
	KnowledgeBaseId() *string
	// Resolves the underlying `CfnKnowledgeBase` from the construct tree.
	// Experimental.
	Reflections() KnowledgeBaseReflection
	// The service role that Amazon Bedrock assumes to operate the knowledge base.
	//
	// To add permissions to the role, use `addToRolePolicy()` or grant to the
	// knowledge base itself, which is an `IGrantable`.
	//
	// Undefined for imported knowledge bases whose role was not provided.
	// Experimental.
	Role() interfacesawsiam.IRoleRef
}

// The jsii proxy for IKnowledgeBase
type jsiiProxy_IKnowledgeBase struct {
	internal.Type__awsiamIGrantable
	internal.Type__interfacesawsbedrockIKnowledgeBaseRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IKnowledgeBase) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := i.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (i *jsiiProxy_IKnowledgeBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IKnowledgeBase) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IKnowledgeBase) Grants() KnowledgeBaseGrants {
	var returns KnowledgeBaseGrants
	_jsii_.Get(
		j,
		"grants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBase) KnowledgeBaseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBase) KnowledgeBaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBase) Reflections() KnowledgeBaseReflection {
	var returns KnowledgeBaseReflection
	_jsii_.Get(
		j,
		"reflections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBase) Role() interfacesawsiam.IRoleRef {
	var returns interfacesawsiam.IRoleRef
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBase) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBase) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBase) KnowledgeBaseRef() *interfacesawsbedrock.KnowledgeBaseReference {
	var returns *interfacesawsbedrock.KnowledgeBaseReference
	_jsii_.Get(
		j,
		"knowledgeBaseRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

