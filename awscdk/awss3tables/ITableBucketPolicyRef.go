package awss3tables

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3tables/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TableBucketPolicy.
// Experimental.
type ITableBucketPolicyRef interface {
	constructs.IConstruct
	// A reference to a TableBucketPolicy resource.
	// Experimental.
	TableBucketPolicyRef() *TableBucketPolicyReference
}

// The jsii proxy for ITableBucketPolicyRef
type jsiiProxy_ITableBucketPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITableBucketPolicyRef) TableBucketPolicyRef() *TableBucketPolicyReference {
	var returns *TableBucketPolicyReference
	_jsii_.Get(
		j,
		"tableBucketPolicyRef",
		&returns,
	)
	return returns
}

