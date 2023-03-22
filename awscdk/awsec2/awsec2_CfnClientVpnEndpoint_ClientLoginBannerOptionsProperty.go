package awsec2


// Options for enabling a customizable text banner that will be displayed on AWS provided clients when a VPN session is established.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientLoginBannerOptionsProperty := &clientLoginBannerOptionsProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	bannerText: jsii.String("bannerText"),
//   }
//
type CfnClientVpnEndpoint_ClientLoginBannerOptionsProperty struct {
	// Enable or disable a customizable text banner that will be displayed on AWS provided clients when a VPN session is established.
	//
	// Valid values: `true | false`
	//
	// Default value: `false`.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Customizable text that will be displayed in a banner on AWS provided clients when a VPN session is established.
	//
	// UTF-8 encoded characters only. Maximum of 1400 characters.
	BannerText *string `field:"optional" json:"bannerText" yaml:"bannerText"`
}

