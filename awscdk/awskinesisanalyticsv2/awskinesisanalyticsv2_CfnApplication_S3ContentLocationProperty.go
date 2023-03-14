package awskinesisanalyticsv2


// The location of an application or a custom artifact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ContentLocationProperty := &s3ContentLocationProperty{
//   	bucketArn: jsii.String("bucketArn"),
//   	fileKey: jsii.String("fileKey"),
//
//   	// the properties below are optional
//   	objectVersion: jsii.String("objectVersion"),
//   }
//
type CfnApplication_S3ContentLocationProperty struct {
	// The Amazon Resource Name (ARN) for the S3 bucket containing the application code.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The file key for the object containing the application code.
	FileKey *string `field:"required" json:"fileKey" yaml:"fileKey"`
	// The version of the object containing the application code.
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

