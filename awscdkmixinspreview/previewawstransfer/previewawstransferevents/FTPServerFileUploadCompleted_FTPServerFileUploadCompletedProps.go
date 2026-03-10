package previewawstransferevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.transfer@FTPServerFileUploadCompleted event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPServerFileUploadCompletedProps := &FTPServerFileUploadCompletedProps{
//   	Bytes: []*string{
//   		jsii.String("bytes"),
//   	},
//   	ClientIp: []*string{
//   		jsii.String("clientIp"),
//   	},
//   	EndTimestamp: []*string{
//   		jsii.String("endTimestamp"),
//   	},
//   	Etag: []*string{
//   		jsii.String("etag"),
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
//   	FilePath: []*string{
//   		jsii.String("filePath"),
//   	},
//   	Protocol: []*string{
//   		jsii.String("protocol"),
//   	},
//   	ServerId: []*string{
//   		jsii.String("serverId"),
//   	},
//   	SessionId: []*string{
//   		jsii.String("sessionId"),
//   	},
//   	StartTimestamp: []*string{
//   		jsii.String("startTimestamp"),
//   	},
//   	StatusCode: []*string{
//   		jsii.String("statusCode"),
//   	},
//   	Username: []*string{
//   		jsii.String("username"),
//   	},
//   }
//
// Experimental.
type FTPServerFileUploadCompleted_FTPServerFileUploadCompletedProps struct {
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
	// end-timestamp property.
	//
	// Specify an array of string values to match this event if the actual value of end-timestamp is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndTimestamp *[]*string `field:"optional" json:"endTimestamp" yaml:"endTimestamp"`
	// etag property.
	//
	// Specify an array of string values to match this event if the actual value of etag is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Etag *[]*string `field:"optional" json:"etag" yaml:"etag"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// file-path property.
	//
	// Specify an array of string values to match this event if the actual value of file-path is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FilePath *[]*string `field:"optional" json:"filePath" yaml:"filePath"`
	// protocol property.
	//
	// Specify an array of string values to match this event if the actual value of protocol is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Protocol *[]*string `field:"optional" json:"protocol" yaml:"protocol"`
	// server-id property.
	//
	// Specify an array of string values to match this event if the actual value of server-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ServerId *[]*string `field:"optional" json:"serverId" yaml:"serverId"`
	// session-id property.
	//
	// Specify an array of string values to match this event if the actual value of session-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SessionId *[]*string `field:"optional" json:"sessionId" yaml:"sessionId"`
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
	// username property.
	//
	// Specify an array of string values to match this event if the actual value of username is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Username *[]*string `field:"optional" json:"username" yaml:"username"`
}

