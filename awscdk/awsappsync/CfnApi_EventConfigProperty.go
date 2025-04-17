package awsappsync


// Describes the authorization configuration for connections, message publishing, message subscriptions, and logging for an Event API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventConfigProperty := &EventConfigProperty{
//   	AuthProviders: []interface{}{
//   		&AuthProviderProperty{
//   			AuthType: jsii.String("authType"),
//
//   			// the properties below are optional
//   			CognitoConfig: &CognitoConfigProperty{
//   				AwsRegion: jsii.String("awsRegion"),
//   				UserPoolId: jsii.String("userPoolId"),
//
//   				// the properties below are optional
//   				AppIdClientRegex: jsii.String("appIdClientRegex"),
//   			},
//   			LambdaAuthorizerConfig: &LambdaAuthorizerConfigProperty{
//   				AuthorizerUri: jsii.String("authorizerUri"),
//
//   				// the properties below are optional
//   				AuthorizerResultTtlInSeconds: jsii.Number(123),
//   				IdentityValidationExpression: jsii.String("identityValidationExpression"),
//   			},
//   			OpenIdConnectConfig: &OpenIDConnectConfigProperty{
//   				Issuer: jsii.String("issuer"),
//
//   				// the properties below are optional
//   				AuthTtl: jsii.Number(123),
//   				ClientId: jsii.String("clientId"),
//   				IatTtl: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ConnectionAuthModes: []interface{}{
//   		&AuthModeProperty{
//   			AuthType: jsii.String("authType"),
//   		},
//   	},
//   	DefaultPublishAuthModes: []interface{}{
//   		&AuthModeProperty{
//   			AuthType: jsii.String("authType"),
//   		},
//   	},
//   	DefaultSubscribeAuthModes: []interface{}{
//   		&AuthModeProperty{
//   			AuthType: jsii.String("authType"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	LogConfig: &EventLogConfigProperty{
//   		CloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   		LogLevel: jsii.String("logLevel"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-eventconfig.html
//
type CfnApi_EventConfigProperty struct {
	// A list of authorization providers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-eventconfig.html#cfn-appsync-api-eventconfig-authproviders
	//
	AuthProviders interface{} `field:"required" json:"authProviders" yaml:"authProviders"`
	// A list of valid authorization modes for the Event API connections.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-eventconfig.html#cfn-appsync-api-eventconfig-connectionauthmodes
	//
	ConnectionAuthModes interface{} `field:"required" json:"connectionAuthModes" yaml:"connectionAuthModes"`
	// A list of valid authorization modes for the Event API publishing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-eventconfig.html#cfn-appsync-api-eventconfig-defaultpublishauthmodes
	//
	DefaultPublishAuthModes interface{} `field:"required" json:"defaultPublishAuthModes" yaml:"defaultPublishAuthModes"`
	// A list of valid authorization modes for the Event API subscriptions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-eventconfig.html#cfn-appsync-api-eventconfig-defaultsubscribeauthmodes
	//
	DefaultSubscribeAuthModes interface{} `field:"required" json:"defaultSubscribeAuthModes" yaml:"defaultSubscribeAuthModes"`
	// The CloudWatch Logs configuration for the Event API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-eventconfig.html#cfn-appsync-api-eventconfig-logconfig
	//
	LogConfig interface{} `field:"optional" json:"logConfig" yaml:"logConfig"`
}

