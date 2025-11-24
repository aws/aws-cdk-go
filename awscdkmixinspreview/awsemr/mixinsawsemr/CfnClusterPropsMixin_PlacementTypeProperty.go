package mixinsawsemr


// `PlacementType` is a property of the `AWS::EMR::Cluster` resource.
//
// `PlacementType` determines the Amazon EC2 Availability Zone configuration of the cluster (job flow).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   placementTypeProperty := &PlacementTypeProperty{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-placementtype.html
//
type CfnClusterPropsMixin_PlacementTypeProperty struct {
	// The Amazon EC2 Availability Zone for the cluster.
	//
	// `AvailabilityZone` is used for uniform instance groups, while `AvailabilityZones` (plural) is used for instance fleets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-placementtype.html#cfn-emr-cluster-placementtype-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
}

