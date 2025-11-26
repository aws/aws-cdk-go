package previewawstransferevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Connector aws.transfer@SFTPConnectorFileRetrieveFailed event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorFileRetrieveFailedProps := &SFTPConnectorFileRetrieveFailedProps{
//   	Bytes: []*string{
//   		jsii.String("bytes"),
//   	},
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
//   	FailureCode: []*string{
//   		jsii.String("failureCode"),
//   	},
//   	FailureMessage: []*string{
//   		jsii.String("failureMessage"),
//   	},
//   	FilePath: []*string{
//   		jsii.String("filePath"),
//   	},
//   	FileTransferId: []*string{
//   		jsii.String("fileTransferId"),
//   	},
//   	LocalDirectoryPath: []*string{
//   		jsii.String("localDirectoryPath"),
//   	},
//   	LocalFileLocation: &LocalFileLocation{
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
//   	Operation: []*string{
//   		jsii.String("operation"),
//   	},
//   	StartTimestamp: []*string{
//   		jsii.String("startTimestamp"),
//   	},
//   	StatusCode: []*string{
//   		jsii.String("statusCode"),
//   	},
//   	TransferId: []*string{
//   		jsii.String("transferId"),
//   	},
//   	Url: []*string{
//   		jsii.String("url"),
//   	},
//   }
//
// Experimental.
type ConnectorEvents_SFTPConnectorFileRetrieveFailed_SFTPConnectorFileRetrieveFailedProps struct {
	// bytes property.
	//
	// Specify an array of string values to match this event if the actual value of bytes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Bytes *[]*string `field:"optional" json:"bytes" yaml:"bytes"`
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
	// file-path property.
	//
	// Specify an array of string values to match this event if the actual value of file-path is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FilePath *[]*string `field:"optional" json:"filePath" yaml:"filePath"`
	// file-transfer-id property.
	//
	// Specify an array of string values to match this event if the actual value of file-transfer-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FileTransferId *[]*string `field:"optional" json:"fileTransferId" yaml:"fileTransferId"`
	// local-directory-path property.
	//
	// Specify an array of string values to match this event if the actual value of local-directory-path is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LocalDirectoryPath *[]*string `field:"optional" json:"localDirectoryPath" yaml:"localDirectoryPath"`
	// local-file-location property.
	//
	// Specify an array of string values to match this event if the actual value of local-file-location is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LocalFileLocation *ConnectorEvents_SFTPConnectorFileRetrieveFailed_LocalFileLocation `field:"optional" json:"localFileLocation" yaml:"localFileLocation"`
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
	// transfer-id property.
	//
	// Specify an array of string values to match this event if the actual value of transfer-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TransferId *[]*string `field:"optional" json:"transferId" yaml:"transferId"`
	// url property.
	//
	// Specify an array of string values to match this event if the actual value of url is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Url *[]*string `field:"optional" json:"url" yaml:"url"`
}

