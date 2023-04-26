package awssagemaker


// The configuration for the file system and kernels in a SageMaker image running as a KernelGateway app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kernelGatewayImageConfigProperty := &KernelGatewayImageConfigProperty{
//   	KernelSpecs: []interface{}{
//   		&KernelSpecProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			DisplayName: jsii.String("displayName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	FileSystemConfig: &FileSystemConfigProperty{
//   		DefaultGid: jsii.Number(123),
//   		DefaultUid: jsii.Number(123),
//   		MountPath: jsii.String("mountPath"),
//   	},
//   }
//
type CfnAppImageConfig_KernelGatewayImageConfigProperty struct {
	// The specification of the Jupyter kernels in the image.
	KernelSpecs interface{} `field:"required" json:"kernelSpecs" yaml:"kernelSpecs"`
	// The Amazon Elastic File System (EFS) storage configuration for a SageMaker image.
	FileSystemConfig interface{} `field:"optional" json:"fileSystemConfig" yaml:"fileSystemConfig"`
}

