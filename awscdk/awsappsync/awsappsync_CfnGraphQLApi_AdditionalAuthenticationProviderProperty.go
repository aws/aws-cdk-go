package awsappsync


// Describes an additional authentication provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   additionalAuthenticationProviderProperty := &additionalAuthenticationProviderProperty{
//   	authenticationType: jsii.String("authenticationType"),
//
//   	// the properties below are optional
//   	lambdaAuthorizerConfig: &lambdaAuthorizerConfigProperty{
//   		authorizerResultTtlInSeconds: jsii.Number(123),
//   		authorizerUri: jsii.String("authorizerUri"),
//   		identityValidationExpression: jsii.String("identityValidationExpression"),
//   	},
//   	openIdConnectConfig: &openIDConnectConfigProperty{
//   		authTtl: jsii.Number(123),
//   		clientId: jsii.String("clientId"),
//   		iatTtl: jsii.Number(123),
//   		issuer: jsii.String("issuer"),
//   	},
//   	userPoolConfig: &cognitoUserPoolConfigProperty{
//   		appIdClientRegex: jsii.String("appIdClientRegex"),
//   		awsRegion: jsii.String("awsRegion"),
//   		userPoolId: jsii.String("userPoolId"),
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

