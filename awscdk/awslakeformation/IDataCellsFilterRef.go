package awslakeformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataCellsFilter.
// Experimental.
type IDataCellsFilterRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DataCellsFilter resource.
	// Experimental.
	DataCellsFilterRef() *DataCellsFilterReference
}

// The jsii proxy for IDataCellsFilterRef
type jsiiProxy_IDataCellsFilterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDataCellsFilterRef) DataCellsFilterRef() *DataCellsFilterReference {
	var returns *DataCellsFilterReference
	_jsii_.Get(
		j,
		"dataCellsFilterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataCellsFilterRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataCellsFilterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

