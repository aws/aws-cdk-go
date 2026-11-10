package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDomain`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainProps := &CfnDomainProps{
//   	IdentityProviders: []*string{
//   		jsii.String("identityProviders"),
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	IdentityProviderConfiguration: &IdentityProviderConfigurationProperty{
//   		IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   			IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domain.html
//
type CfnDomainProps struct {
	// The identity providers configured for the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domain.html#cfn-cloudwatch-domain-identityproviders
	//
	IdentityProviders *[]*string `field:"required" json:"identityProviders" yaml:"identityProviders"`
	// A name that identifies the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domain.html#cfn-cloudwatch-domain-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Identity provider configuration for the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domain.html#cfn-cloudwatch-domain-identityproviderconfiguration
	//
	IdentityProviderConfiguration interface{} `field:"optional" json:"identityProviderConfiguration" yaml:"identityProviderConfiguration"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domain.html#cfn-cloudwatch-domain-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

