package awssagemaker


// Describes the input source and the way the transform job consumes it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   transformInputProperty := &TransformInputProperty{
//   	CompressionType: jsii.String("compressionType"),
//   	ContentType: jsii.String("contentType"),
//   	DataSource: &DataSourceProperty{
//   		S3DataSource: &S3DataSourceProperty{
//   			S3DataType: jsii.String("s3DataType"),
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	SplitType: jsii.String("splitType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-transforminput.html
//
type CfnTransformJobPropsMixin_TransformInputProperty struct {
	// If your transform data is compressed, specify the compression type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-transforminput.html#cfn-sagemaker-transformjob-transforminput-compressiontype
	//
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// The multipurpose internet mail extension (MIME) type of the data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-transforminput.html#cfn-sagemaker-transformjob-transforminput-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Describes the location of the channel data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-transforminput.html#cfn-sagemaker-transformjob-transforminput-datasource
	//
	DataSource interface{} `field:"optional" json:"dataSource" yaml:"dataSource"`
	// The method to use to split the transform job's data files into smaller batches.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-transforminput.html#cfn-sagemaker-transformjob-transforminput-splittype
	//
	SplitType *string `field:"optional" json:"splitType" yaml:"splitType"`
}

