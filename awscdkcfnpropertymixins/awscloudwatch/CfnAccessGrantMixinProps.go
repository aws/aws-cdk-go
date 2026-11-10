package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAccessGrantPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnAccessGrantMixinProps := &CfnAccessGrantMixinProps{
//   	DomainId: jsii.String("domainId"),
//   	Name: jsii.String("name"),
//   	Permission: jsii.String("permission"),
//   	Principal: &AccessGrantPrincipalProperty{
//   		PrincipalAttributes: []interface{}{
//   			&AccessGrantPrincipalAttributeProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		PrincipalId: jsii.String("principalId"),
//   		PrincipalType: jsii.String("principalType"),
//   	},
//   	ScopedActions: []interface{}{
//   		&ScopedActionsProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			ContextConditions: map[string][]*string{
//   				"contextConditionsKey": []*string{
//   					jsii.String("contextConditions"),
//   				},
//   			},
//   			Resources: []interface{}{
//   				&ResourceScopeProperty{
//   					ResourceArns: []*string{
//   						jsii.String("resourceArns"),
//   					},
//   					ResourceType: jsii.String("resourceType"),
//   					RowScopeGroups: []interface{}{
//   						[]interface{}{
//   							&RowScopeProperty{
//   								Field: jsii.String("field"),
//   								Operator: jsii.String("operator"),
//   								Values: []*string{
//   									jsii.String("values"),
//   								},
//   							},
//   						},
//   					},
//   					SignalTypes: []*string{
//   						jsii.String("signalTypes"),
//   					},
//   					Tags: map[string]*string{
//   						"tagsKey": jsii.String("tags"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	SpaceId: jsii.String("spaceId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-accessgrant.html
//
type CfnAccessGrantMixinProps struct {
	// The ID of the domain the grant belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-accessgrant.html#cfn-cloudwatch-accessgrant-domainid
	//
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// A name that identifies the access grant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-accessgrant.html#cfn-cloudwatch-accessgrant-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The permission granted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-accessgrant.html#cfn-cloudwatch-accessgrant-permission
	//
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
	// The principal receiving the grant.
	//
	// Specify PrincipalId, PrincipalAttributes, or both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-accessgrant.html#cfn-cloudwatch-accessgrant-principal
	//
	Principal interface{} `field:"optional" json:"principal" yaml:"principal"`
	// Groups of actions allowed by the grant, each with the resource scopes and conditions that limit those actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-accessgrant.html#cfn-cloudwatch-accessgrant-scopedactions
	//
	ScopedActions interface{} `field:"optional" json:"scopedActions" yaml:"scopedActions"`
	// The space this grant applies to.
	//
	// Domain-scoped grants are returned by ListDomainAccessGrantsForOrganization instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-accessgrant.html#cfn-cloudwatch-accessgrant-spaceid
	//
	SpaceId *string `field:"optional" json:"spaceId" yaml:"spaceId"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-accessgrant.html#cfn-cloudwatch-accessgrant-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

