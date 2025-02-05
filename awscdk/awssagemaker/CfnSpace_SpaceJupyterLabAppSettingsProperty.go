package awssagemaker


// The settings for the JupyterLab application within a space.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spaceJupyterLabAppSettingsProperty := &SpaceJupyterLabAppSettingsProperty{
//   	AppLifecycleManagement: &SpaceAppLifecycleManagementProperty{
//   		IdleSettings: &SpaceIdleSettingsProperty{
//   			IdleTimeoutInMinutes: jsii.Number(123),
//   		},
//   	},
//   	CodeRepositories: []interface{}{
//   		&CodeRepositoryProperty{
//   			RepositoryUrl: jsii.String("repositoryUrl"),
//   		},
//   	},
//   	DefaultResourceSpec: &ResourceSpecProperty{
//   		InstanceType: jsii.String("instanceType"),
//   		LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   		SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   		SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacejupyterlabappsettings.html
//
type CfnSpace_SpaceJupyterLabAppSettingsProperty struct {
	// Settings that are used to configure and manage the lifecycle of JupyterLab applications in a space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacejupyterlabappsettings.html#cfn-sagemaker-space-spacejupyterlabappsettings-applifecyclemanagement
	//
	AppLifecycleManagement interface{} `field:"optional" json:"appLifecycleManagement" yaml:"appLifecycleManagement"`
	// A list of Git repositories that SageMaker automatically displays to users for cloning in the JupyterLab application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacejupyterlabappsettings.html#cfn-sagemaker-space-spacejupyterlabappsettings-coderepositories
	//
	CodeRepositories interface{} `field:"optional" json:"codeRepositories" yaml:"codeRepositories"`
	// Specifies the ARNs of a SageMaker image and SageMaker image version, and the instance type that the version runs on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacejupyterlabappsettings.html#cfn-sagemaker-space-spacejupyterlabappsettings-defaultresourcespec
	//
	DefaultResourceSpec interface{} `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
}

