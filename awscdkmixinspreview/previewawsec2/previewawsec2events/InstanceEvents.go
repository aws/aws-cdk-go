package previewawsec2events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
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
	// EventBridge event pattern for Instance AWS API Call via CloudTrail.
	// Experimental.
	AwsAPICallViaCloudTrailPattern(options *InstanceEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern
	// EventBridge event pattern for Instance EC2 Instance State-change Notification.
	// Experimental.
	EC2InstanceStateChangeNotificationPattern(options *InstanceEvents_EC2InstanceStateChangeNotification_EC2InstanceStateChangeNotificationProps) *awsevents.EventPattern
	// EventBridge event pattern for Instance EC2 Spot Instance Interruption Warning.
	// Experimental.
	EC2SpotInstanceInterruptionWarningPattern(options *InstanceEvents_EC2SpotInstanceInterruptionWarning_EC2SpotInstanceInterruptionWarningProps) *awsevents.EventPattern
}

// The jsii proxy struct for InstanceEvents
type jsiiProxy_InstanceEvents struct {
	_ byte // padding
}

// Create InstanceEvents from a Instance reference.
// Experimental.
func InstanceEvents_FromInstance(instanceRef interfacesawsec2.IInstanceRef) InstanceEvents {
	_init_.Initialize()

	if err := validateInstanceEvents_FromInstanceParameters(instanceRef); err != nil {
		panic(err)
	}
	var returns InstanceEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.events.InstanceEvents",
		"fromInstance",
		[]interface{}{instanceRef},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceEvents) AwsAPICallViaCloudTrailPattern(options *InstanceEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern {
	if err := i.validateAwsAPICallViaCloudTrailPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		i,
		"awsAPICallViaCloudTrailPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceEvents) EC2InstanceStateChangeNotificationPattern(options *InstanceEvents_EC2InstanceStateChangeNotification_EC2InstanceStateChangeNotificationProps) *awsevents.EventPattern {
	if err := i.validateEC2InstanceStateChangeNotificationPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		i,
		"eC2InstanceStateChangeNotificationPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceEvents) EC2SpotInstanceInterruptionWarningPattern(options *InstanceEvents_EC2SpotInstanceInterruptionWarning_EC2SpotInstanceInterruptionWarningProps) *awsevents.EventPattern {
	if err := i.validateEC2SpotInstanceInterruptionWarningPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		i,
		"eC2SpotInstanceInterruptionWarningPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

