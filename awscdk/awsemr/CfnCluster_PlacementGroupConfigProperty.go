package awsemr


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   placementGroupConfigProperty := &PlacementGroupConfigProperty{
//   	InstanceRole: jsii.String("instanceRole"),
//
//   	// the properties below are optional
//   	PlacementStrategy: jsii.String("placementStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-placementgroupconfig.html
//
type CfnCluster_PlacementGroupConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-placementgroupconfig.html#cfn-emr-cluster-placementgroupconfig-instancerole
	//
	InstanceRole *string `field:"required" json:"instanceRole" yaml:"instanceRole"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-placementgroupconfig.html#cfn-emr-cluster-placementgroupconfig-placementstrategy
	//
	PlacementStrategy *string `field:"optional" json:"placementStrategy" yaml:"placementStrategy"`
}

