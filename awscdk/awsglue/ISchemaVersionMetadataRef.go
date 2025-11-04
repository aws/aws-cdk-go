package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SchemaVersionMetadata.
// Experimental.
type ISchemaVersionMetadataRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SchemaVersionMetadata resource.
	// Experimental.
	SchemaVersionMetadataRef() *SchemaVersionMetadataReference
}

// The jsii proxy for ISchemaVersionMetadataRef
type jsiiProxy_ISchemaVersionMetadataRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ISchemaVersionMetadataRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

