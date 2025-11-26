package previewawssagemakermixins


// Settings related to idle shutdown of Studio applications in a space.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   spaceIdleSettingsProperty := &SpaceIdleSettingsProperty{
//   	IdleTimeoutInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spaceidlesettings.html
//
type CfnSpacePropsMixin_SpaceIdleSettingsProperty struct {
	// The time that SageMaker waits after the application becomes idle before shutting it down.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spaceidlesettings.html#cfn-sagemaker-space-spaceidlesettings-idletimeoutinminutes
	//
	IdleTimeoutInMinutes *float64 `field:"optional" json:"idleTimeoutInMinutes" yaml:"idleTimeoutInMinutes"`
}

