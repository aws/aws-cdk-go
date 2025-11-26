package previewawsopsworksevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_opsworks.events.InstanceEvents",
		reflect.TypeOf((*InstanceEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "opsWorksAlertPattern", GoMethod: "OpsWorksAlertPattern"},
			_jsii_.MemberMethod{JsiiMethod: "opsWorksCommandStateChangePattern", GoMethod: "OpsWorksCommandStateChangePattern"},
			_jsii_.MemberMethod{JsiiMethod: "opsWorksInstanceStateChangePattern", GoMethod: "OpsWorksInstanceStateChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_InstanceEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_opsworks.events.InstanceEvents.OpsWorksAlert",
		reflect.TypeOf((*InstanceEvents_OpsWorksAlert)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_InstanceEvents_OpsWorksAlert{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opsworks.events.InstanceEvents.OpsWorksAlert.OpsWorksAlertProps",
		reflect.TypeOf((*InstanceEvents_OpsWorksAlert_OpsWorksAlertProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_opsworks.events.InstanceEvents.OpsWorksCommandStateChange",
		reflect.TypeOf((*InstanceEvents_OpsWorksCommandStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_InstanceEvents_OpsWorksCommandStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opsworks.events.InstanceEvents.OpsWorksCommandStateChange.OpsWorksCommandStateChangeProps",
		reflect.TypeOf((*InstanceEvents_OpsWorksCommandStateChange_OpsWorksCommandStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_opsworks.events.InstanceEvents.OpsWorksInstanceStateChange",
		reflect.TypeOf((*InstanceEvents_OpsWorksInstanceStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_InstanceEvents_OpsWorksInstanceStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opsworks.events.InstanceEvents.OpsWorksInstanceStateChange.OpsWorksInstanceStateChangeProps",
		reflect.TypeOf((*InstanceEvents_OpsWorksInstanceStateChange_OpsWorksInstanceStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_opsworks.events.StackEvents",
		reflect.TypeOf((*StackEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "opsWorksDeploymentStateChangePattern", GoMethod: "OpsWorksDeploymentStateChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_StackEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_opsworks.events.StackEvents.OpsWorksDeploymentStateChange",
		reflect.TypeOf((*StackEvents_OpsWorksDeploymentStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_StackEvents_OpsWorksDeploymentStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opsworks.events.StackEvents.OpsWorksDeploymentStateChange.OpsWorksDeploymentStateChangeProps",
		reflect.TypeOf((*StackEvents_OpsWorksDeploymentStateChange_OpsWorksDeploymentStateChangeProps)(nil)).Elem(),
	)
}
