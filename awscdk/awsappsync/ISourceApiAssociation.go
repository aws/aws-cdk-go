package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappsync"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for AppSync Source Api Association.
type ISourceApiAssociation interface {
	awscdk.IResource
	interfacesawsappsync.ISourceApiAssociationRef
	// The association arn.
	AssociationArn() *string
	// The association id.
	AssociationId() *string
	// The merged api in the association.
	MergedApi() IGraphqlApi
	// The source api in the association.
	SourceApi() IGraphqlApi
}

// The jsii proxy for ISourceApiAssociation
type jsiiProxy_ISourceApiAssociation struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsappsyncISourceApiAssociationRef
}

func (i *jsiiProxy_ISourceApiAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ISourceApiAssociation) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISourceApiAssociation) AssociationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISourceApiAssociation) AssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISourceApiAssociation) MergedApi() IGraphqlApi {
	var returns IGraphqlApi
	_jsii_.Get(
		j,
		"mergedApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISourceApiAssociation) SourceApi() IGraphqlApi {
	var returns IGraphqlApi
	_jsii_.Get(
		j,
		"sourceApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISourceApiAssociation) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISourceApiAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISourceApiAssociation) SourceApiAssociationRef() *interfacesawsappsync.SourceApiAssociationReference {
	var returns *interfacesawsappsync.SourceApiAssociationReference
	_jsii_.Get(
		j,
		"sourceApiAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISourceApiAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

