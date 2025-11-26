package previewawsconnectcampaignsv2mixins


// Contains preview outbound mode timeout configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   timeoutConfigProperty := &TimeoutConfigProperty{
//   	DurationInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-timeoutconfig.html
//
type CfnCampaignPropsMixin_TimeoutConfigProperty struct {
	// Duration in seconds for the countdown timer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-timeoutconfig.html#cfn-connectcampaignsv2-campaign-timeoutconfig-durationinseconds
	//
	DurationInSeconds *float64 `field:"optional" json:"durationInSeconds" yaml:"durationInSeconds"`
}

