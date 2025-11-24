package mixinsawssagemaker


// Settings related to idle shutdown of Studio applications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   idleSettingsProperty := &IdleSettingsProperty{
//   	IdleTimeoutInMinutes: jsii.Number(123),
//   	LifecycleManagement: jsii.String("lifecycleManagement"),
//   	MaxIdleTimeoutInMinutes: jsii.Number(123),
//   	MinIdleTimeoutInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-idlesettings.html
//
type CfnDomainPropsMixin_IdleSettingsProperty struct {
	// The time that SageMaker waits after the application becomes idle before shutting it down.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-idlesettings.html#cfn-sagemaker-domain-idlesettings-idletimeoutinminutes
	//
	IdleTimeoutInMinutes *float64 `field:"optional" json:"idleTimeoutInMinutes" yaml:"idleTimeoutInMinutes"`
	// Indicates whether idle shutdown is activated for the application type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-idlesettings.html#cfn-sagemaker-domain-idlesettings-lifecyclemanagement
	//
	LifecycleManagement *string `field:"optional" json:"lifecycleManagement" yaml:"lifecycleManagement"`
	// The maximum value in minutes that custom idle shutdown can be set to by the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-idlesettings.html#cfn-sagemaker-domain-idlesettings-maxidletimeoutinminutes
	//
	MaxIdleTimeoutInMinutes *float64 `field:"optional" json:"maxIdleTimeoutInMinutes" yaml:"maxIdleTimeoutInMinutes"`
	// The minimum value in minutes that custom idle shutdown can be set to by the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-idlesettings.html#cfn-sagemaker-domain-idlesettings-minidletimeoutinminutes
	//
	MinIdleTimeoutInMinutes *float64 `field:"optional" json:"minIdleTimeoutInMinutes" yaml:"minIdleTimeoutInMinutes"`
}

