package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// An array of key-value pairs to apply to this resource.
//
// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
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
type CfnCapacityReservation_TagSpecificationProperty struct {
	// The type of resource to tag.
	//
	// Specify `capacity-reservation` .
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The tags to apply to the resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

