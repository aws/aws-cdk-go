package awssagemaker


// Configuration for downloading input data from Amazon S3 into the processing container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3InputProperty := &S3InputProperty{
//   	S3DataType: jsii.String("s3DataType"),
//   	S3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	LocalPath: jsii.String("localPath"),
//   	S3CompressionType: jsii.String("s3CompressionType"),
//   	S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   	S3InputMode: jsii.String("s3InputMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-s3input.html
//
type CfnProcessingJob_S3InputProperty struct {
	// Whether you use an `S3Prefix` or a `ManifestFile` for the data type.
	//
	// If you choose `S3Prefix` , `S3Uri` identifies a key name prefix. Amazon SageMaker uses all objects with the specified key name prefix for the processing job. If you choose `ManifestFile` , `S3Uri` identifies an object that is a manifest file containing a list of object keys that you want Amazon SageMaker to use for the processing job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-s3input.html#cfn-sagemaker-processingjob-s3input-s3datatype
	//
	S3DataType *string `field:"required" json:"s3DataType" yaml:"s3DataType"`
	// The URI of the Amazon S3 prefix Amazon SageMaker downloads data required to run a processing job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-s3input.html#cfn-sagemaker-processingjob-s3input-s3uri
	//
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// The local path in your container where you want Amazon SageMaker to write input data to.
	//
	// `LocalPath` is an absolute path to the input data and must begin with `/opt/ml/processing/` . `LocalPath` is a required parameter when `AppManaged` is `False` (default).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-s3input.html#cfn-sagemaker-processingjob-s3input-localpath
	//
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
	// Whether to GZIP-decompress the data in Amazon S3 as it is streamed into the processing container.
	//
	// `Gzip` can only be used when `Pipe` mode is specified as the `S3InputMode` . In `Pipe` mode, Amazon SageMaker streams input data from the source directly to your container without using the EBS volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-s3input.html#cfn-sagemaker-processingjob-s3input-s3compressiontype
	//
	S3CompressionType *string `field:"optional" json:"s3CompressionType" yaml:"s3CompressionType"`
	// Whether to distribute the data from Amazon S3 to all processing instances with `FullyReplicated` , or whether the data from Amazon S3 is sharded by Amazon S3 key, downloading one shard of data to each processing instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-s3input.html#cfn-sagemaker-processingjob-s3input-s3datadistributiontype
	//
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Whether to use `File` or `Pipe` input mode.
	//
	// In File mode, Amazon SageMaker copies the data from the input source onto the local ML storage volume before starting your processing container. This is the most commonly used input mode. In `Pipe` mode, Amazon SageMaker streams input data from the source directly to your processing container into named pipes without using the ML storage volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-s3input.html#cfn-sagemaker-processingjob-s3input-s3inputmode
	//
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
}

