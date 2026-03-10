package previewawsconnectevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.connect@ContactLensRealtimeRulesMatched.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   contactLensRealtimeRulesMatched := awscdkmixinspreview.Events.NewContactLensRealtimeRulesMatched()
//
// Experimental.
type ContactLensRealtimeRulesMatched interface {
}

// The jsii proxy struct for ContactLensRealtimeRulesMatched
type jsiiProxy_ContactLensRealtimeRulesMatched struct {
	_ byte // padding
}

// Experimental.
func NewContactLensRealtimeRulesMatched() ContactLensRealtimeRulesMatched {
	_init_.Initialize()

	j := jsiiProxy_ContactLensRealtimeRulesMatched{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.events.ContactLensRealtimeRulesMatched",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewContactLensRealtimeRulesMatched_Override(c ContactLensRealtimeRulesMatched) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.events.ContactLensRealtimeRulesMatched",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for Contact Lens Realtime Rules Matched.
// Experimental.
func ContactLensRealtimeRulesMatched_ContactLensRealtimeRulesMatchedPattern(options *ContactLensRealtimeRulesMatched_ContactLensRealtimeRulesMatchedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateContactLensRealtimeRulesMatched_ContactLensRealtimeRulesMatchedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_connect.events.ContactLensRealtimeRulesMatched",
		"contactLensRealtimeRulesMatchedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

