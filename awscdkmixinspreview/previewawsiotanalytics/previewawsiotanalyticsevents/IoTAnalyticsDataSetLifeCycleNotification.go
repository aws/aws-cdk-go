package previewawsiotanalyticsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.iotanalytics@IoTAnalyticsDataSetLifeCycleNotification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ioTAnalyticsDataSetLifeCycleNotification := awscdkmixinspreview.Events.NewIoTAnalyticsDataSetLifeCycleNotification()
//
// Experimental.
type IoTAnalyticsDataSetLifeCycleNotification interface {
}

// The jsii proxy struct for IoTAnalyticsDataSetLifeCycleNotification
type jsiiProxy_IoTAnalyticsDataSetLifeCycleNotification struct {
	_ byte // padding
}

// Experimental.
func NewIoTAnalyticsDataSetLifeCycleNotification() IoTAnalyticsDataSetLifeCycleNotification {
	_init_.Initialize()

	j := jsiiProxy_IoTAnalyticsDataSetLifeCycleNotification{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.IoTAnalyticsDataSetLifeCycleNotification",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewIoTAnalyticsDataSetLifeCycleNotification_Override(i IoTAnalyticsDataSetLifeCycleNotification) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.IoTAnalyticsDataSetLifeCycleNotification",
		nil, // no parameters
		i,
	)
}

// EventBridge event pattern for IoT Analytics DataSet Lifecycle Notification.
// Experimental.
func IoTAnalyticsDataSetLifeCycleNotification_IoTAnalyticsDataSetLifeCycleNotificationPattern(options *IoTAnalyticsDataSetLifeCycleNotification_IoTAnalyticsDataSetLifeCycleNotificationProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateIoTAnalyticsDataSetLifeCycleNotification_IoTAnalyticsDataSetLifeCycleNotificationPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.IoTAnalyticsDataSetLifeCycleNotification",
		"ioTAnalyticsDataSetLifeCycleNotificationPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

