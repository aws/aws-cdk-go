package previewawstransferevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Connector aws.transfer@SFTPConnectorDirectoryListingFailed event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorDirectoryListingFailedProps := &SFTPConnectorDirectoryListingFailedProps{
//   	ConnectorId: []*string{
//   		jsii.String("connectorId"),
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
//   	ListingId: []*string{
//   		jsii.String("listingId"),
//   	},
//   	MaxItems: []*string{
//   		jsii.String("maxItems"),
//   	},
//   	OutputDirectoryPath: []*string{
//   		jsii.String("outputDirectoryPath"),
//   	},
//   	RemoteDirectoryPath: []*string{
//   		jsii.String("remoteDirectoryPath"),
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
type ConnectorEvents_SFTPConnectorDirectoryListingFailed_SFTPConnectorDirectoryListingFailedProps struct {
	// connector-id property.
	//
	// Specify an array of string values to match this event if the actual value of connector-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Connector reference.
	//
	// Experimental.
	ConnectorId *[]*string `field:"optional" json:"connectorId" yaml:"connectorId"`
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
	// listing-id property.
	//
	// Specify an array of string values to match this event if the actual value of listing-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ListingId *[]*string `field:"optional" json:"listingId" yaml:"listingId"`
	// max-items property.
	//
	// Specify an array of string values to match this event if the actual value of max-items is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MaxItems *[]*string `field:"optional" json:"maxItems" yaml:"maxItems"`
	// output-directory-path property.
	//
	// Specify an array of string values to match this event if the actual value of output-directory-path is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OutputDirectoryPath *[]*string `field:"optional" json:"outputDirectoryPath" yaml:"outputDirectoryPath"`
	// remote-directory-path property.
	//
	// Specify an array of string values to match this event if the actual value of remote-directory-path is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RemoteDirectoryPath *[]*string `field:"optional" json:"remoteDirectoryPath" yaml:"remoteDirectoryPath"`
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

