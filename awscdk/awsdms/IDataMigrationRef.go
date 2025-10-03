package awsdms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataMigration.
// Experimental.
type IDataMigrationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDataMigrationRef
type jsiiProxy_IDataMigrationRef struct {
	internal.Type__constructsIConstruct
}

