package awsappfabric

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAppAuthorization`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppAuthorizationProps := &CfnAppAuthorizationProps{
//   	App: jsii.String("app"),
//   	AppBundleIdentifier: jsii.String("appBundleIdentifier"),
//   	AuthType: jsii.String("authType"),
//   	Tenant: &TenantProperty{
//   		TenantDisplayName: jsii.String("tenantDisplayName"),
//   		TenantIdentifier: jsii.String("tenantIdentifier"),
//   	},
//
//   	// the properties below are optional
//   	Credential: &CredentialProperty{
//   		ApiKeyCredential: &ApiKeyCredentialProperty{
//   			ApiKey: jsii.String("apiKey"),
//   		},
//   		Oauth2Credential: &Oauth2CredentialProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-appauthorization.html
//
type CfnAppAuthorizationProps struct {
	// The name of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-appauthorization.html#cfn-appfabric-appauthorization-app
	//
	App *string `field:"required" json:"app" yaml:"app"`
	// The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-appauthorization.html#cfn-appfabric-appauthorization-appbundleidentifier
	//
	AppBundleIdentifier *string `field:"required" json:"appBundleIdentifier" yaml:"appBundleIdentifier"`
	// The authorization type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-appauthorization.html#cfn-appfabric-appauthorization-authtype
	//
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// Contains information about an application tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-appauthorization.html#cfn-appfabric-appauthorization-tenant
	//
	Tenant interface{} `field:"required" json:"tenant" yaml:"tenant"`
	// Contains credential information for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-appauthorization.html#cfn-appfabric-appauthorization-credential
	//
	Credential interface{} `field:"optional" json:"credential" yaml:"credential"`
	// Tags to apply to the app authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-appauthorization.html#cfn-appfabric-appauthorization-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

