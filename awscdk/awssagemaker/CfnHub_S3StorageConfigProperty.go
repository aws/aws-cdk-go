package awssagemaker


// The Amazon S3 storage configuration for the hub.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3StorageConfigProperty := &S3StorageConfigProperty{
//   	S3OutputPath: jsii.String("s3OutputPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-hub-s3storageconfig.html
//
type CfnHub_S3StorageConfigProperty struct {
	// The Amazon S3 bucket prefix for hosting hub content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-hub-s3storageconfig.html#cfn-sagemaker-hub-s3storageconfig-s3outputpath
	//
	S3OutputPath *string `field:"optional" json:"s3OutputPath" yaml:"s3OutputPath"`
}

