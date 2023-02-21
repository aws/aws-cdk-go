package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTrafficMirrorFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficMirrorFilterProps := &cfnTrafficMirrorFilterProps{
//   	description: jsii.String("description"),
//   	networkServices: []*string{
//   		jsii.String("networkServices"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTrafficMirrorFilterProps struct {
	// The description of the Traffic Mirror filter.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The network service traffic that is associated with the Traffic Mirror filter.
	//
	// Valid values are `amazon-dns` .
	NetworkServices *[]*string `field:"optional" json:"networkServices" yaml:"networkServices"`
	// The tags to assign to a Traffic Mirror filter.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

