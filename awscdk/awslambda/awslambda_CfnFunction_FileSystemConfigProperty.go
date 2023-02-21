package awslambda


// Details about the connection between a Lambda function and an [Amazon EFS file system](https://docs.aws.amazon.com/lambda/latest/dg/configuration-filesystem.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileSystemConfigProperty := &fileSystemConfigProperty{
//   	arn: jsii.String("arn"),
//   	localMountPath: jsii.String("localMountPath"),
//   }
//
type CfnFunction_FileSystemConfigProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon EFS access point that provides access to the file system.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The path where the function can access the file system, starting with `/mnt/` .
	LocalMountPath *string `field:"required" json:"localMountPath" yaml:"localMountPath"`
}

