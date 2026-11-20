package awssso

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTrustedTokenIssuer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrustedTokenIssuerProps := &CfnTrustedTokenIssuerProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//   	TrustedTokenIssuerConfiguration: &TrustedTokenIssuerConfigurationProperty{
//   		OidcJwtConfiguration: &OidcJwtConfigurationProperty{
//   			ClaimAttributePath: jsii.String("claimAttributePath"),
//   			IdentityStoreAttributePath: jsii.String("identityStoreAttributePath"),
//   			IssuerUrl: jsii.String("issuerUrl"),
//   			JwksRetrievalOption: jsii.String("jwksRetrievalOption"),
//   		},
//   	},
//   	TrustedTokenIssuerType: jsii.String("trustedTokenIssuerType"),
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-trustedtokenissuer.html
//
type CfnTrustedTokenIssuerProps struct {
	// The ARN of the IAM Identity Center instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-trustedtokenissuer.html#cfn-sso-trustedtokenissuer-instancearn
	//
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the trusted token issuer configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-trustedtokenissuer.html#cfn-sso-trustedtokenissuer-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The configuration of a trusted token issuer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-trustedtokenissuer.html#cfn-sso-trustedtokenissuer-trustedtokenissuerconfiguration
	//
	TrustedTokenIssuerConfiguration interface{} `field:"required" json:"trustedTokenIssuerConfiguration" yaml:"trustedTokenIssuerConfiguration"`
	// The type of the trusted token issuer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-trustedtokenissuer.html#cfn-sso-trustedtokenissuer-trustedtokenissuertype
	//
	TrustedTokenIssuerType *string `field:"required" json:"trustedTokenIssuerType" yaml:"trustedTokenIssuerType"`
	// A list of tags to attach to the trusted token issuer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-trustedtokenissuer.html#cfn-sso-trustedtokenissuer-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

