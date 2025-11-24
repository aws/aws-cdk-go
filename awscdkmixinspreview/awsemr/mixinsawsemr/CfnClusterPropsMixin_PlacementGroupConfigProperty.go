package mixinsawsemr


// Placement group configuration for an Amazon EMR cluster.
//
// The configuration specifies the placement strategy that can be applied to instance roles during cluster creation.
//
// To use this configuration, consider attaching managed policy AmazonElasticMapReducePlacementGroupPolicy to the Amazon EMR role.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   placementGroupConfigProperty := &PlacementGroupConfigProperty{
//   	InstanceRole: jsii.String("instanceRole"),
//   	PlacementStrategy: jsii.String("placementStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-placementgroupconfig.html
//
type CfnClusterPropsMixin_PlacementGroupConfigProperty struct {
	// Role of the instance in the cluster.
	//
	// Starting with Amazon EMR release 5.23.0, the only supported instance role is `MASTER` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-placementgroupconfig.html#cfn-emr-cluster-placementgroupconfig-instancerole
	//
	InstanceRole *string `field:"optional" json:"instanceRole" yaml:"instanceRole"`
	// Amazon EC2 Placement Group strategy associated with instance role.
	//
	// Starting with Amazon EMR release 5.23.0, the only supported placement strategy is `SPREAD` for the `MASTER` instance role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-placementgroupconfig.html#cfn-emr-cluster-placementgroupconfig-placementstrategy
	//
	PlacementStrategy *string `field:"optional" json:"placementStrategy" yaml:"placementStrategy"`
}

