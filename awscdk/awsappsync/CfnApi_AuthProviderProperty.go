package awsappsync


// An auth provider for the AppSync API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authProviderProperty := &AuthProviderProperty{
//   	AuthType: jsii.String("authType"),
//
//   	// the properties below are optional
//   	CognitoConfig: &CognitoConfigProperty{
//   		AwsRegion: jsii.String("awsRegion"),
//   		UserPoolId: jsii.String("userPoolId"),
//
//   		// the properties below are optional
//   		AppIdClientRegex: jsii.String("appIdClientRegex"),
//   	},
//   	LambdaAuthorizerConfig: &LambdaAuthorizerConfigProperty{
//   		AuthorizerUri: jsii.String("authorizerUri"),
//
//   		// the properties below are optional
//   		AuthorizerResultTtlInSeconds: jsii.Number(123),
//   		IdentityValidationExpression: jsii.String("identityValidationExpression"),
//   	},
//   	OpenIdConnectConfig: &OpenIDConnectConfigProperty{
//   		Issuer: jsii.String("issuer"),
//
//   		// the properties below are optional
//   		AuthTtl: jsii.Number(123),
//   		ClientId: jsii.String("clientId"),
//   		IatTtl: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-authprovider.html
//
type CfnApi_AuthProviderProperty struct {
	// Security configuration for your AppSync API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-authprovider.html#cfn-appsync-api-authprovider-authtype
	//
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// Optional authorization configuration for using Amazon Cognito user pools with your API endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-authprovider.html#cfn-appsync-api-authprovider-cognitoconfig
	//
	CognitoConfig interface{} `field:"optional" json:"cognitoConfig" yaml:"cognitoConfig"`
	// A LambdaAuthorizerConfig holds configuration on how to authorize AWS AppSync API access when using the AWS_LAMBDA authorizer mode.
	//
	// Be aware that an AWS AppSync API may have only one Lambda authorizer configured at a time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-authprovider.html#cfn-appsync-api-authprovider-lambdaauthorizerconfig
	//
	LambdaAuthorizerConfig interface{} `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// The OpenID Connect configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-authprovider.html#cfn-appsync-api-authprovider-openidconnectconfig
	//
	OpenIdConnectConfig interface{} `field:"optional" json:"openIdConnectConfig" yaml:"openIdConnectConfig"`
}

