package previewawsconnectevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.connect@ContactLensPostCallRulesMatched.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   contactLensPostCallRulesMatched := awscdkmixinspreview.Events.NewContactLensPostCallRulesMatched()
//
// Experimental.
type ContactLensPostCallRulesMatched interface {
}

// The jsii proxy struct for ContactLensPostCallRulesMatched
type jsiiProxy_ContactLensPostCallRulesMatched struct {
	_ byte // padding
}

// Experimental.
func NewContactLensPostCallRulesMatched() ContactLensPostCallRulesMatched {
	_init_.Initialize()

	j := jsiiProxy_ContactLensPostCallRulesMatched{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.events.ContactLensPostCallRulesMatched",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewContactLensPostCallRulesMatched_Override(c ContactLensPostCallRulesMatched) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.events.ContactLensPostCallRulesMatched",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for Contact Lens Post Call Rules Matched.
// Experimental.
func ContactLensPostCallRulesMatched_EventPattern(options *ContactLensPostCallRulesMatched_ContactLensPostCallRulesMatchedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateContactLensPostCallRulesMatched_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_connect.events.ContactLensPostCallRulesMatched",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

