package awssynthetics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGroupProps := &CfnGroupProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ResourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-group.html
//
type CfnGroupProps struct {
	// A name for the group. It can include any Unicode characters.
	//
	// The names for all groups in your account, across all Regions, must be unique.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-group.html#cfn-synthetics-group-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARNs of the canaries that you want to associate with this group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-group.html#cfn-synthetics-group-resourcearns
	//
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// The list of key-value pairs that are associated with the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-group.html#cfn-synthetics-group-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

