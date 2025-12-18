package awseks


// Specify granular overrides for specific repair actions.
//
// These overrides control the repair action and the repair delay time before a node is considered eligible for repair. If you use this, you must specify all the values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeRepairConfigOverridesProperty := &NodeRepairConfigOverridesProperty{
//   	MinRepairWaitTimeMins: jsii.Number(123),
//   	NodeMonitoringCondition: jsii.String("nodeMonitoringCondition"),
//   	NodeUnhealthyReason: jsii.String("nodeUnhealthyReason"),
//   	RepairAction: jsii.String("repairAction"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-noderepairconfigoverrides.html
//
type CfnNodegroup_NodeRepairConfigOverridesProperty struct {
	// Specify the minimum time in minutes to wait before attempting to repair a node with this specific `nodeMonitoringCondition` and `nodeUnhealthyReason` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-noderepairconfigoverrides.html#cfn-eks-nodegroup-noderepairconfigoverrides-minrepairwaittimemins
	//
	MinRepairWaitTimeMins *float64 `field:"optional" json:"minRepairWaitTimeMins" yaml:"minRepairWaitTimeMins"`
	// Specify an unhealthy condition reported by the node monitoring agent that this override would apply to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-noderepairconfigoverrides.html#cfn-eks-nodegroup-noderepairconfigoverrides-nodemonitoringcondition
	//
	NodeMonitoringCondition *string `field:"optional" json:"nodeMonitoringCondition" yaml:"nodeMonitoringCondition"`
	// Specify a reason reported by the node monitoring agent that this override would apply to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-noderepairconfigoverrides.html#cfn-eks-nodegroup-noderepairconfigoverrides-nodeunhealthyreason
	//
	NodeUnhealthyReason *string `field:"optional" json:"nodeUnhealthyReason" yaml:"nodeUnhealthyReason"`
	// Specify the repair action to take for nodes when all of the specified conditions are met.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-noderepairconfigoverrides.html#cfn-eks-nodegroup-noderepairconfigoverrides-repairaction
	//
	RepairAction *string `field:"optional" json:"repairAction" yaml:"repairAction"`
}

