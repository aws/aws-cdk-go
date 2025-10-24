package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApi`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiProps := &CfnApiProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	EventConfig: &EventConfigProperty{
//   		AuthProviders: []interface{}{
//   			&AuthProviderProperty{
//   				AuthType: jsii.String("authType"),
//
//   				// the properties below are optional
//   				CognitoConfig: &CognitoConfigProperty{
//   					AwsRegion: jsii.String("awsRegion"),
//   					UserPoolId: jsii.String("userPoolId"),
//
//   					// the properties below are optional
//   					AppIdClientRegex: jsii.String("appIdClientRegex"),
//   				},
//   				LambdaAuthorizerConfig: &LambdaAuthorizerConfigProperty{
//   					AuthorizerUri: jsii.String("authorizerUri"),
//
//   					// the properties below are optional
//   					AuthorizerResultTtlInSeconds: jsii.Number(123),
//   					IdentityValidationExpression: jsii.String("identityValidationExpression"),
//   				},
//   				OpenIdConnectConfig: &OpenIDConnectConfigProperty{
//   					Issuer: jsii.String("issuer"),
//
//   					// the properties below are optional
//   					AuthTtl: jsii.Number(123),
//   					ClientId: jsii.String("clientId"),
//   					IatTtl: jsii.Number(123),
//   				},
//   			},
//   		},
//   		ConnectionAuthModes: []interface{}{
//   			&AuthModeProperty{
//   				AuthType: jsii.String("authType"),
//   			},
//   		},
//   		DefaultPublishAuthModes: []interface{}{
//   			&AuthModeProperty{
//   				AuthType: jsii.String("authType"),
//   			},
//   		},
//   		DefaultSubscribeAuthModes: []interface{}{
//   			&AuthModeProperty{
//   				AuthType: jsii.String("authType"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		LogConfig: &EventLogConfigProperty{
//   			CloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   	},
//   	OwnerContact: jsii.String("ownerContact"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-api.html
//
type CfnApiProps struct {
	// The name of the `Api` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-api.html#cfn-appsync-api-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Describes the authorization configuration for connections, message publishing, message subscriptions, and logging for an Event API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-api.html#cfn-appsync-api-eventconfig
	//
	EventConfig interface{} `field:"optional" json:"eventConfig" yaml:"eventConfig"`
	// The owner contact information for an API resource.
	//
	// This field accepts any string input with a length of 0 - 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-api.html#cfn-appsync-api-ownercontact
	//
	OwnerContact *string `field:"optional" json:"ownerContact" yaml:"ownerContact"`
	// A set of tags (key-value pairs) for this API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-api.html#cfn-appsync-api-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

