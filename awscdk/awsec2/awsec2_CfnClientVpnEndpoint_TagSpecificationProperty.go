package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The tags to apply to a resource when the resource is being created.
//
// > The `Valid Values` lists all the resource types that can be tagged. However, the action you're using might not support tagging all of these resource types. If you try to tag a resource type that is unsupported for the action you're using, you'll get an error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagSpecificationProperty := &tagSpecificationProperty{
//   	resourceType: jsii.String("resourceType"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnClientVpnEndpoint_TagSpecificationProperty struct {
	// The type of resource to tag.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// The tags to apply to the resource.
	Tags *[]*awscdk.CfnTag `field:"required" json:"tags" yaml:"tags"`
}

