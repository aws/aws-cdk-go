package awsemr


// `PlacementType` is a property of the `AWS::EMR::Cluster` resource.
//
// `PlacementType` determines the Amazon EC2 Availability Zone configuration of the cluster (job flow).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   placementTypeProperty := &placementTypeProperty{
//   	availabilityZone: jsii.String("availabilityZone"),
//   }
//
type CfnCluster_PlacementTypeProperty struct {
	// The Amazon EC2 Availability Zone for the cluster.
	//
	// `AvailabilityZone` is used for uniform instance groups, while `AvailabilityZones` (plural) is used for instance fleets.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
}

