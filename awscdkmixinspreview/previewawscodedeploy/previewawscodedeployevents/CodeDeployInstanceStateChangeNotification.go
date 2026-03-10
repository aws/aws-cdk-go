package previewawscodedeployevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.codedeploy@CodeDeployInstanceStateChangeNotification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeDeployInstanceStateChangeNotification := awscdkmixinspreview.Events.NewCodeDeployInstanceStateChangeNotification()
//
// Experimental.
type CodeDeployInstanceStateChangeNotification interface {
}

// The jsii proxy struct for CodeDeployInstanceStateChangeNotification
type jsiiProxy_CodeDeployInstanceStateChangeNotification struct {
	_ byte // padding
}

// Experimental.
func NewCodeDeployInstanceStateChangeNotification() CodeDeployInstanceStateChangeNotification {
	_init_.Initialize()

	j := jsiiProxy_CodeDeployInstanceStateChangeNotification{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codedeploy.events.CodeDeployInstanceStateChangeNotification",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCodeDeployInstanceStateChangeNotification_Override(c CodeDeployInstanceStateChangeNotification) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codedeploy.events.CodeDeployInstanceStateChangeNotification",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for CodeDeploy Instance State-change Notification.
// Experimental.
func CodeDeployInstanceStateChangeNotification_CodeDeployInstanceStateChangeNotificationPattern(options *CodeDeployInstanceStateChangeNotification_CodeDeployInstanceStateChangeNotificationProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCodeDeployInstanceStateChangeNotification_CodeDeployInstanceStateChangeNotificationPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codedeploy.events.CodeDeployInstanceStateChangeNotification",
		"codeDeployInstanceStateChangeNotificationPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

