package awssagemaker


// The CodeEditor app settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeEditorAppSettingsProperty := &CodeEditorAppSettingsProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-codeeditorappsettings.html
//
type CfnDomain_CodeEditorAppSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-codeeditorappsettings.html#cfn-sagemaker-domain-codeeditorappsettings-defaultresourcespec
	//
	DefaultResourceSpec interface{} `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// A list of LifecycleConfigArns available for use with CodeEditor apps.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-codeeditorappsettings.html#cfn-sagemaker-domain-codeeditorappsettings-lifecycleconfigarns
	//
	LifecycleConfigArns *[]*string `field:"optional" json:"lifecycleConfigArns" yaml:"lifecycleConfigArns"`
}

