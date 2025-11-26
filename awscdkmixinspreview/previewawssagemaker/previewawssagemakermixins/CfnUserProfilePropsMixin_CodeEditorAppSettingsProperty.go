package previewawssagemakermixins


// The Code Editor application settings.
//
// For more information about Code Editor, see [Get started with Code Editor in Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/code-editor.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeEditorAppSettingsProperty := &CodeEditorAppSettingsProperty{
//   	AppLifecycleManagement: &AppLifecycleManagementProperty{
//   		IdleSettings: &IdleSettingsProperty{
//   			IdleTimeoutInMinutes: jsii.Number(123),
//   			LifecycleManagement: jsii.String("lifecycleManagement"),
//   			MaxIdleTimeoutInMinutes: jsii.Number(123),
//   			MinIdleTimeoutInMinutes: jsii.Number(123),
//   		},
//   	},
//   	BuiltInLifecycleConfigArn: jsii.String("builtInLifecycleConfigArn"),
//   	CustomImages: []interface{}{
//   		&CustomImageProperty{
//   			AppImageConfigName: jsii.String("appImageConfigName"),
//   			ImageName: jsii.String("imageName"),
//   			ImageVersionNumber: jsii.Number(123),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-codeeditorappsettings.html
//
type CfnUserProfilePropsMixin_CodeEditorAppSettingsProperty struct {
	// Settings that are used to configure and manage the lifecycle of CodeEditor applications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-codeeditorappsettings.html#cfn-sagemaker-userprofile-codeeditorappsettings-applifecyclemanagement
	//
	AppLifecycleManagement interface{} `field:"optional" json:"appLifecycleManagement" yaml:"appLifecycleManagement"`
	// The lifecycle configuration that runs before the default lifecycle configuration.
	//
	// It can override changes made in the default lifecycle configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-codeeditorappsettings.html#cfn-sagemaker-userprofile-codeeditorappsettings-builtinlifecycleconfigarn
	//
	BuiltInLifecycleConfigArn *string `field:"optional" json:"builtInLifecycleConfigArn" yaml:"builtInLifecycleConfigArn"`
	// A list of custom SageMaker images that are configured to run as a Code Editor app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-codeeditorappsettings.html#cfn-sagemaker-userprofile-codeeditorappsettings-customimages
	//
	CustomImages interface{} `field:"optional" json:"customImages" yaml:"customImages"`
	// The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the Code Editor app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-codeeditorappsettings.html#cfn-sagemaker-userprofile-codeeditorappsettings-defaultresourcespec
	//
	DefaultResourceSpec interface{} `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// The Amazon Resource Name (ARN) of the Code Editor application lifecycle configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-codeeditorappsettings.html#cfn-sagemaker-userprofile-codeeditorappsettings-lifecycleconfigarns
	//
	LifecycleConfigArns *[]*string `field:"optional" json:"lifecycleConfigArns" yaml:"lifecycleConfigArns"`
}

