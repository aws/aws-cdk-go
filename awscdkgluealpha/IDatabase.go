package awscdkgluealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkgluealpha/v2/internal"
)

// Experimental.
type IDatabase interface {
	awscdk.IResource
	// The catalog this database belongs to.
	// Experimental.
	Catalog() ICatalog
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

func (j *jsiiProxy_IDatabase) Catalog() ICatalog {
	var returns ICatalog
	_jsii_.Get(
		j,
		"catalog",
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

