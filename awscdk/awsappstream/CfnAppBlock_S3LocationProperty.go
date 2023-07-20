package awsappstream


// The S3 location of the app block.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &S3LocationProperty{
//   	S3Bucket: jsii.String("s3Bucket"),
//
//   	// the properties below are optional
//   	S3Key: jsii.String("s3Key"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-appblock-s3location.html
//
type CfnAppBlock_S3LocationProperty struct {
	// The S3 bucket of the app block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-appblock-s3location.html#cfn-appstream-appblock-s3location-s3bucket
	//
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The S3 key of the S3 object of the virtual hard disk.
	//
	// This is required when it's used by `SetupScriptDetails` and `PostSetupScriptDetails` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-appblock-s3location.html#cfn-appstream-appblock-s3location-s3key
	//
	S3Key *string `field:"optional" json:"s3Key" yaml:"s3Key"`
}

