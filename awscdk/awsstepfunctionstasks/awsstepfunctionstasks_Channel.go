package awsstepfunctionstasks


// Describes the training, validation or test dataset and the Amazon S3 location where it is stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var s3Location s3Location
//
//   channel := &channel{
//   	channelName: jsii.String("channelName"),
//   	dataSource: &dataSource{
//   		s3DataSource: &s3DataSource{
//   			s3Location: s3Location,
//
//   			// the properties below are optional
//   			attributeNames: []*string{
//   				jsii.String("attributeNames"),
//   			},
//   			s3DataDistributionType: awscdk.Aws_stepfunctions_tasks.s3DataDistributionType_FULLY_REPLICATED,
//   			s3DataType: awscdk.*Aws_stepfunctions_tasks.s3DataType_MANIFEST_FILE,
//   		},
//   	},
//
//   	// the properties below are optional
//   	compressionType: awscdk.*Aws_stepfunctions_tasks.compressionType_NONE,
//   	contentType: jsii.String("contentType"),
//   	inputMode: awscdk.*Aws_stepfunctions_tasks.inputMode_PIPE,
//   	recordWrapperType: awscdk.*Aws_stepfunctions_tasks.recordWrapperType_NONE,
//   	shuffleConfig: &shuffleConfig{
//   		seed: jsii.Number(123),
//   	},
//   }
//
type Channel struct {
	// Name of the channel.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// Location of the channel data.
	DataSource *DataSource `field:"required" json:"dataSource" yaml:"dataSource"`
	// Compression type if training data is compressed.
	CompressionType CompressionType `field:"optional" json:"compressionType" yaml:"compressionType"`
	// The MIME type of the data.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Input mode to use for the data channel in a training job.
	InputMode InputMode `field:"optional" json:"inputMode" yaml:"inputMode"`
	// Specify RecordIO as the value when input data is in raw format but the training algorithm requires the RecordIO format.
	//
	// In this case, Amazon SageMaker wraps each individual S3 object in a RecordIO record.
	// If the input data is already in RecordIO format, you don't need to set this attribute.
	RecordWrapperType RecordWrapperType `field:"optional" json:"recordWrapperType" yaml:"recordWrapperType"`
	// Shuffle config option for input data in a channel.
	ShuffleConfig *ShuffleConfig `field:"optional" json:"shuffleConfig" yaml:"shuffleConfig"`
}

