package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsglue/internal"
)

// Experimental.
type ITable interface {
	awscdk.IResource
	// Experimental.
	TableArn() *string
	// Experimental.
	TableName() *string
}

// The jsii proxy for ITable
type jsiiProxy_ITable struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ITable) TableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

