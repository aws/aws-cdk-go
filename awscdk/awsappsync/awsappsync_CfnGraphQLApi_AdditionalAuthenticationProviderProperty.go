package awsappsync


// Describes an additional authentication provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   additionalAuthenticationProviderProperty := &AdditionalAuthenticationProviderProperty{
//   	AuthenticationType: jsii.String("authenticationType"),
//
//   	// the properties below are optional
//   	LambdaAuthorizerConfig: &LambdaAuthorizerConfigProperty{
//   		AuthorizerResultTtlInSeconds: jsii.Number(123),
//   		AuthorizerUri: jsii.String("authorizerUri"),
//   		IdentityValidationExpression: jsii.String("identityValidationExpression"),
//   	},
//   	OpenIdConnectConfig: &OpenIDConnectConfigProperty{
//   		AuthTtl: jsii.Number(123),
//   		ClientId: jsii.String("clientId"),
//   		IatTtl: jsii.Number(123),
//   		Issuer: jsii.String("issuer"),
//   	},
//   	UserPoolConfig: &CognitoUserPoolConfigProperty{
//   		AppIdClientRegex: jsii.String("appIdClientRegex"),
//   		AwsRegion: jsii.String("awsRegion"),
//   		UserPoolId: jsii.String("userPoolId"),
//   	},
//   }
//
type CfnGraphQLApi_AdditionalAuthenticationProviderProperty struct {
	// The authentication type for API key, AWS Identity and Access Management , OIDC, Amazon Cognito user pools , or AWS Lambda .
	//
	// Valid Values: `API_KEY` | `AWS_IAM` | `OPENID_CONNECT` | `AMAZON_COGNITO_USER_POOLS` | `AWS_LAMBDA`.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// Configuration for AWS Lambda function authorization.
	LambdaAuthorizerConfig interface{} `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// The OIDC configuration.
	OpenIdConnectConfig interface{} `field:"optional" json:"openIdConnectConfig" yaml:"openIdConnectConfig"`
	// The Amazon Cognito user pool configuration.
	UserPoolConfig interface{} `field:"optional" json:"userPoolConfig" yaml:"userPoolConfig"`
}

