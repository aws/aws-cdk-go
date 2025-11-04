package awsneptune

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsneptune/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBParameterGroup.
// Experimental.
type IDBParameterGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DBParameterGroup resource.
	// Experimental.
	DbParameterGroupRef() *DBParameterGroupReference
}

// The jsii proxy for IDBParameterGroupRef
type jsiiProxy_IDBParameterGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDBParameterGroupRef) DbParameterGroupRef() *DBParameterGroupReference {
	var returns *DBParameterGroupReference
	_jsii_.Get(
		j,
		"dbParameterGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBParameterGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBParameterGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

