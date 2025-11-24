package mixinsawssagemaker


// Settings that are used to configure and manage the lifecycle of Amazon SageMaker Studio applications in a space.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   spaceAppLifecycleManagementProperty := &SpaceAppLifecycleManagementProperty{
//   	IdleSettings: &SpaceIdleSettingsProperty{
//   		IdleTimeoutInMinutes: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spaceapplifecyclemanagement.html
//
type CfnSpacePropsMixin_SpaceAppLifecycleManagementProperty struct {
	// Settings related to idle shutdown of Studio applications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spaceapplifecyclemanagement.html#cfn-sagemaker-space-spaceapplifecyclemanagement-idlesettings
	//
	IdleSettings interface{} `field:"optional" json:"idleSettings" yaml:"idleSettings"`
}

