package awssam


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
	// `CfnFunction.FileSystemConfigProperty.Arn`.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// `CfnFunction.FileSystemConfigProperty.LocalMountPath`.
	LocalMountPath *string `field:"optional" json:"localMountPath" yaml:"localMountPath"`
}

