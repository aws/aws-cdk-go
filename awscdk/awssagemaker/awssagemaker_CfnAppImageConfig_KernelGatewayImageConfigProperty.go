package awssagemaker


// The configuration for the file system and kernels in a SageMaker image running as a KernelGateway app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kernelGatewayImageConfigProperty := &kernelGatewayImageConfigProperty{
//   	kernelSpecs: []interface{}{
//   		&kernelSpecProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			displayName: jsii.String("displayName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	fileSystemConfig: &fileSystemConfigProperty{
//   		defaultGid: jsii.Number(123),
//   		defaultUid: jsii.Number(123),
//   		mountPath: jsii.String("mountPath"),
//   	},
//   }
//
type CfnAppImageConfig_KernelGatewayImageConfigProperty struct {
	// The specification of the Jupyter kernels in the image.
	KernelSpecs interface{} `field:"required" json:"kernelSpecs" yaml:"kernelSpecs"`
	// The Amazon Elastic File System (EFS) storage configuration for a SageMaker image.
	FileSystemConfig interface{} `field:"optional" json:"fileSystemConfig" yaml:"fileSystemConfig"`
}

