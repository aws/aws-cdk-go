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
//   channel := &Channel{
//   	ChannelName: jsii.String("channelName"),
//   	DataSource: &DataSource{
//   		S3DataSource: &S3DataSource{
//   			S3Location: s3Location,
//
//   			// the properties below are optional
//   			AttributeNames: []*string{
//   				jsii.String("attributeNames"),
//   			},
//   			S3DataDistributionType: awscdk.Aws_stepfunctions_tasks.S3DataDistributionType_FULLY_REPLICATED,
//   			S3DataType: awscdk.*Aws_stepfunctions_tasks.S3DataType_MANIFEST_FILE,
//   		},
//   	},
//
//   	// the properties below are optional
//   	CompressionType: awscdk.*Aws_stepfunctions_tasks.CompressionType_NONE,
//   	ContentType: jsii.String("contentType"),
//   	InputMode: awscdk.*Aws_stepfunctions_tasks.InputMode_PIPE,
//   	RecordWrapperType: awscdk.*Aws_stepfunctions_tasks.RecordWrapperType_NONE,
//   	ShuffleConfig: &ShuffleConfig{
//   		Seed: jsii.Number(123),
//   	},
//   }
//
// Experimental.
type Channel struct {
	// Name of the channel.
	// Experimental.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// Location of the channel data.
	// Experimental.
	DataSource *DataSource `field:"required" json:"dataSource" yaml:"dataSource"`
	// Compression type if training data is compressed.
	// Experimental.
	CompressionType CompressionType `field:"optional" json:"compressionType" yaml:"compressionType"`
	// The MIME type of the data.
	// Experimental.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Input mode to use for the data channel in a training job.
	// Experimental.
	InputMode InputMode `field:"optional" json:"inputMode" yaml:"inputMode"`
	// Specify RecordIO as the value when input data is in raw format but the training algorithm requires the RecordIO format.
	//
	// In this case, Amazon SageMaker wraps each individual S3 object in a RecordIO record.
	// If the input data is already in RecordIO format, you don't need to set this attribute.
	// Experimental.
	RecordWrapperType RecordWrapperType `field:"optional" json:"recordWrapperType" yaml:"recordWrapperType"`
	// Shuffle config option for input data in a channel.
	// Experimental.
	ShuffleConfig *ShuffleConfig `field:"optional" json:"shuffleConfig" yaml:"shuffleConfig"`
}

