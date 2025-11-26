package previewawslexmixins


// Specifies an Amazon S3 bucket for logging audio conversations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3BucketLogDestinationProperty := &S3BucketLogDestinationProperty{
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	LogPrefix: jsii.String("logPrefix"),
//   	S3BucketArn: jsii.String("s3BucketArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-s3bucketlogdestination.html
//
type CfnBotPropsMixin_S3BucketLogDestinationProperty struct {
	// The Amazon Resource Name (ARN) of an AWS Key Management Service (KMS) key for encrypting audio log files stored in an Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-s3bucketlogdestination.html#cfn-lex-bot-s3bucketlogdestination-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The S3 prefix to assign to audio log files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-s3bucketlogdestination.html#cfn-lex-bot-s3bucketlogdestination-logprefix
	//
	LogPrefix *string `field:"optional" json:"logPrefix" yaml:"logPrefix"`
	// The Amazon Resource Name (ARN) of an Amazon S3 bucket where audio log files are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-s3bucketlogdestination.html#cfn-lex-bot-s3bucketlogdestination-s3bucketarn
	//
	S3BucketArn *string `field:"optional" json:"s3BucketArn" yaml:"s3BucketArn"`
}

