package awslakeformation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataCellsFilter.
// Experimental.
type IDataCellsFilterRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDataCellsFilterRef
type jsiiProxy_IDataCellsFilterRef struct {
	internal.Type__constructsIConstruct
}

