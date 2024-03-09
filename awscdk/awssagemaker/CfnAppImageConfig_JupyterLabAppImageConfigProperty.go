package awssagemaker


// The configuration for the file system and kernels in a SageMaker image running as a JupyterLab app.
//
// The `FileSystemConfig` object is not supported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jupyterLabAppImageConfigProperty := &JupyterLabAppImageConfigProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-jupyterlabappimageconfig.html
//
type CfnAppImageConfig_JupyterLabAppImageConfigProperty struct {
	// The configuration used to run the application image container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-jupyterlabappimageconfig.html#cfn-sagemaker-appimageconfig-jupyterlabappimageconfig-containerconfig
	//
	ContainerConfig interface{} `field:"optional" json:"containerConfig" yaml:"containerConfig"`
}

