package previewawssagemakerevents


// Type definition for RequestParametersItem_2.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParametersItem2 := &RequestParametersItem2{
//   	ChannelName: []*string{
//   		jsii.String("channelName"),
//   	},
//   	ContentType: []*string{
//   		jsii.String("contentType"),
//   	},
//   	DataSource: &DataSource1{
//   		S3DataSource: &S3DataSource1{
//   			S3DataDistributionType: []*string{
//   				jsii.String("s3DataDistributionType"),
//   			},
//   			S3DataType: []*string{
//   				jsii.String("s3DataType"),
//   			},
//   			S3Uri: []*string{
//   				jsii.String("s3Uri"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type ModelEvents_AWSAPICallViaCloudTrail_RequestParametersItem2 struct {
	// channelName property.
	//
	// Specify an array of string values to match this event if the actual value of channelName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ChannelName *[]*string `field:"optional" json:"channelName" yaml:"channelName"`
	// contentType property.
	//
	// Specify an array of string values to match this event if the actual value of contentType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContentType *[]*string `field:"optional" json:"contentType" yaml:"contentType"`
	// dataSource property.
	//
	// Specify an array of string values to match this event if the actual value of dataSource is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DataSource *ModelEvents_AWSAPICallViaCloudTrail_DataSource1 `field:"optional" json:"dataSource" yaml:"dataSource"`
}

