package awscdkgluealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkgluealpha/v2/internal"
)

// Experimental.
type IDatabase interface {
	awscdk.IResource
	// The ARN of the catalog.
	// Experimental.
	CatalogArn() *string
	// The catalog id of the database (usually, the AWS account id).
	// Experimental.
	CatalogId() *string
	// The ARN of the database.
	// Experimental.
	DatabaseArn() *string
	// The name of the database.
	// Experimental.
	DatabaseName() *string
}

// The jsii proxy for IDatabase
type jsiiProxy_IDatabase struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IDatabase) CatalogArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabase) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabase) DatabaseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabase) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

