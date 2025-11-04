package awsdocdb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdocdb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBClusterParameterGroup.
// Experimental.
type IDBClusterParameterGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DBClusterParameterGroup resource.
	// Experimental.
	DbClusterParameterGroupRef() *DBClusterParameterGroupReference
}

// The jsii proxy for IDBClusterParameterGroupRef
type jsiiProxy_IDBClusterParameterGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDBClusterParameterGroupRef) DbClusterParameterGroupRef() *DBClusterParameterGroupReference {
	var returns *DBClusterParameterGroupReference
	_jsii_.Get(
		j,
		"dbClusterParameterGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBClusterParameterGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBClusterParameterGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

