package awscodepipeline


// Properties for defining a `CfnWebhook`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWebhookProps := &CfnWebhookProps{
//   	Authentication: jsii.String("authentication"),
//   	AuthenticationConfiguration: &WebhookAuthConfigurationProperty{
//   		AllowedIpRange: jsii.String("allowedIpRange"),
//   		SecretToken: jsii.String("secretToken"),
//   	},
//   	Filters: []interface{}{
//   		&WebhookFilterRuleProperty{
//   			JsonPath: jsii.String("jsonPath"),
//
//   			// the properties below are optional
//   			MatchEquals: jsii.String("matchEquals"),
//   		},
//   	},
//   	TargetAction: jsii.String("targetAction"),
//   	TargetPipeline: jsii.String("targetPipeline"),
//   	TargetPipelineVersion: jsii.Number(123),
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   	RegisterWithThirdParty: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html
//
type CfnWebhookProps struct {
	// Supported options are GITHUB_HMAC, IP, and UNAUTHENTICATED.
	//
	// - For information about the authentication scheme implemented by GITHUB_HMAC, see [Securing your webhooks](https://docs.aws.amazon.com/https://developer.github.com/webhooks/securing/) on the GitHub Developer website.
	// - IP rejects webhooks trigger requests unless they originate from an IP address in the IP range whitelisted in the authentication configuration.
	// - UNAUTHENTICATED accepts all webhook trigger requests regardless of origin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-authentication
	//
	Authentication *string `field:"required" json:"authentication" yaml:"authentication"`
	// Properties that configure the authentication applied to incoming webhook trigger requests.
	//
	// The required properties depend on the authentication type. For GITHUB_HMAC, only the `SecretToken` property must be set. For IP, only the `AllowedIPRange` property must be set to a valid CIDR range. For UNAUTHENTICATED, no properties can be set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-authenticationconfiguration
	//
	AuthenticationConfiguration interface{} `field:"required" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// A list of rules applied to the body/payload sent in the POST request to a webhook URL.
	//
	// All defined rules must pass for the request to be accepted and the pipeline started.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-filters
	//
	Filters interface{} `field:"required" json:"filters" yaml:"filters"`
	// The name of the action in a pipeline you want to connect to the webhook.
	//
	// The action must be from the source (first) stage of the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-targetaction
	//
	TargetAction *string `field:"required" json:"targetAction" yaml:"targetAction"`
	// The name of the pipeline you want to connect to the webhook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-targetpipeline
	//
	TargetPipeline *string `field:"required" json:"targetPipeline" yaml:"targetPipeline"`
	// The version number of the pipeline to be connected to the trigger request.
	//
	// Required: Yes
	//
	// Type: Integer
	//
	// Update requires: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-targetpipelineversion
	//
	TargetPipelineVersion *float64 `field:"required" json:"targetPipelineVersion" yaml:"targetPipelineVersion"`
	// The name of the webhook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Configures a connection between the webhook that was created and the external tool with events to be detected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-registerwiththirdparty
	//
	RegisterWithThirdParty interface{} `field:"optional" json:"registerWithThirdParty" yaml:"registerWithThirdParty"`
}

