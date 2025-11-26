package previewawsmedialivemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   esamProperty := &EsamProperty{
//   	AcquisitionPointId: jsii.String("acquisitionPointId"),
//   	AdAvailOffset: jsii.Number(123),
//   	PasswordParam: jsii.String("passwordParam"),
//   	PoisEndpoint: jsii.String("poisEndpoint"),
//   	Username: jsii.String("username"),
//   	ZoneIdentity: jsii.String("zoneIdentity"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-esam.html
//
type CfnChannelPropsMixin_EsamProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-esam.html#cfn-medialive-channel-esam-acquisitionpointid
	//
	AcquisitionPointId *string `field:"optional" json:"acquisitionPointId" yaml:"acquisitionPointId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-esam.html#cfn-medialive-channel-esam-adavailoffset
	//
	AdAvailOffset *float64 `field:"optional" json:"adAvailOffset" yaml:"adAvailOffset"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-esam.html#cfn-medialive-channel-esam-passwordparam
	//
	PasswordParam *string `field:"optional" json:"passwordParam" yaml:"passwordParam"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-esam.html#cfn-medialive-channel-esam-poisendpoint
	//
	PoisEndpoint *string `field:"optional" json:"poisEndpoint" yaml:"poisEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-esam.html#cfn-medialive-channel-esam-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-esam.html#cfn-medialive-channel-esam-zoneidentity
	//
	ZoneIdentity *string `field:"optional" json:"zoneIdentity" yaml:"zoneIdentity"`
}

