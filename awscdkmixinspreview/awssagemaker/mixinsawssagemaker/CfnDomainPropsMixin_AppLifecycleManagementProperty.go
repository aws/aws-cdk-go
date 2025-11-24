package mixinsawssagemaker


// Settings that are used to configure and manage the lifecycle of Amazon SageMaker Studio applications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-applifecyclemanagement.html
//
type CfnDomainPropsMixin_AppLifecycleManagementProperty struct {
	// Settings related to idle shutdown of Studio applications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-applifecyclemanagement.html#cfn-sagemaker-domain-applifecyclemanagement-idlesettings
	//
	IdleSettings interface{} `field:"optional" json:"idleSettings" yaml:"idleSettings"`
}

