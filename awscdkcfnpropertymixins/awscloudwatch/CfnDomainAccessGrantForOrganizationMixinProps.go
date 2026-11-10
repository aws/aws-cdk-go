package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDomainAccessGrantForOrganizationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDomainAccessGrantForOrganizationMixinProps := &CfnDomainAccessGrantForOrganizationMixinProps{
//   	DomainId: jsii.String("domainId"),
//   	Name: jsii.String("name"),
//   	Permission: jsii.String("permission"),
//   	Principal: &OrganizationAccessGrantPrincipalProperty{
//   		PrincipalAttributes: []interface{}{
//   			&PrincipalAttributeProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		PrincipalId: jsii.String("principalId"),
//   		PrincipalType: jsii.String("principalType"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domainaccessgrantfororganization.html
//
type CfnDomainAccessGrantForOrganizationMixinProps struct {
	// The ID of the organization domain the grant belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domainaccessgrantfororganization.html#cfn-cloudwatch-domainaccessgrantfororganization-domainid
	//
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// A name that identifies the access grant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domainaccessgrantfororganization.html#cfn-cloudwatch-domainaccessgrantfororganization-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The permission granted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domainaccessgrantfororganization.html#cfn-cloudwatch-domainaccessgrantfororganization-permission
	//
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
	// The principal the grant is issued to.
	//
	// Supply PrincipalId, or PrincipalAttributes for an attribute-based grant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domainaccessgrantfororganization.html#cfn-cloudwatch-domainaccessgrantfororganization-principal
	//
	Principal interface{} `field:"optional" json:"principal" yaml:"principal"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-domainaccessgrantfororganization.html#cfn-cloudwatch-domainaccessgrantfororganization-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

