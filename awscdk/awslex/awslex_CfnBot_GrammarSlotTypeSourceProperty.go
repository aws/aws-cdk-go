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
type CfnBot_GrammarSlotTypeSourceProperty struct {
	// The name of the Amazon S3 bucket that contains the grammar source.
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// The path to the grammar in the Amazon S3 bucket.
	S3ObjectKey *string `field:"required" json:"s3ObjectKey" yaml:"s3ObjectKey"`
	// The AWS KMS key required to decrypt the contents of the grammar, if any.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

