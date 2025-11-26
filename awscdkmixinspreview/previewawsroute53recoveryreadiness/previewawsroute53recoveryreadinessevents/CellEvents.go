package previewawsroute53recoveryreadinessevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53recoveryreadiness"
)

// EventBridge event patterns for Cell.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cellRef ICellRef
//
//   cellEvents := awscdkmixinspreview.Events.CellEvents_FromCell(cellRef)
//
// Experimental.
type CellEvents interface {
	// EventBridge event pattern for Cell Route 53 Application Recovery Controller cell readiness status change.
	// Experimental.
	Route53ApplicationRecoveryControllerCellReadinessStatusChangePattern(options *CellEvents_Route53ApplicationRecoveryControllerCellReadinessStatusChange_Route53ApplicationRecoveryControllerCellReadinessStatusChangeProps) *awsevents.EventPattern
}

// The jsii proxy struct for CellEvents
type jsiiProxy_CellEvents struct {
	_ byte // padding
}

// Create CellEvents from a Cell reference.
// Experimental.
func CellEvents_FromCell(cellRef interfacesawsroute53recoveryreadiness.ICellRef) CellEvents {
	_init_.Initialize()

	if err := validateCellEvents_FromCellParameters(cellRef); err != nil {
		panic(err)
	}
	var returns CellEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.CellEvents",
		"fromCell",
		[]interface{}{cellRef},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CellEvents) Route53ApplicationRecoveryControllerCellReadinessStatusChangePattern(options *CellEvents_Route53ApplicationRecoveryControllerCellReadinessStatusChange_Route53ApplicationRecoveryControllerCellReadinessStatusChangeProps) *awsevents.EventPattern {
	if err := c.validateRoute53ApplicationRecoveryControllerCellReadinessStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"route53ApplicationRecoveryControllerCellReadinessStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

