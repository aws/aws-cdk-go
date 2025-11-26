package previewawspinpointmixins


// Properties for CfnADMChannelPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnADMChannelMixinProps := &CfnADMChannelMixinProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-admchannel.html
//
type CfnADMChannelMixinProps struct {
	// The unique identifier for the Amazon Pinpoint application that the ADM channel applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-admchannel.html#cfn-pinpoint-admchannel-applicationid
	//
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// The Client ID that you received from Amazon to send messages by using ADM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-admchannel.html#cfn-pinpoint-admchannel-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The Client Secret that you received from Amazon to send messages by using ADM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-admchannel.html#cfn-pinpoint-admchannel-clientsecret
	//
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Specifies whether to enable the ADM channel for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-admchannel.html#cfn-pinpoint-admchannel-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

