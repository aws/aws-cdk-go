package previewawsiotanalyticsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.iotanalytics@IoTAnalyticsDataSetLifeCycleNotification event types for Dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ioTAnalyticsDataSetLifeCycleNotification := #error#.NewIoTAnalyticsDataSetLifeCycleNotification()
//
// Experimental.
type DatasetEvents_IoTAnalyticsDataSetLifeCycleNotification interface {
}

// The jsii proxy struct for DatasetEvents_IoTAnalyticsDataSetLifeCycleNotification
type jsiiProxy_DatasetEvents_IoTAnalyticsDataSetLifeCycleNotification struct {
	_ byte // padding
}

// Experimental.
func NewDatasetEvents_IoTAnalyticsDataSetLifeCycleNotification() DatasetEvents_IoTAnalyticsDataSetLifeCycleNotification {
	_init_.Initialize()

	j := jsiiProxy_DatasetEvents_IoTAnalyticsDataSetLifeCycleNotification{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.DatasetEvents.IoTAnalyticsDataSetLifeCycleNotification",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatasetEvents_IoTAnalyticsDataSetLifeCycleNotification_Override(d DatasetEvents_IoTAnalyticsDataSetLifeCycleNotification) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.DatasetEvents.IoTAnalyticsDataSetLifeCycleNotification",
		nil, // no parameters
		d,
	)
}

