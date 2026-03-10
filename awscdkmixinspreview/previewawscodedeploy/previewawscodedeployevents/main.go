package previewawscodedeployevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codedeploy.events.CodeDeployDeploymentStateChangeNotification",
		reflect.TypeOf((*CodeDeployDeploymentStateChangeNotification)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CodeDeployDeploymentStateChangeNotification{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.events.CodeDeployDeploymentStateChangeNotification.CodeDeployDeploymentStateChangeNotificationProps",
		reflect.TypeOf((*CodeDeployDeploymentStateChangeNotification_CodeDeployDeploymentStateChangeNotificationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codedeploy.events.CodeDeployInstanceStateChangeNotification",
		reflect.TypeOf((*CodeDeployInstanceStateChangeNotification)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CodeDeployInstanceStateChangeNotification{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codedeploy.events.CodeDeployInstanceStateChangeNotification.CodeDeployInstanceStateChangeNotificationProps",
		reflect.TypeOf((*CodeDeployInstanceStateChangeNotification_CodeDeployInstanceStateChangeNotificationProps)(nil)).Elem(),
	)
}
