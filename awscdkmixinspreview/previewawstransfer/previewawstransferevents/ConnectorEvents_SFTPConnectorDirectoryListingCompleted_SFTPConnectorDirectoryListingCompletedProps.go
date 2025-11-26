package previewawstransferevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Connector aws.transfer@SFTPConnectorDirectoryListingCompleted event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorDirectoryListingCompletedProps := &SFTPConnectorDirectoryListingCompletedProps{
//   	ConnectorId: []*string{
//   		jsii.String("connectorId"),
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
//   	ItemCount: []*string{
//   		jsii.String("itemCount"),
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
//   	OutputFileLocation: &OutputFileLocation{
//   		Bucket: []*string{
//   			jsii.String("bucket"),
//   		},
//   		Domain: []*string{
//   			jsii.String("domain"),
//   		},
//   		Key: []*string{
//   			jsii.String("key"),
//   		},
//   	},
//   	RemoteDirectoryPath: []*string{
//   		jsii.String("remoteDirectoryPath"),
//   	},
//   	StartTimestamp: []*string{
//   		jsii.String("startTimestamp"),
//   	},
//   	StatusCode: []*string{
//   		jsii.String("statusCode"),
//   	},
//   	Truncated: []*string{
//   		jsii.String("truncated"),
//   	},
//   	Url: []*string{
//   		jsii.String("url"),
//   	},
//   }
//
// Experimental.
type ConnectorEvents_SFTPConnectorDirectoryListingCompleted_SFTPConnectorDirectoryListingCompletedProps struct {
	// connector-id property.
	//
	// Specify an array of string values to match this event if the actual value of connector-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Connector reference.
	//
	// Experimental.
	ConnectorId *[]*string `field:"optional" json:"connectorId" yaml:"connectorId"`
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
	// item-count property.
	//
	// Specify an array of string values to match this event if the actual value of item-count is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ItemCount *[]*string `field:"optional" json:"itemCount" yaml:"itemCount"`
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
	// output-file-location property.
	//
	// Specify an array of string values to match this event if the actual value of output-file-location is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OutputFileLocation *ConnectorEvents_SFTPConnectorDirectoryListingCompleted_OutputFileLocation `field:"optional" json:"outputFileLocation" yaml:"outputFileLocation"`
	// remote-directory-path property.
	//
	// Specify an array of string values to match this event if the actual value of remote-directory-path is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RemoteDirectoryPath *[]*string `field:"optional" json:"remoteDirectoryPath" yaml:"remoteDirectoryPath"`
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
	// truncated property.
	//
	// Specify an array of string values to match this event if the actual value of truncated is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Truncated *[]*string `field:"optional" json:"truncated" yaml:"truncated"`
	// url property.
	//
	// Specify an array of string values to match this event if the actual value of url is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Url *[]*string `field:"optional" json:"url" yaml:"url"`
}

