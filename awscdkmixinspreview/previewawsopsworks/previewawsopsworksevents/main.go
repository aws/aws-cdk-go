package previewawsopsworksevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksAlert",
		reflect.TypeOf((*OpsWorksAlert)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_OpsWorksAlert{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksAlert.OpsWorksAlertProps",
		reflect.TypeOf((*OpsWorksAlert_OpsWorksAlertProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksCommandStateChange",
		reflect.TypeOf((*OpsWorksCommandStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_OpsWorksCommandStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksCommandStateChange.OpsWorksCommandStateChangeProps",
		reflect.TypeOf((*OpsWorksCommandStateChange_OpsWorksCommandStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksDeploymentStateChange",
		reflect.TypeOf((*OpsWorksDeploymentStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_OpsWorksDeploymentStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksDeploymentStateChange.OpsWorksDeploymentStateChangeProps",
		reflect.TypeOf((*OpsWorksDeploymentStateChange_OpsWorksDeploymentStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksInstanceStateChange",
		reflect.TypeOf((*OpsWorksInstanceStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_OpsWorksInstanceStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksInstanceStateChange.OpsWorksInstanceStateChangeProps",
		reflect.TypeOf((*OpsWorksInstanceStateChange_OpsWorksInstanceStateChangeProps)(nil)).Elem(),
	)
}
