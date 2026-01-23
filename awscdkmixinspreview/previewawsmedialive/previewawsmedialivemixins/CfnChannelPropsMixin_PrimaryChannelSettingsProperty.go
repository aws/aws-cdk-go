package previewawsmedialivemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   primaryChannelSettingsProperty := &PrimaryChannelSettingsProperty{
//   	LinkedChannelType: jsii.String("linkedChannelType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-primarychannelsettings.html
//
type CfnChannelPropsMixin_PrimaryChannelSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-primarychannelsettings.html#cfn-medialive-channel-primarychannelsettings-linkedchanneltype
	//
	LinkedChannelType *string `field:"optional" json:"linkedChannelType" yaml:"linkedChannelType"`
}

