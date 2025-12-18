package previewawseksmixins


// The node auto repair configuration for the node group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   nodeRepairConfigProperty := &NodeRepairConfigProperty{
//   	Enabled: jsii.Boolean(false),
//   	MaxParallelNodesRepairedCount: jsii.Number(123),
//   	MaxParallelNodesRepairedPercentage: jsii.Number(123),
//   	MaxUnhealthyNodeThresholdCount: jsii.Number(123),
//   	MaxUnhealthyNodeThresholdPercentage: jsii.Number(123),
//   	NodeRepairConfigOverrides: []interface{}{
//   		&NodeRepairConfigOverridesProperty{
//   			MinRepairWaitTimeMins: jsii.Number(123),
//   			NodeMonitoringCondition: jsii.String("nodeMonitoringCondition"),
//   			NodeUnhealthyReason: jsii.String("nodeUnhealthyReason"),
//   			RepairAction: jsii.String("repairAction"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-noderepairconfig.html
//
type CfnNodegroupPropsMixin_NodeRepairConfigProperty struct {
	// Specifies whether to enable node auto repair for the node group.
	//
	// Node auto repair is disabled by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-noderepairconfig.html#cfn-eks-nodegroup-noderepairconfig-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specify the maximum number of nodes that can be repaired concurrently or in parallel, expressed as a count of unhealthy nodes.
	//
	// This gives you finer-grained control over the pace of node replacements. When using this, you cannot also set `maxParallelNodesRepairedPercentage` at the same time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-noderepairconfig.html#cfn-eks-nodegroup-noderepairconfig-maxparallelnodesrepairedcount
	//
	MaxParallelNodesRepairedCount *float64 `field:"optional" json:"maxParallelNodesRepairedCount" yaml:"maxParallelNodesRepairedCount"`
	// Specify the maximum number of nodes that can be repaired concurrently or in parallel, expressed as a percentage of unhealthy nodes.
	//
	// This gives you finer-grained control over the pace of node replacements. When using this, you cannot also set `maxParallelNodesRepairedCount` at the same time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-noderepairconfig.html#cfn-eks-nodegroup-noderepairconfig-maxparallelnodesrepairedpercentage
	//
	MaxParallelNodesRepairedPercentage *float64 `field:"optional" json:"maxParallelNodesRepairedPercentage" yaml:"maxParallelNodesRepairedPercentage"`
	// Specify a count threshold of unhealthy nodes, above which node auto repair actions will stop.
	//
	// When using this, you cannot also set `maxUnhealthyNodeThresholdPercentage` at the same time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-noderepairconfig.html#cfn-eks-nodegroup-noderepairconfig-maxunhealthynodethresholdcount
	//
	MaxUnhealthyNodeThresholdCount *float64 `field:"optional" json:"maxUnhealthyNodeThresholdCount" yaml:"maxUnhealthyNodeThresholdCount"`
	// Specify a percentage threshold of unhealthy nodes, above which node auto repair actions will stop.
	//
	// When using this, you cannot also set `maxUnhealthyNodeThresholdCount` at the same time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-noderepairconfig.html#cfn-eks-nodegroup-noderepairconfig-maxunhealthynodethresholdpercentage
	//
	MaxUnhealthyNodeThresholdPercentage *float64 `field:"optional" json:"maxUnhealthyNodeThresholdPercentage" yaml:"maxUnhealthyNodeThresholdPercentage"`
	// Specify granular overrides for specific repair actions.
	//
	// These overrides control the repair action and the repair delay time before a node is considered eligible for repair. If you use this, you must specify all the values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-noderepairconfig.html#cfn-eks-nodegroup-noderepairconfig-noderepairconfigoverrides
	//
	NodeRepairConfigOverrides interface{} `field:"optional" json:"nodeRepairConfigOverrides" yaml:"nodeRepairConfigOverrides"`
}

