package awspinpoint


// Properties for defining a `CfnApplicationSettings`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationSettingsProps := &cfnApplicationSettingsProps{
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	campaignHook: &campaignHookProperty{
//   		lambdaFunctionName: jsii.String("lambdaFunctionName"),
//   		mode: jsii.String("mode"),
//   		webUrl: jsii.String("webUrl"),
//   	},
//   	cloudWatchMetricsEnabled: jsii.Boolean(false),
//   	limits: &limitsProperty{
//   		daily: jsii.Number(123),
//   		maximumDuration: jsii.Number(123),
//   		messagesPerSecond: jsii.Number(123),
//   		total: jsii.Number(123),
//   	},
//   	quietTime: &quietTimeProperty{
//   		end: jsii.String("end"),
//   		start: jsii.String("start"),
//   	},
//   }
//
type CfnApplicationSettingsProps struct {
	// The unique identifier for the Amazon Pinpoint application.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The settings for the Lambda function to use by default as a code hook for campaigns in the application.
	//
	// To override these settings for a specific campaign, use the Campaign resource to define custom Lambda function settings for the campaign.
	CampaignHook interface{} `field:"optional" json:"campaignHook" yaml:"campaignHook"`
	// Specifies whether to enable application-related alarms in Amazon CloudWatch.
	CloudWatchMetricsEnabled interface{} `field:"optional" json:"cloudWatchMetricsEnabled" yaml:"cloudWatchMetricsEnabled"`
	// The default sending limits for campaigns in the application.
	//
	// To override these limits for a specific campaign, use the Campaign resource to define custom limits for the campaign.
	Limits interface{} `field:"optional" json:"limits" yaml:"limits"`
	// The default quiet time for campaigns in the application.
	//
	// Quiet time is a specific time range when campaigns don't send messages to endpoints, if all the following conditions are met:
	//
	// - The `EndpointDemographic.Timezone` property of the endpoint is set to a valid value.
	//
	// - The current time in the endpoint's time zone is later than or equal to the time specified by the `QuietTime.Start` property for the application (or a campaign that has custom quiet time settings).
	//
	// - The current time in the endpoint's time zone is earlier than or equal to the time specified by the `QuietTime.End` property for the application (or a campaign that has custom quiet time settings).
	//
	// If any of the preceding conditions isn't met, the endpoint will receive messages from a campaign, even if quiet time is enabled.
	//
	// To override the default quiet time settings for a specific campaign, use the Campaign resource to define a custom quiet time for the campaign.
	QuietTime interface{} `field:"optional" json:"quietTime" yaml:"quietTime"`
}

