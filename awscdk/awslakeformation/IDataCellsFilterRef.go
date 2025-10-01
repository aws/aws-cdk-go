package awslakeformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataCellsFilter.
// Experimental.
type IDataCellsFilterRef interface {
	constructs.IConstruct
	// A reference to a DataCellsFilter resource.
	// Experimental.
	DataCellsFilterRef() *DataCellsFilterReference
}

// The jsii proxy for IDataCellsFilterRef
type jsiiProxy_IDataCellsFilterRef struct {
	internal.Type__constructsIConstruct
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

