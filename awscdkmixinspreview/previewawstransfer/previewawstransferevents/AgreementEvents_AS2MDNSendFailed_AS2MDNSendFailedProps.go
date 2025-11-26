package previewawstransferevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Agreement aws.transfer@AS2MDNSendFailed event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aS2MDNSendFailedProps := &AS2MDNSendFailedProps{
//   	AgreementId: []*string{
//   		jsii.String("agreementId"),
//   	},
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
//   	ClientIp: []*string{
//   		jsii.String("clientIp"),
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
//   	FailureCode: []*string{
//   		jsii.String("failureCode"),
//   	},
//   	FailureMessage: []*string{
//   		jsii.String("failureMessage"),
//   	},
//   	MdnMessageId: []*string{
//   		jsii.String("mdnMessageId"),
//   	},
//   	MdnSubject: []*string{
//   		jsii.String("mdnSubject"),
//   	},
//   	MdnType: []*string{
//   		jsii.String("mdnType"),
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
//   	ServerId: []*string{
//   		jsii.String("serverId"),
//   	},
//   	StartTimestamp: []*string{
//   		jsii.String("startTimestamp"),
//   	},
//   	StatusCode: []*string{
//   		jsii.String("statusCode"),
//   	},
//   }
//
// Experimental.
type AgreementEvents_AS2MDNSendFailed_AS2MDNSendFailedProps struct {
	// agreement-id property.
	//
	// Specify an array of string values to match this event if the actual value of agreement-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Agreement reference.
	//
	// Experimental.
	AgreementId *[]*string `field:"optional" json:"agreementId" yaml:"agreementId"`
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
	// client-ip property.
	//
	// Specify an array of string values to match this event if the actual value of client-ip is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ClientIp *[]*string `field:"optional" json:"clientIp" yaml:"clientIp"`
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
	// mdn-type property.
	//
	// Specify an array of string values to match this event if the actual value of mdn-type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MdnType *[]*string `field:"optional" json:"mdnType" yaml:"mdnType"`
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
	S3Attributes *AgreementEvents_AS2MDNSendFailed_S3Attributes `field:"optional" json:"s3Attributes" yaml:"s3Attributes"`
	// server-id property.
	//
	// Specify an array of string values to match this event if the actual value of server-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ServerId *[]*string `field:"optional" json:"serverId" yaml:"serverId"`
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
}

