package awssagemaker


// The settings for the JupyterLab application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jupyterLabAppSettingsProperty := &JupyterLabAppSettingsProperty{
//   	CodeRepositories: []interface{}{
//   		&CodeRepositoryProperty{
//   			RepositoryUrl: jsii.String("repositoryUrl"),
//   		},
//   	},
//   	CustomImages: []interface{}{
//   		&CustomImageProperty{
//   			AppImageConfigName: jsii.String("appImageConfigName"),
//   			ImageName: jsii.String("imageName"),
//
//   			// the properties below are optional
//   			ImageVersionNumber: jsii.Number(123),
//   		},
//   	},
//   	DefaultResourceSpec: &ResourceSpecProperty{
//   		InstanceType: jsii.String("instanceType"),
//   		SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   		SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   	},
//   	LifecycleConfigArns: []*string{
//   		jsii.String("lifecycleConfigArns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-jupyterlabappsettings.html
//
type CfnUserProfile_JupyterLabAppSettingsProperty struct {
	// A list of Git repositories that SageMaker automatically displays to users for cloning in the JupyterLab application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-jupyterlabappsettings.html#cfn-sagemaker-userprofile-jupyterlabappsettings-coderepositories
	//
	CodeRepositories interface{} `field:"optional" json:"codeRepositories" yaml:"codeRepositories"`
	// A list of custom SageMaker images that are configured to run as a JupyterLab app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-jupyterlabappsettings.html#cfn-sagemaker-userprofile-jupyterlabappsettings-customimages
	//
	CustomImages interface{} `field:"optional" json:"customImages" yaml:"customImages"`
	// The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the JupyterLab app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-jupyterlabappsettings.html#cfn-sagemaker-userprofile-jupyterlabappsettings-defaultresourcespec
	//
	DefaultResourceSpec interface{} `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// The Amazon Resource Name (ARN) of the lifecycle configurations attached to the user profile or domain.
	//
	// To remove a lifecycle config, you must set `LifecycleConfigArns` to an empty list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-jupyterlabappsettings.html#cfn-sagemaker-userprofile-jupyterlabappsettings-lifecycleconfigarns
	//
	LifecycleConfigArns *[]*string `field:"optional" json:"lifecycleConfigArns" yaml:"lifecycleConfigArns"`
}

