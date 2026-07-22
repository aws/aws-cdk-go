package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnOAuthClientApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOAuthClientApplicationProps := &CfnOAuthClientApplicationProps{
//   	Name: jsii.String("name"),
//   	OAuthClientApplicationId: jsii.String("oAuthClientApplicationId"),
//   	OAuthClientAuthenticationType: jsii.String("oAuthClientAuthenticationType"),
//   	OAuthTokenEndpointUrl: jsii.String("oAuthTokenEndpointUrl"),
//
//   	// the properties below are optional
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	DataSourceType: jsii.String("dataSourceType"),
//   	IdentityProviderVpcConnectionProperties: &IdentityProviderVpcConnectionPropertiesProperty{
//   		VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   	},
//   	OAuthAuthorizationEndpointUrl: jsii.String("oAuthAuthorizationEndpointUrl"),
//   	OAuthScopes: jsii.String("oAuthScopes"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-oauthclientapplication.html
//
type CfnOAuthClientApplicationProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-oauthclientapplication.html#cfn-quicksight-oauthclientapplication-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-oauthclientapplication.html#cfn-quicksight-oauthclientapplication-oauthclientapplicationid
	//
	OAuthClientApplicationId *string `field:"required" json:"oAuthClientApplicationId" yaml:"oAuthClientApplicationId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-oauthclientapplication.html#cfn-quicksight-oauthclientapplication-oauthclientauthenticationtype
	//
	OAuthClientAuthenticationType *string `field:"required" json:"oAuthClientAuthenticationType" yaml:"oAuthClientAuthenticationType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-oauthclientapplication.html#cfn-quicksight-oauthclientapplication-oauthtokenendpointurl
	//
	OAuthTokenEndpointUrl *string `field:"required" json:"oAuthTokenEndpointUrl" yaml:"oAuthTokenEndpointUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-oauthclientapplication.html#cfn-quicksight-oauthclientapplication-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-oauthclientapplication.html#cfn-quicksight-oauthclientapplication-clientsecret
	//
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-oauthclientapplication.html#cfn-quicksight-oauthclientapplication-datasourcetype
	//
	DataSourceType *string `field:"optional" json:"dataSourceType" yaml:"dataSourceType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-oauthclientapplication.html#cfn-quicksight-oauthclientapplication-identityprovidervpcconnectionproperties
	//
	IdentityProviderVpcConnectionProperties interface{} `field:"optional" json:"identityProviderVpcConnectionProperties" yaml:"identityProviderVpcConnectionProperties"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-oauthclientapplication.html#cfn-quicksight-oauthclientapplication-oauthauthorizationendpointurl
	//
	OAuthAuthorizationEndpointUrl *string `field:"optional" json:"oAuthAuthorizationEndpointUrl" yaml:"oAuthAuthorizationEndpointUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-oauthclientapplication.html#cfn-quicksight-oauthclientapplication-oauthscopes
	//
	OAuthScopes *string `field:"optional" json:"oAuthScopes" yaml:"oAuthScopes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-oauthclientapplication.html#cfn-quicksight-oauthclientapplication-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

