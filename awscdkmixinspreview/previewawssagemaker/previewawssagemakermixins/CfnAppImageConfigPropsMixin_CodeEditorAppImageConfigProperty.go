package previewawssagemakermixins


// The configuration for the file system and kernels in a SageMaker image running as a Code Editor app.
//
// The `FileSystemConfig` object is not supported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeEditorAppImageConfigProperty := &CodeEditorAppImageConfigProperty{
//   	ContainerConfig: &ContainerConfigProperty{
//   		ContainerArguments: []*string{
//   			jsii.String("containerArguments"),
//   		},
//   		ContainerEntrypoint: []*string{
//   			jsii.String("containerEntrypoint"),
//   		},
//   		ContainerEnvironmentVariables: []interface{}{
//   			&CustomImageContainerEnvironmentVariableProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-codeeditorappimageconfig.html
//
type CfnAppImageConfigPropsMixin_CodeEditorAppImageConfigProperty struct {
	// The container configuration for a SageMaker image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-codeeditorappimageconfig.html#cfn-sagemaker-appimageconfig-codeeditorappimageconfig-containerconfig
	//
	ContainerConfig interface{} `field:"optional" json:"containerConfig" yaml:"containerConfig"`
}

