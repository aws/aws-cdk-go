package previewawstransferevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Connector aws.transfer@SFTPConnectorRemoteDeleteFailed event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorRemoteDeleteFailedProps := &SFTPConnectorRemoteDeleteFailedProps{
//   	ConnectorId: []*string{
//   		jsii.String("connectorId"),
//   	},
//   	DeleteId: []*string{
//   		jsii.String("deleteId"),
//   	},
//   	DeletePath: []*string{
//   		jsii.String("deletePath"),
//   	},
//   	EndTimestamp: []*string{
//   		jsii.String("endTimestamp"),
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
//   	FailureCode: []*string{
//   		jsii.String("failureCode"),
//   	},
//   	FailureMessage: []*string{
//   		jsii.String("failureMessage"),
//   	},
//   	Operation: []*string{
//   		jsii.String("operation"),
//   	},
//   	StartTimestamp: []*string{
//   		jsii.String("startTimestamp"),
//   	},
//   	StatusCode: []*string{
//   		jsii.String("statusCode"),
//   	},
//   	Url: []*string{
//   		jsii.String("url"),
//   	},
//   }
//
// Experimental.
type ConnectorEvents_SFTPConnectorRemoteDeleteFailed_SFTPConnectorRemoteDeleteFailedProps struct {
	// connector-id property.
	//
	// Specify an array of string values to match this event if the actual value of connector-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Connector reference.
	//
	// Experimental.
	ConnectorId *[]*string `field:"optional" json:"connectorId" yaml:"connectorId"`
	// delete-id property.
	//
	// Specify an array of string values to match this event if the actual value of delete-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DeleteId *[]*string `field:"optional" json:"deleteId" yaml:"deleteId"`
	// delete-path property.
	//
	// Specify an array of string values to match this event if the actual value of delete-path is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DeletePath *[]*string `field:"optional" json:"deletePath" yaml:"deletePath"`
	// end-timestamp property.
	//
	// Specify an array of string values to match this event if the actual value of end-timestamp is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndTimestamp *[]*string `field:"optional" json:"endTimestamp" yaml:"endTimestamp"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// failure-code property.
	//
	// Specify an array of string values to match this event if the actual value of failure-code is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FailureCode *[]*string `field:"optional" json:"failureCode" yaml:"failureCode"`
	// failure-message property.
	//
	// Specify an array of string values to match this event if the actual value of failure-message is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FailureMessage *[]*string `field:"optional" json:"failureMessage" yaml:"failureMessage"`
	// operation property.
	//
	// Specify an array of string values to match this event if the actual value of operation is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Operation *[]*string `field:"optional" json:"operation" yaml:"operation"`
	// start-timestamp property.
	//
	// Specify an array of string values to match this event if the actual value of start-timestamp is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartTimestamp *[]*string `field:"optional" json:"startTimestamp" yaml:"startTimestamp"`
	// status-code property.
	//
	// Specify an array of string values to match this event if the actual value of status-code is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StatusCode *[]*string `field:"optional" json:"statusCode" yaml:"statusCode"`
	// url property.
	//
	// Specify an array of string values to match this event if the actual value of url is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Url *[]*string `field:"optional" json:"url" yaml:"url"`
}

