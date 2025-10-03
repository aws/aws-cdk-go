package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SchemaVersion.
// Experimental.
type ISchemaVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISchemaVersionRef
type jsiiProxy_ISchemaVersionRef struct {
	internal.Type__constructsIConstruct
}

