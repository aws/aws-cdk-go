package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StoredQuery.
// Experimental.
type IStoredQueryRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStoredQueryRef
type jsiiProxy_IStoredQueryRef struct {
	internal.Type__constructsIConstruct
}

