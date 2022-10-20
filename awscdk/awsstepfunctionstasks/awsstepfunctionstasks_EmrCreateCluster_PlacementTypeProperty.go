package awsstepfunctionstasks


// The Amazon EC2 Availability Zone configuration of the cluster (job flow).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   placementTypeProperty := &placementTypeProperty{
//   	availabilityZone: jsii.String("availabilityZone"),
//   	availabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_PlacementType.html
//
type EmrCreateCluster_PlacementTypeProperty struct {
	// The Amazon EC2 Availability Zone for the cluster.
	//
	// AvailabilityZone is used for uniform instance groups, while AvailabilityZones
	// (plural) is used for instance fleets.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// When multiple Availability Zones are specified, Amazon EMR evaluates them and launches instances in the optimal Availability Zone.
	//
	// AvailabilityZones is used for instance fleets, while AvailabilityZone (singular) is used for uniform instance groups.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
}

