package awsdocdb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdocdb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBSubnetGroup.
// Experimental.
type IDBSubnetGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DBSubnetGroup resource.
	// Experimental.
	DbSubnetGroupRef() *DBSubnetGroupReference
}

// The jsii proxy for IDBSubnetGroupRef
type jsiiProxy_IDBSubnetGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDBSubnetGroupRef) DbSubnetGroupRef() *DBSubnetGroupReference {
	var returns *DBSubnetGroupReference
	_jsii_.Get(
		j,
		"dbSubnetGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBSubnetGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBSubnetGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

