package awss3tables

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3tables/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TableBucket.
// Experimental.
type ITableBucketRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TableBucket resource.
	// Experimental.
	TableBucketRef() *TableBucketReference
}

// The jsii proxy for ITableBucketRef
type jsiiProxy_ITableBucketRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ITableBucketRef) TableBucketRef() *TableBucketReference {
	var returns *TableBucketReference
	_jsii_.Get(
		j,
		"tableBucketRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITableBucketRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITableBucketRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

