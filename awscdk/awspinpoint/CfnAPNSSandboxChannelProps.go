package awspinpoint


// Properties for defining a `CfnAPNSSandboxChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAPNSSandboxChannelProps := &CfnAPNSSandboxChannelProps{
//   	ApplicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	BundleId: jsii.String("bundleId"),
//   	Certificate: jsii.String("certificate"),
//   	DefaultAuthenticationMethod: jsii.String("defaultAuthenticationMethod"),
//   	Enabled: jsii.Boolean(false),
//   	PrivateKey: jsii.String("privateKey"),
//   	TeamId: jsii.String("teamId"),
//   	TokenKey: jsii.String("tokenKey"),
//   	TokenKeyId: jsii.String("tokenKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnssandboxchannel.html
//
type CfnAPNSSandboxChannelProps struct {
	// The unique identifier for the Amazon Pinpoint application that the APNs sandbox channel applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnssandboxchannel.html#cfn-pinpoint-apnssandboxchannel-applicationid
	//
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The bundle identifier that's assigned to your iOS app.
	//
	// This identifier is used for APNs tokens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnssandboxchannel.html#cfn-pinpoint-apnssandboxchannel-bundleid
	//
	BundleId *string `field:"optional" json:"bundleId" yaml:"bundleId"`
	// The APNs client certificate that you received from Apple.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with APNs by using an APNs certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnssandboxchannel.html#cfn-pinpoint-apnssandboxchannel-certificate
	//
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// The default authentication method that you want Amazon Pinpoint to use when authenticating with APNs.
	//
	// Valid options are `key` or `certificate` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnssandboxchannel.html#cfn-pinpoint-apnssandboxchannel-defaultauthenticationmethod
	//
	DefaultAuthenticationMethod *string `field:"optional" json:"defaultAuthenticationMethod" yaml:"defaultAuthenticationMethod"`
	// Specifies whether to enable the APNs Sandbox channel for the Amazon Pinpoint application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnssandboxchannel.html#cfn-pinpoint-apnssandboxchannel-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The private key for the APNs client certificate that you want Amazon Pinpoint to use to communicate with APNs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnssandboxchannel.html#cfn-pinpoint-apnssandboxchannel-privatekey
	//
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// The identifier that's assigned to your Apple Developer Account team.
	//
	// This identifier is used for APNs tokens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnssandboxchannel.html#cfn-pinpoint-apnssandboxchannel-teamid
	//
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// The authentication key to use for APNs tokens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnssandboxchannel.html#cfn-pinpoint-apnssandboxchannel-tokenkey
	//
	TokenKey *string `field:"optional" json:"tokenKey" yaml:"tokenKey"`
	// The key identifier that's assigned to your APNs signing key.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with APNs by using APNs tokens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnssandboxchannel.html#cfn-pinpoint-apnssandboxchannel-tokenkeyid
	//
	TokenKeyId *string `field:"optional" json:"tokenKeyId" yaml:"tokenKeyId"`
}

