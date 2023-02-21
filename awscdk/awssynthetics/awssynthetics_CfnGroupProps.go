package awssynthetics

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGroupProps := &cfnGroupProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	resourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnGroupProps struct {
	// A name for the group. It can include any Unicode characters.
	//
	// The names for all groups in your account, across all Regions, must be unique.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARNs of the canaries that you want to associate with this group.
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// The list of key-value pairs that are associated with the group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

