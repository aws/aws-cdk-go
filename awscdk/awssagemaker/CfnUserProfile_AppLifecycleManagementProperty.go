package awssagemaker


// Settings that are used to configure and manage the lifecycle of Amazon SageMaker Studio applications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appLifecycleManagementProperty := &AppLifecycleManagementProperty{
//   	IdleSettings: &IdleSettingsProperty{
//   		IdleTimeoutInMinutes: jsii.Number(123),
//   		LifecycleManagement: jsii.String("lifecycleManagement"),
//   		MaxIdleTimeoutInMinutes: jsii.Number(123),
//   		MinIdleTimeoutInMinutes: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-applifecyclemanagement.html
//
type CfnUserProfile_AppLifecycleManagementProperty struct {
	// Settings related to idle shutdown of Studio applications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-applifecyclemanagement.html#cfn-sagemaker-userprofile-applifecyclemanagement-idlesettings
	//
	IdleSettings interface{} `field:"optional" json:"idleSettings" yaml:"idleSettings"`
}

