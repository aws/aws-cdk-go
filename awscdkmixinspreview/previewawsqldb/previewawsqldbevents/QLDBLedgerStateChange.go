package previewawsqldbevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.qldb@QLDBLedgerStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   qLDBLedgerStateChange := awscdkmixinspreview.Events.NewQLDBLedgerStateChange()
//
// Experimental.
type QLDBLedgerStateChange interface {
}

// The jsii proxy struct for QLDBLedgerStateChange
type jsiiProxy_QLDBLedgerStateChange struct {
	_ byte // padding
}

// Experimental.
func NewQLDBLedgerStateChange() QLDBLedgerStateChange {
	_init_.Initialize()

	j := jsiiProxy_QLDBLedgerStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qldb.events.QLDBLedgerStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewQLDBLedgerStateChange_Override(q QLDBLedgerStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qldb.events.QLDBLedgerStateChange",
		nil, // no parameters
		q,
	)
}

// EventBridge event pattern for QLDB Ledger State Change.
// Experimental.
func QLDBLedgerStateChange_QldbLedgerStateChangePattern(options *QLDBLedgerStateChange_QLDBLedgerStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateQLDBLedgerStateChange_QldbLedgerStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_qldb.events.QLDBLedgerStateChange",
		"qldbLedgerStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

