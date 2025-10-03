package awssecuritylake

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecuritylake/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataLake.
// Experimental.
type IDataLakeRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDataLakeRef
type jsiiProxy_IDataLakeRef struct {
	internal.Type__constructsIConstruct
}

