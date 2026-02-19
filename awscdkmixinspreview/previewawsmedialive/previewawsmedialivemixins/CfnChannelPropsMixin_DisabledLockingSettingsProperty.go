package previewawsmedialivemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   disabledLockingSettingsProperty := &DisabledLockingSettingsProperty{
//   	CustomEpoch: jsii.String("customEpoch"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-disabledlockingsettings.html
//
type CfnChannelPropsMixin_DisabledLockingSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-disabledlockingsettings.html#cfn-medialive-channel-disabledlockingsettings-customepoch
	//
	CustomEpoch *string `field:"optional" json:"customEpoch" yaml:"customEpoch"`
}

