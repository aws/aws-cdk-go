package previewawsqldbevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_qldb.events.QLDBLedgerStateChange",
		reflect.TypeOf((*QLDBLedgerStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_QLDBLedgerStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_qldb.events.QLDBLedgerStateChange.QLDBLedgerStateChangeProps",
		reflect.TypeOf((*QLDBLedgerStateChange_QLDBLedgerStateChangeProps)(nil)).Elem(),
	)
}
