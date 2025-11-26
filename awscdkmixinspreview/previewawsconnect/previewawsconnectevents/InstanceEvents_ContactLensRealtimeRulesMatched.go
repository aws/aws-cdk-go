package previewawsconnectevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.connect@ContactLensRealtimeRulesMatched event types for Instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   contactLensRealtimeRulesMatched := #error#.NewContactLensRealtimeRulesMatched()
//
// Experimental.
type InstanceEvents_ContactLensRealtimeRulesMatched interface {
}

// The jsii proxy struct for InstanceEvents_ContactLensRealtimeRulesMatched
type jsiiProxy_InstanceEvents_ContactLensRealtimeRulesMatched struct {
	_ byte // padding
}

// Experimental.
func NewInstanceEvents_ContactLensRealtimeRulesMatched() InstanceEvents_ContactLensRealtimeRulesMatched {
	_init_.Initialize()

	j := jsiiProxy_InstanceEvents_ContactLensRealtimeRulesMatched{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.events.InstanceEvents.ContactLensRealtimeRulesMatched",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewInstanceEvents_ContactLensRealtimeRulesMatched_Override(i InstanceEvents_ContactLensRealtimeRulesMatched) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.events.InstanceEvents.ContactLensRealtimeRulesMatched",
		nil, // no parameters
		i,
	)
}

