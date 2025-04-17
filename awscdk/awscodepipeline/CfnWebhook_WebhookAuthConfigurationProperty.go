package awscodepipeline


// The authentication applied to incoming webhook trigger requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webhookAuthConfigurationProperty := &WebhookAuthConfigurationProperty{
//   	AllowedIpRange: jsii.String("allowedIpRange"),
//   	SecretToken: jsii.String("secretToken"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-webhook-webhookauthconfiguration.html
//
type CfnWebhook_WebhookAuthConfigurationProperty struct {
	// The property used to configure acceptance of webhooks in an IP address range.
	//
	// For IP, only the `AllowedIPRange` property must be set. This property must be set to a valid CIDR range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-webhook-webhookauthconfiguration.html#cfn-codepipeline-webhook-webhookauthconfiguration-allowediprange
	//
	AllowedIpRange *string `field:"optional" json:"allowedIpRange" yaml:"allowedIpRange"`
	// The property used to configure GitHub authentication. For GITHUB_HMAC, only the `SecretToken` property must be set.
	//
	// > When creating CodePipeline webhooks, do not use your own credentials or reuse the same secret token across multiple webhooks. For optimal security, generate a unique secret token for each webhook you create. The secret token is an arbitrary string that you provide, which GitHub uses to compute and sign the webhook payloads sent to CodePipeline, for protecting the integrity and authenticity of the webhook payloads. Using your own credentials or reusing the same token across multiple webhooks can lead to security vulnerabilities. > If a secret token was provided, it will be redacted in the response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-webhook-webhookauthconfiguration.html#cfn-codepipeline-webhook-webhookauthconfiguration-secrettoken
	//
	SecretToken *string `field:"optional" json:"secretToken" yaml:"secretToken"`
}

