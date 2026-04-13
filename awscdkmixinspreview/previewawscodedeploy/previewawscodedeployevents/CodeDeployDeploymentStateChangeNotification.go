package previewawscodedeployevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.codedeploy@CodeDeployDeploymentStateChangeNotification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeDeployDeploymentStateChangeNotification := awscdkmixinspreview.Events.NewCodeDeployDeploymentStateChangeNotification()
//
// Experimental.
type CodeDeployDeploymentStateChangeNotification interface {
}

// The jsii proxy struct for CodeDeployDeploymentStateChangeNotification
type jsiiProxy_CodeDeployDeploymentStateChangeNotification struct {
	_ byte // padding
}

// Experimental.
func NewCodeDeployDeploymentStateChangeNotification() CodeDeployDeploymentStateChangeNotification {
	_init_.Initialize()

	j := jsiiProxy_CodeDeployDeploymentStateChangeNotification{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codedeploy.events.CodeDeployDeploymentStateChangeNotification",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCodeDeployDeploymentStateChangeNotification_Override(c CodeDeployDeploymentStateChangeNotification) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codedeploy.events.CodeDeployDeploymentStateChangeNotification",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for CodeDeploy Deployment State-change Notification.
// Experimental.
func CodeDeployDeploymentStateChangeNotification_EventPattern(options *CodeDeployDeploymentStateChangeNotification_CodeDeployDeploymentStateChangeNotificationProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCodeDeployDeploymentStateChangeNotification_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codedeploy.events.CodeDeployDeploymentStateChangeNotification",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

