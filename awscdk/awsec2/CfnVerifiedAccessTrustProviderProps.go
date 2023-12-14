package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVerifiedAccessTrustProvider`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVerifiedAccessTrustProviderProps := &CfnVerifiedAccessTrustProviderProps{
//   	PolicyReferenceName: jsii.String("policyReferenceName"),
//   	TrustProviderType: jsii.String("trustProviderType"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DeviceOptions: &DeviceOptionsProperty{
//   		PublicSigningKeyUrl: jsii.String("publicSigningKeyUrl"),
//   		TenantId: jsii.String("tenantId"),
//   	},
//   	DeviceTrustProviderType: jsii.String("deviceTrustProviderType"),
//   	OidcOptions: &OidcOptionsProperty{
//   		AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		Issuer: jsii.String("issuer"),
//   		Scope: jsii.String("scope"),
//   		TokenEndpoint: jsii.String("tokenEndpoint"),
//   		UserInfoEndpoint: jsii.String("userInfoEndpoint"),
//   	},
//   	SseSpecification: &SseSpecificationProperty{
//   		CustomerManagedKeyEnabled: jsii.Boolean(false),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserTrustProviderType: jsii.String("userTrustProviderType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccesstrustprovider.html
//
type CfnVerifiedAccessTrustProviderProps struct {
	// The identifier to be used when working with policy rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccesstrustprovider-policyreferencename
	//
	PolicyReferenceName *string `field:"required" json:"policyReferenceName" yaml:"policyReferenceName"`
	// The type of Verified Access trust provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccesstrustprovider-trustprovidertype
	//
	TrustProviderType *string `field:"required" json:"trustProviderType" yaml:"trustProviderType"`
	// A description for the AWS Verified Access trust provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccesstrustprovider-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The options for device-identity trust provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccesstrustprovider-deviceoptions
	//
	DeviceOptions interface{} `field:"optional" json:"deviceOptions" yaml:"deviceOptions"`
	// The type of device-based trust provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccesstrustprovider-devicetrustprovidertype
	//
	DeviceTrustProviderType *string `field:"optional" json:"deviceTrustProviderType" yaml:"deviceTrustProviderType"`
	// The options for an OpenID Connect-compatible user-identity trust provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccesstrustprovider-oidcoptions
	//
	OidcOptions interface{} `field:"optional" json:"oidcOptions" yaml:"oidcOptions"`
	// The options for additional server side encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccesstrustprovider-ssespecification
	//
	SseSpecification interface{} `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// The tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccesstrustprovider-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of user-based trust provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccesstrustprovider-usertrustprovidertype
	//
	UserTrustProviderType *string `field:"optional" json:"userTrustProviderType" yaml:"userTrustProviderType"`
}

