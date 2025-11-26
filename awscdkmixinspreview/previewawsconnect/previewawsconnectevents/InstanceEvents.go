package previewawsconnectevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect"
)

// EventBridge event patterns for Instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var instanceRef IInstanceRef
//
//   instanceEvents := awscdkmixinspreview.Events.InstanceEvents_FromInstance(instanceRef)
//
// Experimental.
type InstanceEvents interface {
	// EventBridge event pattern for Instance Amazon Connect Contact Event.
	// Experimental.
	CodeConnectContactPattern(options *InstanceEvents_CodeConnectContact_CodeConnectContactProps) *awsevents.EventPattern
	// EventBridge event pattern for Instance Contact Lens Post Call Rules Matched.
	// Experimental.
	ContactLensPostCallRulesMatchedPattern(options *InstanceEvents_ContactLensPostCallRulesMatched_ContactLensPostCallRulesMatchedProps) *awsevents.EventPattern
	// EventBridge event pattern for Instance Contact Lens Realtime Rules Matched.
	// Experimental.
	ContactLensRealtimeRulesMatchedPattern(options *InstanceEvents_ContactLensRealtimeRulesMatched_ContactLensRealtimeRulesMatchedProps) *awsevents.EventPattern
}

// The jsii proxy struct for InstanceEvents
type jsiiProxy_InstanceEvents struct {
	_ byte // padding
}

// Create InstanceEvents from a Instance reference.
// Experimental.
func InstanceEvents_FromInstance(instanceRef interfacesawsconnect.IInstanceRef) InstanceEvents {
	_init_.Initialize()

	if err := validateInstanceEvents_FromInstanceParameters(instanceRef); err != nil {
		panic(err)
	}
	var returns InstanceEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_connect.events.InstanceEvents",
		"fromInstance",
		[]interface{}{instanceRef},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceEvents) CodeConnectContactPattern(options *InstanceEvents_CodeConnectContact_CodeConnectContactProps) *awsevents.EventPattern {
	if err := i.validateCodeConnectContactPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		i,
		"codeConnectContactPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceEvents) ContactLensPostCallRulesMatchedPattern(options *InstanceEvents_ContactLensPostCallRulesMatched_ContactLensPostCallRulesMatchedProps) *awsevents.EventPattern {
	if err := i.validateContactLensPostCallRulesMatchedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		i,
		"contactLensPostCallRulesMatchedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceEvents) ContactLensRealtimeRulesMatchedPattern(options *InstanceEvents_ContactLensRealtimeRulesMatched_ContactLensRealtimeRulesMatchedProps) *awsevents.EventPattern {
	if err := i.validateContactLensRealtimeRulesMatchedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		i,
		"contactLensRealtimeRulesMatchedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

