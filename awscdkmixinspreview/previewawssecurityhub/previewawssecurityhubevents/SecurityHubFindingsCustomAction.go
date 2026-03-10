package previewawssecurityhubevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.securityhub@SecurityHubFindingsCustomAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   securityHubFindingsCustomAction := awscdkmixinspreview.Events.NewSecurityHubFindingsCustomAction()
//
// Experimental.
type SecurityHubFindingsCustomAction interface {
}

// The jsii proxy struct for SecurityHubFindingsCustomAction
type jsiiProxy_SecurityHubFindingsCustomAction struct {
	_ byte // padding
}

// Experimental.
func NewSecurityHubFindingsCustomAction() SecurityHubFindingsCustomAction {
	_init_.Initialize()

	j := jsiiProxy_SecurityHubFindingsCustomAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.events.SecurityHubFindingsCustomAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSecurityHubFindingsCustomAction_Override(s SecurityHubFindingsCustomAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.events.SecurityHubFindingsCustomAction",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for Security Hub Findings - Custom Action.
// Experimental.
func SecurityHubFindingsCustomAction_SecurityHubFindingsCustomActionPattern(options *SecurityHubFindingsCustomAction_SecurityHubFindingsCustomActionProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSecurityHubFindingsCustomAction_SecurityHubFindingsCustomActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_securityhub.events.SecurityHubFindingsCustomAction",
		"securityHubFindingsCustomActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

