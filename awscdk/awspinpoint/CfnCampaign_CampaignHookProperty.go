package awspinpoint


// Specifies settings for invoking an Lambda function that customizes a segment for a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   campaignHookProperty := &CampaignHookProperty{
//   	LambdaFunctionName: jsii.String("lambdaFunctionName"),
//   	Mode: jsii.String("mode"),
//   	WebUrl: jsii.String("webUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignhook.html
//
type CfnCampaign_CampaignHookProperty struct {
	// The name or Amazon Resource Name (ARN) of the Lambda function that Amazon Pinpoint invokes to customize a segment for a campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignhook.html#cfn-pinpoint-campaign-campaignhook-lambdafunctionname
	//
	LambdaFunctionName *string `field:"optional" json:"lambdaFunctionName" yaml:"lambdaFunctionName"`
	// The mode that Amazon Pinpoint uses to invoke the Lambda function. Possible values are:.
	//
	// - `FILTER` - Invoke the function to customize the segment that's used by a campaign.
	// - `DELIVERY` - (Deprecated) Previously, invoked the function to send a campaign through a custom channel. This functionality is not supported anymore. To send a campaign through a custom channel, use the `CustomDeliveryConfiguration` and `CampaignCustomMessage` objects of the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignhook.html#cfn-pinpoint-campaign-campaignhook-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The web URL that Amazon Pinpoint calls to invoke the Lambda function over HTTPS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignhook.html#cfn-pinpoint-campaign-campaignhook-weburl
	//
	WebUrl *string `field:"optional" json:"webUrl" yaml:"webUrl"`
}

