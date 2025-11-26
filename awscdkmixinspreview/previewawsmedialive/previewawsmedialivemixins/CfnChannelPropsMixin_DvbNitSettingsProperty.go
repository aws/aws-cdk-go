package previewawsmedialivemixins


// The configuration of DVB NIT.
//
// The parent of this entity is M2tsSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dvbNitSettingsProperty := &DvbNitSettingsProperty{
//   	NetworkId: jsii.Number(123),
//   	NetworkName: jsii.String("networkName"),
//   	RepInterval: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbnitsettings.html
//
type CfnChannelPropsMixin_DvbNitSettingsProperty struct {
	// The numeric value placed in the Network Information Table (NIT).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbnitsettings.html#cfn-medialive-channel-dvbnitsettings-networkid
	//
	NetworkId *float64 `field:"optional" json:"networkId" yaml:"networkId"`
	// The network name text placed in the networkNameDescriptor inside the Network Information Table (NIT).
	//
	// The maximum length is 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbnitsettings.html#cfn-medialive-channel-dvbnitsettings-networkname
	//
	NetworkName *string `field:"optional" json:"networkName" yaml:"networkName"`
	// The number of milliseconds between instances of this table in the output transport stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbnitsettings.html#cfn-medialive-channel-dvbnitsettings-repinterval
	//
	RepInterval *float64 `field:"optional" json:"repInterval" yaml:"repInterval"`
}

