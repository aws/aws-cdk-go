package previewawsmedialivemixins


// Information about how to connect to the upstream system.
//
// The parent of this entity is InputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkInputSettingsProperty := &NetworkInputSettingsProperty{
//   	HlsInputSettings: &HlsInputSettingsProperty{
//   		Bandwidth: jsii.Number(123),
//   		BufferSegments: jsii.Number(123),
//   		Retries: jsii.Number(123),
//   		RetryInterval: jsii.Number(123),
//   		Scte35Source: jsii.String("scte35Source"),
//   	},
//   	MulticastInputSettings: &MulticastInputSettingsProperty{
//   		SourceIpAddress: jsii.String("sourceIpAddress"),
//   	},
//   	ServerValidation: jsii.String("serverValidation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-networkinputsettings.html
//
type CfnChannelPropsMixin_NetworkInputSettingsProperty struct {
	// Information about how to connect to the upstream system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-networkinputsettings.html#cfn-medialive-channel-networkinputsettings-hlsinputsettings
	//
	HlsInputSettings interface{} `field:"optional" json:"hlsInputSettings" yaml:"hlsInputSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-networkinputsettings.html#cfn-medialive-channel-networkinputsettings-multicastinputsettings
	//
	MulticastInputSettings interface{} `field:"optional" json:"multicastInputSettings" yaml:"multicastInputSettings"`
	// Checks HTTPS server certificates.
	//
	// When set to checkCryptographyOnly, cryptography in the certificate is checked, but not the server's name. Certain subdomains (notably S3 buckets that use dots in the bucket name) don't strictly match the corresponding certificate's wildcard pattern and would otherwise cause the channel to error. This setting is ignored for protocols that do not use HTTPS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-networkinputsettings.html#cfn-medialive-channel-networkinputsettings-servervalidation
	//
	ServerValidation *string `field:"optional" json:"serverValidation" yaml:"serverValidation"`
}

