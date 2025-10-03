package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataSource.
// Experimental.
type IDataSourceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDataSourceRef
type jsiiProxy_IDataSourceRef struct {
	internal.Type__constructsIConstruct
}

