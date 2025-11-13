package interfacesawsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SchemaVersion.
// Experimental.
type ISchemaVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SchemaVersion resource.
	// Experimental.
	SchemaVersionRef() *SchemaVersionReference
}

// The jsii proxy for ISchemaVersionRef
type jsiiProxy_ISchemaVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_ISchemaVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchemaVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

