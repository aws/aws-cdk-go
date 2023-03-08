package awspinpoint


// Specifies settings for invoking an Lambda function that customizes a segment for a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   campaignHookProperty := &campaignHookProperty{
//   	lambdaFunctionName: jsii.String("lambdaFunctionName"),
//   	mode: jsii.String("mode"),
//   	webUrl: jsii.String("webUrl"),
//   }
//
type CfnCampaign_CampaignHookProperty struct {
	// The name or Amazon Resource Name (ARN) of the Lambda function that Amazon Pinpoint invokes to customize a segment for a campaign.
	LambdaFunctionName *string `field:"optional" json:"lambdaFunctionName" yaml:"lambdaFunctionName"`
	// The mode that Amazon Pinpoint uses to invoke the Lambda function. Possible values are:.
	//
	// - `FILTER` - Invoke the function to customize the segment that's used by a campaign.
	// - `DELIVERY` - (Deprecated) Previously, invoked the function to send a campaign through a custom channel. This functionality is not supported anymore. To send a campaign through a custom channel, use the `CustomDeliveryConfiguration` and `CampaignCustomMessage` objects of the campaign.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The web URL that Amazon Pinpoint calls to invoke the Lambda function over HTTPS.
	WebUrl *string `field:"optional" json:"webUrl" yaml:"webUrl"`
}

