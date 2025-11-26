package previewawsopsworksevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsopsworks"
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
	// EventBridge event pattern for Instance OpsWorks Alert.
	// Experimental.
	OpsWorksAlertPattern(options *InstanceEvents_OpsWorksAlert_OpsWorksAlertProps) *awsevents.EventPattern
	// EventBridge event pattern for Instance OpsWorks Command State Change.
	// Experimental.
	OpsWorksCommandStateChangePattern(options *InstanceEvents_OpsWorksCommandStateChange_OpsWorksCommandStateChangeProps) *awsevents.EventPattern
	// EventBridge event pattern for Instance OpsWorks Instance State Change.
	// Experimental.
	OpsWorksInstanceStateChangePattern(options *InstanceEvents_OpsWorksInstanceStateChange_OpsWorksInstanceStateChangeProps) *awsevents.EventPattern
}

// The jsii proxy struct for InstanceEvents
type jsiiProxy_InstanceEvents struct {
	_ byte // padding
}

// Create InstanceEvents from a Instance reference.
// Experimental.
func InstanceEvents_FromInstance(instanceRef interfacesawsopsworks.IInstanceRef) InstanceEvents {
	_init_.Initialize()

	if err := validateInstanceEvents_FromInstanceParameters(instanceRef); err != nil {
		panic(err)
	}
	var returns InstanceEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_opsworks.events.InstanceEvents",
		"fromInstance",
		[]interface{}{instanceRef},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceEvents) OpsWorksAlertPattern(options *InstanceEvents_OpsWorksAlert_OpsWorksAlertProps) *awsevents.EventPattern {
	if err := i.validateOpsWorksAlertPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		i,
		"opsWorksAlertPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceEvents) OpsWorksCommandStateChangePattern(options *InstanceEvents_OpsWorksCommandStateChange_OpsWorksCommandStateChangeProps) *awsevents.EventPattern {
	if err := i.validateOpsWorksCommandStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		i,
		"opsWorksCommandStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceEvents) OpsWorksInstanceStateChangePattern(options *InstanceEvents_OpsWorksInstanceStateChange_OpsWorksInstanceStateChangeProps) *awsevents.EventPattern {
	if err := i.validateOpsWorksInstanceStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		i,
		"opsWorksInstanceStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

