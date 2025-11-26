package previewawssagemakermixins


// The configuration for the file system and kernels in a SageMaker AI image running as a KernelGateway app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kernelGatewayImageConfigProperty := &KernelGatewayImageConfigProperty{
//   	FileSystemConfig: &FileSystemConfigProperty{
//   		DefaultGid: jsii.Number(123),
//   		DefaultUid: jsii.Number(123),
//   		MountPath: jsii.String("mountPath"),
//   	},
//   	KernelSpecs: []interface{}{
//   		&KernelSpecProperty{
//   			DisplayName: jsii.String("displayName"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-kernelgatewayimageconfig.html
//
type CfnAppImageConfigPropsMixin_KernelGatewayImageConfigProperty struct {
	// The Amazon Elastic File System storage configuration for a SageMaker AI image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-kernelgatewayimageconfig.html#cfn-sagemaker-appimageconfig-kernelgatewayimageconfig-filesystemconfig
	//
	FileSystemConfig interface{} `field:"optional" json:"fileSystemConfig" yaml:"fileSystemConfig"`
	// The specification of the Jupyter kernels in the image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-kernelgatewayimageconfig.html#cfn-sagemaker-appimageconfig-kernelgatewayimageconfig-kernelspecs
	//
	KernelSpecs interface{} `field:"optional" json:"kernelSpecs" yaml:"kernelSpecs"`
}

