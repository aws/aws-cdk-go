package mixinsawsconnectcampaignsv2


// The configuration of timezone for recipient.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   localTimeZoneConfigProperty := &LocalTimeZoneConfigProperty{
//   	DefaultTimeZone: jsii.String("defaultTimeZone"),
//   	LocalTimeZoneDetection: []*string{
//   		jsii.String("localTimeZoneDetection"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-localtimezoneconfig.html
//
type CfnCampaignPropsMixin_LocalTimeZoneConfigProperty struct {
	// The timezone to use for all recipients.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-localtimezoneconfig.html#cfn-connectcampaignsv2-campaign-localtimezoneconfig-defaulttimezone
	//
	DefaultTimeZone *string `field:"optional" json:"defaultTimeZone" yaml:"defaultTimeZone"`
	// Detects methods for the recipient's timezone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-localtimezoneconfig.html#cfn-connectcampaignsv2-campaign-localtimezoneconfig-localtimezonedetection
	//
	LocalTimeZoneDetection *[]*string `field:"optional" json:"localTimeZoneDetection" yaml:"localTimeZoneDetection"`
}

