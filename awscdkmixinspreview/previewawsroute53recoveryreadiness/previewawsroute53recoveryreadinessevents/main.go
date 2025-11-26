package previewawsroute53recoveryreadinessevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.CellEvents",
		reflect.TypeOf((*CellEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "route53ApplicationRecoveryControllerCellReadinessStatusChangePattern", GoMethod: "Route53ApplicationRecoveryControllerCellReadinessStatusChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_CellEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.CellEvents.Route53ApplicationRecoveryControllerCellReadinessStatusChange",
		reflect.TypeOf((*CellEvents_Route53ApplicationRecoveryControllerCellReadinessStatusChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CellEvents_Route53ApplicationRecoveryControllerCellReadinessStatusChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.CellEvents.Route53ApplicationRecoveryControllerCellReadinessStatusChange.Route53ApplicationRecoveryControllerCellReadinessStatusChangeProps",
		reflect.TypeOf((*CellEvents_Route53ApplicationRecoveryControllerCellReadinessStatusChange_Route53ApplicationRecoveryControllerCellReadinessStatusChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.CellEvents.Route53ApplicationRecoveryControllerCellReadinessStatusChange.State",
		reflect.TypeOf((*CellEvents_Route53ApplicationRecoveryControllerCellReadinessStatusChange_State)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.ReadinessCheckEvents",
		reflect.TypeOf((*ReadinessCheckEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "route53ApplicationRecoveryControllerReadinessCheckStatusChangePattern", GoMethod: "Route53ApplicationRecoveryControllerReadinessCheckStatusChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_ReadinessCheckEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.ReadinessCheckEvents.Route53ApplicationRecoveryControllerReadinessCheckStatusChange",
		reflect.TypeOf((*ReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.ReadinessCheckEvents.Route53ApplicationRecoveryControllerReadinessCheckStatusChange.Route53ApplicationRecoveryControllerReadinessCheckStatusChangeProps",
		reflect.TypeOf((*ReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange_Route53ApplicationRecoveryControllerReadinessCheckStatusChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.ReadinessCheckEvents.Route53ApplicationRecoveryControllerReadinessCheckStatusChange.State",
		reflect.TypeOf((*ReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange_State)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.RecoveryGroupEvents",
		reflect.TypeOf((*RecoveryGroupEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChangePattern", GoMethod: "Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_RecoveryGroupEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.RecoveryGroupEvents.Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange",
		reflect.TypeOf((*RecoveryGroupEvents_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RecoveryGroupEvents_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.RecoveryGroupEvents.Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange.Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChangeProps",
		reflect.TypeOf((*RecoveryGroupEvents_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.RecoveryGroupEvents.Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange.State",
		reflect.TypeOf((*RecoveryGroupEvents_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange_State)(nil)).Elem(),
	)
}
