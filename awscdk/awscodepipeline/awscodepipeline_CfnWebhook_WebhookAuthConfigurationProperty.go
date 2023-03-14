package awscodepipeline


// The authentication applied to incoming webhook trigger requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webhookAuthConfigurationProperty := &webhookAuthConfigurationProperty{
//   	allowedIpRange: jsii.String("allowedIpRange"),
//   	secretToken: jsii.String("secretToken"),
//   }
//
type CfnWebhook_WebhookAuthConfigurationProperty struct {
	// The property used to configure acceptance of webhooks in an IP address range.
	//
	// For IP, only the `AllowedIPRange` property must be set. This property must be set to a valid CIDR range.
	AllowedIpRange *string `field:"optional" json:"allowedIpRange" yaml:"allowedIpRange"`
	// The property used to configure GitHub authentication.
	//
	// For GITHUB_HMAC, only the `SecretToken` property must be set.
	SecretToken *string `field:"optional" json:"secretToken" yaml:"secretToken"`
}

