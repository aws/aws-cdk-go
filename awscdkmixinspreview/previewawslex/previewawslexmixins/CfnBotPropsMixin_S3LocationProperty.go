package previewawslexmixins


// Defines an Amazon S3 bucket location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3LocationProperty := &S3LocationProperty{
//   	S3Bucket: jsii.String("s3Bucket"),
//   	S3ObjectKey: jsii.String("s3ObjectKey"),
//   	S3ObjectVersion: jsii.String("s3ObjectVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-s3location.html
//
type CfnBotPropsMixin_S3LocationProperty struct {
	// The S3 bucket name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-s3location.html#cfn-lex-bot-s3location-s3bucket
	//
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// The path and file name to the object in the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-s3location.html#cfn-lex-bot-s3location-s3objectkey
	//
	S3ObjectKey *string `field:"optional" json:"s3ObjectKey" yaml:"s3ObjectKey"`
	// The version of the object in the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-s3location.html#cfn-lex-bot-s3location-s3objectversion
	//
	S3ObjectVersion *string `field:"optional" json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
}

