package previewawsconnectevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.connect@ContactLensPostCallRulesMatched event types for Instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   contactLensPostCallRulesMatched := #error#.NewContactLensPostCallRulesMatched()
//
// Experimental.
type InstanceEvents_ContactLensPostCallRulesMatched interface {
}

// The jsii proxy struct for InstanceEvents_ContactLensPostCallRulesMatched
type jsiiProxy_InstanceEvents_ContactLensPostCallRulesMatched struct {
	_ byte // padding
}

// Experimental.
func NewInstanceEvents_ContactLensPostCallRulesMatched() InstanceEvents_ContactLensPostCallRulesMatched {
	_init_.Initialize()

	j := jsiiProxy_InstanceEvents_ContactLensPostCallRulesMatched{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.events.InstanceEvents.ContactLensPostCallRulesMatched",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewInstanceEvents_ContactLensPostCallRulesMatched_Override(i InstanceEvents_ContactLensPostCallRulesMatched) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.events.InstanceEvents.ContactLensPostCallRulesMatched",
		nil, // no parameters
		i,
	)
}

