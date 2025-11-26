package previewawssagemakermixins


// The JupyterServer app settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   jupyterServerAppSettingsProperty := &JupyterServerAppSettingsProperty{
//   	DefaultResourceSpec: &ResourceSpecProperty{
//   		InstanceType: jsii.String("instanceType"),
//   		LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   		SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   		SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   	},
//   	LifecycleConfigArns: []*string{
//   		jsii.String("lifecycleConfigArns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-jupyterserverappsettings.html
//
type CfnDomainPropsMixin_JupyterServerAppSettingsProperty struct {
	// The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the JupyterServer app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-jupyterserverappsettings.html#cfn-sagemaker-domain-jupyterserverappsettings-defaultresourcespec
	//
	DefaultResourceSpec interface{} `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// The Amazon Resource Name (ARN) of the Lifecycle Configurations attached to the JupyterServerApp.
	//
	// If you use this parameter, the `DefaultResourceSpec` parameter is also required.
	//
	// > To remove a Lifecycle Config, you must set `LifecycleConfigArns` to an empty list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-jupyterserverappsettings.html#cfn-sagemaker-domain-jupyterserverappsettings-lifecycleconfigarns
	//
	LifecycleConfigArns *[]*string `field:"optional" json:"lifecycleConfigArns" yaml:"lifecycleConfigArns"`
}

