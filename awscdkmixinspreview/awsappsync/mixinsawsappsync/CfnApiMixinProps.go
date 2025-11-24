package mixinsawsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnApiPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApiMixinProps := &CfnApiMixinProps{
//   	EventConfig: &EventConfigProperty{
//   		AuthProviders: []interface{}{
//   			&AuthProviderProperty{
//   				AuthType: jsii.String("authType"),
//   				CognitoConfig: &CognitoConfigProperty{
//   					AppIdClientRegex: jsii.String("appIdClientRegex"),
//   					AwsRegion: jsii.String("awsRegion"),
//   					UserPoolId: jsii.String("userPoolId"),
//   				},
//   				LambdaAuthorizerConfig: &LambdaAuthorizerConfigProperty{
//   					AuthorizerResultTtlInSeconds: jsii.Number(123),
//   					AuthorizerUri: jsii.String("authorizerUri"),
//   					IdentityValidationExpression: jsii.String("identityValidationExpression"),
//   				},
//   				OpenIdConnectConfig: &OpenIDConnectConfigProperty{
//   					AuthTtl: jsii.Number(123),
//   					ClientId: jsii.String("clientId"),
//   					IatTtl: jsii.Number(123),
//   					Issuer: jsii.String("issuer"),
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
//   		LogConfig: &EventLogConfigProperty{
//   			CloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   	},
//   	Name: jsii.String("name"),
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
type CfnApiMixinProps struct {
	// Describes the authorization configuration for connections, message publishing, message subscriptions, and logging for an Event API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-api.html#cfn-appsync-api-eventconfig
	//
	EventConfig interface{} `field:"optional" json:"eventConfig" yaml:"eventConfig"`
	// The name of the `Api` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-api.html#cfn-appsync-api-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
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

