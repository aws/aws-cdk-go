package interfacesawsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SchemaVersionMetadata.
// Experimental.
type ISchemaVersionMetadataRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SchemaVersionMetadata resource.
	// Experimental.
	SchemaVersionMetadataRef() *SchemaVersionMetadataReference
}

// The jsii proxy for ISchemaVersionMetadataRef
type jsiiProxy_ISchemaVersionMetadataRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_ISchemaVersionMetadataRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchemaVersionMetadataRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

