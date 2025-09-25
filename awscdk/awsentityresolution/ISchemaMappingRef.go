package awsentityresolution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsentityresolution/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SchemaMapping.
// Experimental.
type ISchemaMappingRef interface {
	constructs.IConstruct
	// A reference to a SchemaMapping resource.
	// Experimental.
	SchemaMappingRef() *SchemaMappingReference
}

// The jsii proxy for ISchemaMappingRef
type jsiiProxy_ISchemaMappingRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISchemaMappingRef) SchemaMappingRef() *SchemaMappingReference {
	var returns *SchemaMappingReference
	_jsii_.Get(
		j,
		"schemaMappingRef",
		&returns,
	)
	return returns
}

