package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SchemaVersionMetadata.
// Experimental.
type ISchemaVersionMetadataRef interface {
	constructs.IConstruct
	// A reference to a SchemaVersionMetadata resource.
	// Experimental.
	SchemaVersionMetadataRef() *SchemaVersionMetadataReference
}

// The jsii proxy for ISchemaVersionMetadataRef
type jsiiProxy_ISchemaVersionMetadataRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISchemaVersionMetadataRef) SchemaVersionMetadataRef() *SchemaVersionMetadataReference {
	var returns *SchemaVersionMetadataReference
	_jsii_.Get(
		j,
		"schemaVersionMetadataRef",
		&returns,
	)
	return returns
}

