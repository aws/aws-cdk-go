package previewawsimagebuildermixins


// Amazon S3 logging configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3LogsProperty := &S3LogsProperty{
//   	S3BucketName: jsii.String("s3BucketName"),
//   	S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-infrastructureconfiguration-s3logs.html
//
type CfnInfrastructureConfigurationPropsMixin_S3LogsProperty struct {
	// The S3 bucket in which to store the logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-infrastructureconfiguration-s3logs.html#cfn-imagebuilder-infrastructureconfiguration-s3logs-s3bucketname
	//
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// The Amazon S3 path to the bucket where the logs are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-infrastructureconfiguration-s3logs.html#cfn-imagebuilder-infrastructureconfiguration-s3logs-s3keyprefix
	//
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

