package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDomainPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDomainMixinProps := &CfnDomainMixinProps{
//   	IdentityProviderConfiguration: &IdentityProviderConfigurationProperty{
//   		IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   			IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   		},
//   	},
//   	IdentityProviders: []*string{
//   		jsii.String("identityProviders"),
//   	},
//   	Name: jsii.String("name"),
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
type CfnDomainMixinProps struct {
	// Identity provider configuration for the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domain.html#cfn-cloudwatch-domain-identityproviderconfiguration
	//
	IdentityProviderConfiguration interface{} `field:"optional" json:"identityProviderConfiguration" yaml:"identityProviderConfiguration"`
	// The identity providers configured for the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domain.html#cfn-cloudwatch-domain-identityproviders
	//
	IdentityProviders *[]*string `field:"optional" json:"identityProviders" yaml:"identityProviders"`
	// A name that identifies the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domain.html#cfn-cloudwatch-domain-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domain.html#cfn-cloudwatch-domain-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

