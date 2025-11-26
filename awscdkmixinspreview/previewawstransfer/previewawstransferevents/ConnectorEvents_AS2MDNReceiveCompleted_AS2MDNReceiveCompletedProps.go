package previewawstransferevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Connector aws.transfer@AS2MDNReceiveCompleted event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aS2MDNReceiveCompletedProps := &AS2MDNReceiveCompletedProps{
//   	As2From: []*string{
//   		jsii.String("as2From"),
//   	},
//   	As2MessageId: []*string{
//   		jsii.String("as2MessageId"),
//   	},
//   	As2To: []*string{
//   		jsii.String("as2To"),
//   	},
//   	Bytes: []*string{
//   		jsii.String("bytes"),
//   	},
//   	ConnectorId: []*string{
//   		jsii.String("connectorId"),
//   	},
//   	Disposition: []*string{
//   		jsii.String("disposition"),
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
//   	MdnMessageId: []*string{
//   		jsii.String("mdnMessageId"),
//   	},
//   	MdnSubject: []*string{
//   		jsii.String("mdnSubject"),
//   	},
//   	MessageSubject: []*string{
//   		jsii.String("messageSubject"),
//   	},
//   	RequesterFileName: []*string{
//   		jsii.String("requesterFileName"),
//   	},
//   	S3Attributes: &S3Attributes{
//   		FileBucket: []*string{
//   			jsii.String("fileBucket"),
//   		},
//   		FileKey: []*string{
//   			jsii.String("fileKey"),
//   		},
//   		JsonBucket: []*string{
//   			jsii.String("jsonBucket"),
//   		},
//   		JsonKey: []*string{
//   			jsii.String("jsonKey"),
//   		},
//   		MdnBucket: []*string{
//   			jsii.String("mdnBucket"),
//   		},
//   		MdnKey: []*string{
//   			jsii.String("mdnKey"),
//   		},
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
//   }
//
// Experimental.
type ConnectorEvents_AS2MDNReceiveCompleted_AS2MDNReceiveCompletedProps struct {
	// as2-from property.
	//
	// Specify an array of string values to match this event if the actual value of as2-from is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	As2From *[]*string `field:"optional" json:"as2From" yaml:"as2From"`
	// as2-message-id property.
	//
	// Specify an array of string values to match this event if the actual value of as2-message-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	As2MessageId *[]*string `field:"optional" json:"as2MessageId" yaml:"as2MessageId"`
	// as2-to property.
	//
	// Specify an array of string values to match this event if the actual value of as2-to is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	As2To *[]*string `field:"optional" json:"as2To" yaml:"as2To"`
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
	// disposition property.
	//
	// Specify an array of string values to match this event if the actual value of disposition is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Disposition *[]*string `field:"optional" json:"disposition" yaml:"disposition"`
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
	// mdn-message-id property.
	//
	// Specify an array of string values to match this event if the actual value of mdn-message-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MdnMessageId *[]*string `field:"optional" json:"mdnMessageId" yaml:"mdnMessageId"`
	// mdn-subject property.
	//
	// Specify an array of string values to match this event if the actual value of mdn-subject is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MdnSubject *[]*string `field:"optional" json:"mdnSubject" yaml:"mdnSubject"`
	// message-subject property.
	//
	// Specify an array of string values to match this event if the actual value of message-subject is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MessageSubject *[]*string `field:"optional" json:"messageSubject" yaml:"messageSubject"`
	// requester-file-name property.
	//
	// Specify an array of string values to match this event if the actual value of requester-file-name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequesterFileName *[]*string `field:"optional" json:"requesterFileName" yaml:"requesterFileName"`
	// s3-attributes property.
	//
	// Specify an array of string values to match this event if the actual value of s3-attributes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	S3Attributes *ConnectorEvents_AS2MDNReceiveCompleted_S3Attributes `field:"optional" json:"s3Attributes" yaml:"s3Attributes"`
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
}

