package awslex


// Specifies an Amazon S3 bucket for logging audio conversations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3BucketLogDestinationProperty := &S3BucketLogDestinationProperty{
//   	LogPrefix: jsii.String("logPrefix"),
//   	S3BucketArn: jsii.String("s3BucketArn"),
//
//   	// the properties below are optional
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-s3bucketlogdestination.html
//
type CfnBot_S3BucketLogDestinationProperty struct {
	// The S3 prefix to assign to audio log files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-s3bucketlogdestination.html#cfn-lex-bot-s3bucketlogdestination-logprefix
	//
	LogPrefix *string `field:"required" json:"logPrefix" yaml:"logPrefix"`
	// The Amazon Resource Name (ARN) of an Amazon S3 bucket where audio log files are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-s3bucketlogdestination.html#cfn-lex-bot-s3bucketlogdestination-s3bucketarn
	//
	S3BucketArn *string `field:"required" json:"s3BucketArn" yaml:"s3BucketArn"`
	// The Amazon Resource Name (ARN) of an AWS Key Management Service (KMS) key for encrypting audio log files stored in an Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-s3bucketlogdestination.html#cfn-lex-bot-s3bucketlogdestination-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

