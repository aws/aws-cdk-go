package previewawsmedialivemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cmafIngestOutputSettingsProperty := &CmafIngestOutputSettingsProperty{
//   	NameModifier: jsii.String("nameModifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestoutputsettings.html
//
type CfnChannelPropsMixin_CmafIngestOutputSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestoutputsettings.html#cfn-medialive-channel-cmafingestoutputsettings-namemodifier
	//
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}

