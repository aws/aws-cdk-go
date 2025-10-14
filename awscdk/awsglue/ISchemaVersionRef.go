package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SchemaVersion.
// Experimental.
type ISchemaVersionRef interface {
	constructs.IConstruct
	// A reference to a SchemaVersion resource.
	// Experimental.
	SchemaVersionRef() *SchemaVersionReference
}

// The jsii proxy for ISchemaVersionRef
type jsiiProxy_ISchemaVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISchemaVersionRef) SchemaVersionRef() *SchemaVersionReference {
	var returns *SchemaVersionReference
	_jsii_.Get(
		j,
		"schemaVersionRef",
		&returns,
	)
	return returns
}

