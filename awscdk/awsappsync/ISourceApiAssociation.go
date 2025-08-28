package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
)

// Interface for AppSync Source Api Association.
type ISourceApiAssociation interface {
	awscdk.IResource
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

