package awscleanrooms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdMappingTable.
// Experimental.
type IIdMappingTableRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a IdMappingTable resource.
	// Experimental.
	IdMappingTableRef() *IdMappingTableReference
}

// The jsii proxy for IIdMappingTableRef
type jsiiProxy_IIdMappingTableRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IIdMappingTableRef) IdMappingTableRef() *IdMappingTableReference {
	var returns *IdMappingTableReference
	_jsii_.Get(
		j,
		"idMappingTableRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdMappingTableRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdMappingTableRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

