package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDomainForOrganization`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainForOrganizationProps := &CfnDomainForOrganizationProps{
//   	DomainAccessRoleArn: jsii.String("domainAccessRoleArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domainfororganization.html
//
type CfnDomainForOrganizationProps struct {
	// The ARN of the customer-provided IAM role in the management account used for domain access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domainfororganization.html#cfn-cloudwatch-domainfororganization-domainaccessrolearn
	//
	DomainAccessRoleArn *string `field:"required" json:"domainAccessRoleArn" yaml:"domainAccessRoleArn"`
	// The identity providers configured for the organization domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domainfororganization.html#cfn-cloudwatch-domainfororganization-identityproviders
	//
	IdentityProviders *[]*string `field:"required" json:"identityProviders" yaml:"identityProviders"`
	// A name that identifies the organization domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domainfororganization.html#cfn-cloudwatch-domainfororganization-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Identity provider configuration for the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domainfororganization.html#cfn-cloudwatch-domainfororganization-identityproviderconfiguration
	//
	IdentityProviderConfiguration interface{} `field:"optional" json:"identityProviderConfiguration" yaml:"identityProviderConfiguration"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domainfororganization.html#cfn-cloudwatch-domainfororganization-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

