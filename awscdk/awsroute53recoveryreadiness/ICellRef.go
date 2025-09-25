package awsroute53recoveryreadiness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53recoveryreadiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Cell.
// Experimental.
type ICellRef interface {
	constructs.IConstruct
	// A reference to a Cell resource.
	// Experimental.
	CellRef() *CellReference
}

// The jsii proxy for ICellRef
type jsiiProxy_ICellRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICellRef) CellRef() *CellReference {
	var returns *CellReference
	_jsii_.Get(
		j,
		"cellRef",
		&returns,
	)
	return returns
}

