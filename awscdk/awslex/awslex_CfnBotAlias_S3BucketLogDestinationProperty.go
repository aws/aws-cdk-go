package awslex


// Specifies an Amazon S3 bucket for logging audio conversations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3BucketLogDestinationProperty := &s3BucketLogDestinationProperty{
//   	logPrefix: jsii.String("logPrefix"),
//   	s3BucketArn: jsii.String("s3BucketArn"),
//
//   	// the properties below are optional
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
type CfnBotAlias_S3BucketLogDestinationProperty struct {
	// The S3 prefix to assign to audio log files.
	LogPrefix *string `field:"required" json:"logPrefix" yaml:"logPrefix"`
	// The Amazon Resource Name (ARN) of an Amazon S3 bucket where audio log files are stored.
	S3BucketArn *string `field:"required" json:"s3BucketArn" yaml:"s3BucketArn"`
	// The Amazon Resource Name (ARN) of an AWS Key Management Service key for encrypting audio log files stored in an S3 bucket.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

