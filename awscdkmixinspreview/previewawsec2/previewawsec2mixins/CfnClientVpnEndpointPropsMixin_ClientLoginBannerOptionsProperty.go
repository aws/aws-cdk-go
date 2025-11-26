package previewawsec2mixins


// Options for enabling a customizable text banner that will be displayed on AWS provided clients when a VPN session is established.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clientLoginBannerOptionsProperty := &ClientLoginBannerOptionsProperty{
//   	BannerText: jsii.String("bannerText"),
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-clientloginbanneroptions.html
//
type CfnClientVpnEndpointPropsMixin_ClientLoginBannerOptionsProperty struct {
	// Customizable text that will be displayed in a banner on AWS provided clients when a VPN session is established.
	//
	// UTF-8 encoded characters only. Maximum of 1400 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-clientloginbanneroptions.html#cfn-ec2-clientvpnendpoint-clientloginbanneroptions-bannertext
	//
	BannerText *string `field:"optional" json:"bannerText" yaml:"bannerText"`
	// Enable or disable a customizable text banner that will be displayed on AWS provided clients when a VPN session is established.
	//
	// Valid values: `true | false`
	//
	// Default value: `false`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-clientloginbanneroptions.html#cfn-ec2-clientvpnendpoint-clientloginbanneroptions-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

