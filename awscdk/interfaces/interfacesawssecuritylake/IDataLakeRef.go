package interfacesawssecuritylake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssecuritylake/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataLake.
// Experimental.
type IDataLakeRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DataLake resource.
	// Experimental.
	DataLakeRef() *DataLakeReference
}

// The jsii proxy for IDataLakeRef
type jsiiProxy_IDataLakeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDataLakeRef) DataLakeRef() *DataLakeReference {
	var returns *DataLakeReference
	_jsii_.Get(
		j,
		"dataLakeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataLakeRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataLakeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

