package awsoutposts

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnOutpostPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnOutpostMixinProps := &CfnOutpostMixinProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	SiteId: jsii.String("siteId"),
//   	SupportedHardwareType: jsii.String("supportedHardwareType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-outpost.html
//
type CfnOutpostMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-outpost.html#cfn-outposts-outpost-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-outpost.html#cfn-outposts-outpost-availabilityzoneid
	//
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-outpost.html#cfn-outposts-outpost-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-outpost.html#cfn-outposts-outpost-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-outpost.html#cfn-outposts-outpost-siteid
	//
	SiteId *string `field:"optional" json:"siteId" yaml:"siteId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-outpost.html#cfn-outposts-outpost-supportedhardwaretype
	//
	SupportedHardwareType *string `field:"optional" json:"supportedHardwareType" yaml:"supportedHardwareType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-outpost.html#cfn-outposts-outpost-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

