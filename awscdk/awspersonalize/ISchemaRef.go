package awspersonalize

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspersonalize/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Schema.
// Experimental.
type ISchemaRef interface {
	constructs.IConstruct
	// A reference to a Schema resource.
	// Experimental.
	SchemaRef() *SchemaReference
}

// The jsii proxy for ISchemaRef
type jsiiProxy_ISchemaRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISchemaRef) SchemaRef() *SchemaReference {
	var returns *SchemaReference
	_jsii_.Get(
		j,
		"schemaRef",
		&returns,
	)
	return returns
}

