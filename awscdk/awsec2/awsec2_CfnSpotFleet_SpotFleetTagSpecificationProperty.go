package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// The tags for a Spot Fleet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spotFleetTagSpecificationProperty := &spotFleetTagSpecificationProperty{
//   	resourceType: jsii.String("resourceType"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnSpotFleet_SpotFleetTagSpecificationProperty struct {
	// The type of resource.
	//
	// Currently, the only resource type that is supported is `instance` . To tag the Spot Fleet request on creation, use the `TagSpecifications` parameter in [`SpotFleetRequestConfigData`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SpotFleetRequestConfigData.html) .
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

