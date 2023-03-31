package awsec2


// Describes Spot Instance placement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spotPlacementProperty := &spotPlacementProperty{
//   	availabilityZone: jsii.String("availabilityZone"),
//   	groupName: jsii.String("groupName"),
//   	tenancy: jsii.String("tenancy"),
//   }
//
type CfnSpotFleet_SpotPlacementProperty struct {
	// The Availability Zone.
	//
	// To specify multiple Availability Zones, separate them using commas; for example, "us-west-2a, us-west-2b".
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The name of the placement group.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The tenancy of the instance (if the instance is running in a VPC).
	//
	// An instance with a tenancy of `dedicated` runs on single-tenant hardware. The `host` tenancy is not supported for Spot Instances.
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
}

