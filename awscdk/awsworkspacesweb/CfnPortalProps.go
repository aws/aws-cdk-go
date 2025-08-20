package awsworkspacesweb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPortal`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPortalProps := &CfnPortalProps{
//   	AdditionalEncryptionContext: map[string]*string{
//   		"additionalEncryptionContextKey": jsii.String("additionalEncryptionContext"),
//   	},
//   	AuthenticationType: jsii.String("authenticationType"),
//   	BrowserSettingsArn: jsii.String("browserSettingsArn"),
//   	CustomerManagedKey: jsii.String("customerManagedKey"),
//   	DataProtectionSettingsArn: jsii.String("dataProtectionSettingsArn"),
//   	DisplayName: jsii.String("displayName"),
//   	InstanceType: jsii.String("instanceType"),
//   	IpAccessSettingsArn: jsii.String("ipAccessSettingsArn"),
//   	MaxConcurrentSessions: jsii.Number(123),
//   	NetworkSettingsArn: jsii.String("networkSettingsArn"),
//   	SessionLoggerArn: jsii.String("sessionLoggerArn"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrustStoreArn: jsii.String("trustStoreArn"),
//   	UserAccessLoggingSettingsArn: jsii.String("userAccessLoggingSettingsArn"),
//   	UserSettingsArn: jsii.String("userSettingsArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-portal.html
//
type CfnPortalProps struct {
	// The additional encryption context of the portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-portal.html#cfn-workspacesweb-portal-additionalencryptioncontext
	//
	AdditionalEncryptionContext interface{} `field:"optional" json:"additionalEncryptionContext" yaml:"additionalEncryptionContext"`
	// The type of authentication integration points used when signing into the web portal. Defaults to `Standard` .
	//
	// `Standard` web portals are authenticated directly through your identity provider (IdP). User and group access to your web portal is controlled through your IdP. You need to include an IdP resource in your template to integrate your IdP with your web portal. Completing the configuration for your IdP requires exchanging WorkSpaces Secure Browser’s SP metadata with your IdP’s IdP metadata. If your IdP requires the SP metadata first before returning the IdP metadata, you should follow these steps:
	//
	// 1. Create and deploy a CloudFormation template with a `Standard` portal with no `IdentityProvider` resource.
	//
	// 2. Retrieve the SP metadata using `Fn:GetAtt` , the WorkSpaces Secure Browser console, or by the calling the `GetPortalServiceProviderMetadata` API.
	//
	// 3. Submit the data to your IdP.
	//
	// 4. Add an `IdentityProvider` resource to your CloudFormation template.
	//
	// `IAM Identity Center` web portals are authenticated through AWS IAM Identity Center . They provide additional features, such as IdP-initiated authentication. Identity sources (including external identity provider integration) and other identity provider information must be configured in IAM Identity Center . User and group assignment must be done through the WorkSpaces Secure Browser console. These cannot be configured in CloudFormation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-portal.html#cfn-workspacesweb-portal-authenticationtype
	//
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// The ARN of the browser settings that is associated with this web portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-portal.html#cfn-workspacesweb-portal-browsersettingsarn
	//
	BrowserSettingsArn *string `field:"optional" json:"browserSettingsArn" yaml:"browserSettingsArn"`
	// The customer managed key of the web portal.
	//
	// *Pattern* : `^arn:[\w+=\/,.@-]+:kms:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:key\/[a-zA-Z0-9-]+$`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-portal.html#cfn-workspacesweb-portal-customermanagedkey
	//
	CustomerManagedKey *string `field:"optional" json:"customerManagedKey" yaml:"customerManagedKey"`
	// The ARN of the data protection settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-portal.html#cfn-workspacesweb-portal-dataprotectionsettingsarn
	//
	DataProtectionSettingsArn *string `field:"optional" json:"dataProtectionSettingsArn" yaml:"dataProtectionSettingsArn"`
	// The name of the web portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-portal.html#cfn-workspacesweb-portal-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The type and resources of the underlying instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-portal.html#cfn-workspacesweb-portal-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The ARN of the IP access settings that is associated with the web portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-portal.html#cfn-workspacesweb-portal-ipaccesssettingsarn
	//
	IpAccessSettingsArn *string `field:"optional" json:"ipAccessSettingsArn" yaml:"ipAccessSettingsArn"`
	// The maximum number of concurrent sessions for the portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-portal.html#cfn-workspacesweb-portal-maxconcurrentsessions
	//
	MaxConcurrentSessions *float64 `field:"optional" json:"maxConcurrentSessions" yaml:"maxConcurrentSessions"`
	// The ARN of the network settings that is associated with the web portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-portal.html#cfn-workspacesweb-portal-networksettingsarn
	//
	NetworkSettingsArn *string `field:"optional" json:"networkSettingsArn" yaml:"networkSettingsArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-portal.html#cfn-workspacesweb-portal-sessionloggerarn
	//
	SessionLoggerArn *string `field:"optional" json:"sessionLoggerArn" yaml:"sessionLoggerArn"`
	// The tags to add to the web portal.
	//
	// A tag is a key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-portal.html#cfn-workspacesweb-portal-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN of the trust store that is associated with the web portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-portal.html#cfn-workspacesweb-portal-truststorearn
	//
	TrustStoreArn *string `field:"optional" json:"trustStoreArn" yaml:"trustStoreArn"`
	// The ARN of the user access logging settings that is associated with the web portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-portal.html#cfn-workspacesweb-portal-useraccessloggingsettingsarn
	//
	UserAccessLoggingSettingsArn *string `field:"optional" json:"userAccessLoggingSettingsArn" yaml:"userAccessLoggingSettingsArn"`
	// The ARN of the user settings that is associated with the web portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-portal.html#cfn-workspacesweb-portal-usersettingsarn
	//
	UserSettingsArn *string `field:"optional" json:"userSettingsArn" yaml:"userSettingsArn"`
}

