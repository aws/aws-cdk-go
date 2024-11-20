package awsconnectcampaignsv2


// Local time zone config.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnCampaign_LocalTimeZoneConfigProperty struct {
	// Time Zone Id in the IANA format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-localtimezoneconfig.html#cfn-connectcampaignsv2-campaign-localtimezoneconfig-defaulttimezone
	//
	DefaultTimeZone *string `field:"optional" json:"defaultTimeZone" yaml:"defaultTimeZone"`
	// Local TimeZone Detection method list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-localtimezoneconfig.html#cfn-connectcampaignsv2-campaign-localtimezoneconfig-localtimezonedetection
	//
	LocalTimeZoneDetection *[]*string `field:"optional" json:"localTimeZoneDetection" yaml:"localTimeZoneDetection"`
}

