package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// The tags to apply to a resource when the resource is being created.
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
type CfnCapacityReservationFleet_TagSpecificationProperty struct {
	// The type of resource to tag on creation. Specify `capacity-reservation-fleet` .
	//
	// To tag a resource after it has been created, see [CreateTags](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTags.html) .
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The tags to apply to the resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

