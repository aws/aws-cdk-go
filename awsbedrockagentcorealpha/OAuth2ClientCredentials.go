package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// OAuth2 client identifier and secret registered with the identity provider (all vendors).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var secretValue SecretValue
//
//   oAuth2ClientCredentials := &OAuth2ClientCredentials{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: secretValue,
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type OAuth2ClientCredentials struct {
	// OAuth2 client identifier.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// OAuth2 client secret.
	//
	// **NOTE:** The client secret will be included in the CloudFormation template as part of synthesis.
	// The service stores the secret in Secrets Manager after creation, but the value is visible
	// in the template and deployment history. Use `SecretValue.unsafePlainText()` to explicitly
	// acknowledge plaintext, or pass a reference from another construct to avoid embedding the
	// literal value.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ClientSecret awscdk.SecretValue `field:"required" json:"clientSecret" yaml:"clientSecret"`
}

