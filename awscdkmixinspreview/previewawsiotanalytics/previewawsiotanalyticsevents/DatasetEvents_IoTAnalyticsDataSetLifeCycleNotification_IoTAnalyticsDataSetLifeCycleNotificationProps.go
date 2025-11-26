package previewawsiotanalyticsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Dataset aws.iotanalytics@IoTAnalyticsDataSetLifeCycleNotification event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ioTAnalyticsDataSetLifeCycleNotificationProps := &IoTAnalyticsDataSetLifeCycleNotificationProps{
//   	ContentDeliveryRuleIndex: []*string{
//   		jsii.String("contentDeliveryRuleIndex"),
//   	},
//   	DatasetName: []*string{
//   		jsii.String("datasetName"),
//   	},
//   	EventDetailVersion: []*string{
//   		jsii.String("eventDetailVersion"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	Message: []*string{
//   		jsii.String("message"),
//   	},
//   	State: []*string{
//   		jsii.String("state"),
//   	},
//   	VersionId: []*string{
//   		jsii.String("versionId"),
//   	},
//   }
//
// Experimental.
type DatasetEvents_IoTAnalyticsDataSetLifeCycleNotification_IoTAnalyticsDataSetLifeCycleNotificationProps struct {
	// content-delivery-rule-index property.
	//
	// Specify an array of string values to match this event if the actual value of content-delivery-rule-index is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContentDeliveryRuleIndex *[]*string `field:"optional" json:"contentDeliveryRuleIndex" yaml:"contentDeliveryRuleIndex"`
	// dataset-name property.
	//
	// Specify an array of string values to match this event if the actual value of dataset-name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Dataset reference.
	//
	// Experimental.
	DatasetName *[]*string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// event-detail-version property.
	//
	// Specify an array of string values to match this event if the actual value of event-detail-version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventDetailVersion *[]*string `field:"optional" json:"eventDetailVersion" yaml:"eventDetailVersion"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// message property.
	//
	// Specify an array of string values to match this event if the actual value of message is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Message *[]*string `field:"optional" json:"message" yaml:"message"`
	// state property.
	//
	// Specify an array of string values to match this event if the actual value of state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	State *[]*string `field:"optional" json:"state" yaml:"state"`
	// version-id property.
	//
	// Specify an array of string values to match this event if the actual value of version-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VersionId *[]*string `field:"optional" json:"versionId" yaml:"versionId"`
}

