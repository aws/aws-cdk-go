package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnIPAMScope`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIPAMScopeProps := &cfnIPAMScopeProps{
//   	ipamId: jsii.String("ipamId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnIPAMScopeProps struct {
	// The ID of the IPAM for which you're creating this scope.
	IpamId *string `field:"required" json:"ipamId" yaml:"ipamId"`
	// The description of the scope.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The key/value combination of a tag assigned to the resource.
	//
	// Use the tag key in the filter name and the tag value as the filter value. For example, to find all resources that have a tag with the key `Owner` and the value `TeamA` , specify `tag:Owner` for the filter name and `TeamA` for the filter value.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

