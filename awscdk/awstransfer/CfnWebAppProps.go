package awstransfer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnWebApp`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWebAppProps := &CfnWebAppProps{
//   	IdentityProviderDetails: &IdentityProviderDetailsProperty{
//   		ApplicationArn: jsii.String("applicationArn"),
//   		InstanceArn: jsii.String("instanceArn"),
//   		Role: jsii.String("role"),
//   	},
//
//   	// the properties below are optional
//   	AccessEndpoint: jsii.String("accessEndpoint"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WebAppCustomization: &WebAppCustomizationProperty{
//   		FaviconFile: jsii.String("faviconFile"),
//   		LogoFile: jsii.String("logoFile"),
//   		Title: jsii.String("title"),
//   	},
//   	WebAppEndpointPolicy: jsii.String("webAppEndpointPolicy"),
//   	WebAppUnits: &WebAppUnitsProperty{
//   		Provisioned: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-webapp.html
//
type CfnWebAppProps struct {
	// You can provide a structure that contains the details for the identity provider to use with your web app.
	//
	// For more details about this parameter, see [Configure your identity provider for Transfer Family web apps](https://docs.aws.amazon.com//transfer/latest/userguide/webapp-identity-center.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-webapp.html#cfn-transfer-webapp-identityproviderdetails
	//
	IdentityProviderDetails interface{} `field:"required" json:"identityProviderDetails" yaml:"identityProviderDetails"`
	// The `AccessEndpoint` is the URL that you provide to your users for them to interact with the Transfer Family web app.
	//
	// You can specify a custom URL or use the default value.
	//
	// Before you enter a custom URL for this parameter, follow the steps described in [Update your access endpoint with a custom URL](https://docs.aws.amazon.com//transfer/latest/userguide/webapp-customize.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-webapp.html#cfn-transfer-webapp-accessendpoint
	//
	AccessEndpoint *string `field:"optional" json:"accessEndpoint" yaml:"accessEndpoint"`
	// Key-value pairs that can be used to group and search for web apps.
	//
	// Tags are metadata attached to web apps for any purpose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-webapp.html#cfn-transfer-webapp-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A structure that contains the customization fields for the web app.
	//
	// You can provide a title, logo, and icon to customize the appearance of your web app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-webapp.html#cfn-transfer-webapp-webappcustomization
	//
	WebAppCustomization interface{} `field:"optional" json:"webAppCustomization" yaml:"webAppCustomization"`
	// Setting for the type of endpoint policy for the web app. The default value is `STANDARD` .
	//
	// If your web app was created in an AWS GovCloud (US) Region , the value of this parameter can be `FIPS` , which indicates the web app endpoint is FIPS-compliant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-webapp.html#cfn-transfer-webapp-webappendpointpolicy
	//
	WebAppEndpointPolicy *string `field:"optional" json:"webAppEndpointPolicy" yaml:"webAppEndpointPolicy"`
	// A union that contains the value for number of concurrent connections or the user sessions on your web app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-webapp.html#cfn-transfer-webapp-webappunits
	//
	WebAppUnits interface{} `field:"optional" json:"webAppUnits" yaml:"webAppUnits"`
}

