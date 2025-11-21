package awslex


// Describes the Amazon S3 bucket name and location for the grammar that is the source for the slot type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grammarSlotTypeSourceProperty := &GrammarSlotTypeSourceProperty{
//   	S3BucketName: jsii.String("s3BucketName"),
//   	S3ObjectKey: jsii.String("s3ObjectKey"),
//
//   	// the properties below are optional
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-grammarslottypesource.html
//
type CfnBot_GrammarSlotTypeSourceProperty struct {
	// The name of the Amazon S3 bucket that contains the grammar source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-grammarslottypesource.html#cfn-lex-bot-grammarslottypesource-s3bucketname
	//
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// The path to the grammar in the Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-grammarslottypesource.html#cfn-lex-bot-grammarslottypesource-s3objectkey
	//
	S3ObjectKey *string `field:"required" json:"s3ObjectKey" yaml:"s3ObjectKey"`
	// The AWS  key required to decrypt the contents of the grammar, if any.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-grammarslottypesource.html#cfn-lex-bot-grammarslottypesource-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

