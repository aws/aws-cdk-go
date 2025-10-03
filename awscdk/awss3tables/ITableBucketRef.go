package awss3tables

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3tables/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TableBucket.
// Experimental.
type ITableBucketRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITableBucketRef
type jsiiProxy_ITableBucketRef struct {
	internal.Type__constructsIConstruct
}

