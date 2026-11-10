package awsroute53profiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnProfileAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnProfileAssociationMixinProps := &CfnProfileAssociationMixinProps{
//   	Arn: jsii.String("arn"),
//   	Name: jsii.String("name"),
//   	Priority: jsii.Number(123),
//   	ProfileId: jsii.String("profileId"),
//   	ResourceId: jsii.String("resourceId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53profiles-profileassociation.html
//
type CfnProfileAssociationMixinProps struct {
	// The Amazon Resource Name (ARN) of the profile association to a VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53profiles-profileassociation.html#cfn-route53profiles-profileassociation-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Name of the Profile association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53profiles-profileassociation.html#cfn-route53profiles-profileassociation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The priority of the profile association, used to order the profiles applied to a VPC when multiple profiles are associated with it.
	//
	// Lower values have higher precedence. Omit for the first profile associated with a VPC. This orders whole profiles on a VPC and is distinct from the DNS Firewall rule group priority carried in a ProfileResourceAssociation's ResourceProperties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53profiles-profileassociation.html#cfn-route53profiles-profileassociation-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// ID of the Profile.
	//
	// Update to this property requires update to the `ResourceId` property as well, because you can only associate one Profile per VPC. For more information, see [Route 53 Profiles](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/profiles.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53profiles-profileassociation.html#cfn-route53profiles-profileassociation-profileid
	//
	ProfileId *string `field:"optional" json:"profileId" yaml:"profileId"`
	// The ID of the VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53profiles-profileassociation.html#cfn-route53profiles-profileassociation-resourceid
	//
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53profiles-profileassociation.html#cfn-route53profiles-profileassociation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

