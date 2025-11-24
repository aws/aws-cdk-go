package mixinsawscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnIdentityPoolPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var cognitoEvents interface{}
//   var supportedLoginProviders interface{}
//
//   cfnIdentityPoolMixinProps := &CfnIdentityPoolMixinProps{
//   	AllowClassicFlow: jsii.Boolean(false),
//   	AllowUnauthenticatedIdentities: jsii.Boolean(false),
//   	CognitoEvents: cognitoEvents,
//   	CognitoIdentityProviders: []interface{}{
//   		&CognitoIdentityProviderProperty{
//   			ClientId: jsii.String("clientId"),
//   			ProviderName: jsii.String("providerName"),
//   			ServerSideTokenCheck: jsii.Boolean(false),
//   		},
//   	},
//   	CognitoStreams: &CognitoStreamsProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		StreamingStatus: jsii.String("streamingStatus"),
//   		StreamName: jsii.String("streamName"),
//   	},
//   	DeveloperProviderName: jsii.String("developerProviderName"),
//   	IdentityPoolName: jsii.String("identityPoolName"),
//   	IdentityPoolTags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	OpenIdConnectProviderArns: []*string{
//   		jsii.String("openIdConnectProviderArns"),
//   	},
//   	PushSync: &PushSyncProperty{
//   		ApplicationArns: []*string{
//   			jsii.String("applicationArns"),
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	SamlProviderArns: []*string{
//   		jsii.String("samlProviderArns"),
//   	},
//   	SupportedLoginProviders: supportedLoginProviders,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html
//
type CfnIdentityPoolMixinProps struct {
	// Enables the Basic (Classic) authentication flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-allowclassicflow
	//
	AllowClassicFlow interface{} `field:"optional" json:"allowClassicFlow" yaml:"allowClassicFlow"`
	// Specifies whether the identity pool supports unauthenticated logins.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-allowunauthenticatedidentities
	//
	AllowUnauthenticatedIdentities interface{} `field:"optional" json:"allowUnauthenticatedIdentities" yaml:"allowUnauthenticatedIdentities"`
	// The events to configure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitoevents
	//
	CognitoEvents interface{} `field:"optional" json:"cognitoEvents" yaml:"cognitoEvents"`
	// The Amazon Cognito user pools and their client IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitoidentityproviders
	//
	CognitoIdentityProviders interface{} `field:"optional" json:"cognitoIdentityProviders" yaml:"cognitoIdentityProviders"`
	// Configuration options for configuring Amazon Cognito streams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitostreams
	//
	CognitoStreams interface{} `field:"optional" json:"cognitoStreams" yaml:"cognitoStreams"`
	// The "domain" Amazon Cognito uses when referencing your users.
	//
	// This name acts as a placeholder that allows your backend and the Amazon Cognito service to communicate about the developer provider. For the `DeveloperProviderName` , you can use letters and periods (.), underscores (_), and dashes (-).
	//
	// *Minimum length* : 1
	//
	// *Maximum length* : 100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-developerprovidername
	//
	DeveloperProviderName *string `field:"optional" json:"developerProviderName" yaml:"developerProviderName"`
	// The name of your Amazon Cognito identity pool.
	//
	// *Minimum length* : 1
	//
	// *Maximum length* : 128
	//
	// *Pattern* : `[\w\s+=,.@-]+`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-identitypoolname
	//
	IdentityPoolName *string `field:"optional" json:"identityPoolName" yaml:"identityPoolName"`
	// Tags to assign to the identity pool.
	//
	// A tag is a label that you can apply to identity pools to categorize and manage them in different ways, such as by purpose, owner, environment, or other criteria.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-identitypooltags
	//
	IdentityPoolTags *[]*awscdk.CfnTag `field:"optional" json:"identityPoolTags" yaml:"identityPoolTags"`
	// The Amazon Resource Names (ARNs) of the OpenID connect providers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-openidconnectproviderarns
	//
	OpenIdConnectProviderArns *[]*string `field:"optional" json:"openIdConnectProviderArns" yaml:"openIdConnectProviderArns"`
	// The configuration options to be applied to the identity pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-pushsync
	//
	PushSync interface{} `field:"optional" json:"pushSync" yaml:"pushSync"`
	// The Amazon Resource Names (ARNs) of the Security Assertion Markup Language (SAML) providers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-samlproviderarns
	//
	SamlProviderArns *[]*string `field:"optional" json:"samlProviderArns" yaml:"samlProviderArns"`
	// Key-value pairs that map provider names to provider app IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-supportedloginproviders
	//
	SupportedLoginProviders interface{} `field:"optional" json:"supportedLoginProviders" yaml:"supportedLoginProviders"`
}

