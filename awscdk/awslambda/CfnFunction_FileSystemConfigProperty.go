package awslambda


// Details about the connection between a Lambda function and an [Amazon EFS file system](https://docs.aws.amazon.com/lambda/latest/dg/configuration-filesystem.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileSystemConfigProperty := &FileSystemConfigProperty{
//   	Arn: jsii.String("arn"),
//   	LocalMountPath: jsii.String("localMountPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-filesystemconfig.html
//
type CfnFunction_FileSystemConfigProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon EFS access point that provides access to the file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-filesystemconfig.html#cfn-lambda-function-filesystemconfig-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The path where the function can access the file system, starting with `/mnt/` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-filesystemconfig.html#cfn-lambda-function-filesystemconfig-localmountpath
	//
	LocalMountPath *string `field:"required" json:"localMountPath" yaml:"localMountPath"`
}

