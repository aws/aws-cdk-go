package previewawssagemakerevents


// Type definition for SageMakerTrainingJobStateChangeItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerTrainingJobStateChangeItem := &SageMakerTrainingJobStateChangeItem{
//   	ChannelName: []*string{
//   		jsii.String("channelName"),
//   	},
//   	CompressionType: []*string{
//   		jsii.String("compressionType"),
//   	},
//   	ContentType: []*string{
//   		jsii.String("contentType"),
//   	},
//   	DataSource: &DataSource{
//   		S3DataSource: &S3DataSource{
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
//   	RecordWrapperType: []*string{
//   		jsii.String("recordWrapperType"),
//   	},
//   }
//
// Experimental.
type SageMakerTrainingJobStateChange_SageMakerTrainingJobStateChangeItem struct {
	// ChannelName property.
	//
	// Specify an array of string values to match this event if the actual value of ChannelName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ChannelName *[]*string `field:"optional" json:"channelName" yaml:"channelName"`
	// CompressionType property.
	//
	// Specify an array of string values to match this event if the actual value of CompressionType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CompressionType *[]*string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// ContentType property.
	//
	// Specify an array of string values to match this event if the actual value of ContentType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContentType *[]*string `field:"optional" json:"contentType" yaml:"contentType"`
	// DataSource property.
	//
	// Specify an array of string values to match this event if the actual value of DataSource is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DataSource *SageMakerTrainingJobStateChange_DataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
	// RecordWrapperType property.
	//
	// Specify an array of string values to match this event if the actual value of RecordWrapperType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RecordWrapperType *[]*string `field:"optional" json:"recordWrapperType" yaml:"recordWrapperType"`
}

