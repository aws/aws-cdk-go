package previewawsiotanalyticsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiotanalytics"
)

// EventBridge event patterns for Dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var datasetRef IDatasetRef
//
//   datasetEvents := awscdkmixinspreview.Events.DatasetEvents_FromDataset(datasetRef)
//
// Experimental.
type DatasetEvents interface {
	// EventBridge event pattern for Dataset IoT Analytics DataSet Lifecycle Notification.
	// Experimental.
	IoTAnalyticsDataSetLifeCycleNotificationPattern(options *DatasetEvents_IoTAnalyticsDataSetLifeCycleNotification_IoTAnalyticsDataSetLifeCycleNotificationProps) *awsevents.EventPattern
}

// The jsii proxy struct for DatasetEvents
type jsiiProxy_DatasetEvents struct {
	_ byte // padding
}

// Create DatasetEvents from a Dataset reference.
// Experimental.
func DatasetEvents_FromDataset(datasetRef interfacesawsiotanalytics.IDatasetRef) DatasetEvents {
	_init_.Initialize()

	if err := validateDatasetEvents_FromDatasetParameters(datasetRef); err != nil {
		panic(err)
	}
	var returns DatasetEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.DatasetEvents",
		"fromDataset",
		[]interface{}{datasetRef},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasetEvents) IoTAnalyticsDataSetLifeCycleNotificationPattern(options *DatasetEvents_IoTAnalyticsDataSetLifeCycleNotification_IoTAnalyticsDataSetLifeCycleNotificationProps) *awsevents.EventPattern {
	if err := d.validateIoTAnalyticsDataSetLifeCycleNotificationPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"ioTAnalyticsDataSetLifeCycleNotificationPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

