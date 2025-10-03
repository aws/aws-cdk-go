package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SchemaVersionMetadata.
// Experimental.
type ISchemaVersionMetadataRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISchemaVersionMetadataRef
type jsiiProxy_ISchemaVersionMetadataRef struct {
	internal.Type__constructsIConstruct
}

