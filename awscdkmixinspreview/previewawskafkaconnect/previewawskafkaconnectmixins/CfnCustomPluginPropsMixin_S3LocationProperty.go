package previewawskafkaconnectmixins


// The location of an object in Amazon S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3LocationProperty := &S3LocationProperty{
//   	BucketArn: jsii.String("bucketArn"),
//   	FileKey: jsii.String("fileKey"),
//   	ObjectVersion: jsii.String("objectVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-customplugin-s3location.html
//
type CfnCustomPluginPropsMixin_S3LocationProperty struct {
	// The Amazon Resource Name (ARN) of an S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-customplugin-s3location.html#cfn-kafkaconnect-customplugin-s3location-bucketarn
	//
	BucketArn *string `field:"optional" json:"bucketArn" yaml:"bucketArn"`
	// The file key for an object in an S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-customplugin-s3location.html#cfn-kafkaconnect-customplugin-s3location-filekey
	//
	FileKey *string `field:"optional" json:"fileKey" yaml:"fileKey"`
	// The version of an object in an S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-customplugin-s3location.html#cfn-kafkaconnect-customplugin-s3location-objectversion
	//
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

