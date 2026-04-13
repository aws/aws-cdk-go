package awseks


// The warm pool configuration for the node group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   warmPoolConfigProperty := &WarmPoolConfigProperty{
//   	Enabled: jsii.Boolean(false),
//   	MaxGroupPreparedCapacity: jsii.Number(123),
//   	MinSize: jsii.Number(123),
//   	PoolState: jsii.String("poolState"),
//   	ReuseOnScaleIn: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-warmpoolconfig.html
//
type CfnNodegroup_WarmPoolConfigProperty struct {
	// Enable or disable warm pool for the node group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-warmpoolconfig.html#cfn-eks-nodegroup-warmpoolconfig-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum number of instances that are allowed to be in the warm pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-warmpoolconfig.html#cfn-eks-nodegroup-warmpoolconfig-maxgrouppreparedcapacity
	//
	MaxGroupPreparedCapacity *float64 `field:"optional" json:"maxGroupPreparedCapacity" yaml:"maxGroupPreparedCapacity"`
	// The minimum number of instances to maintain in the warm pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-warmpoolconfig.html#cfn-eks-nodegroup-warmpoolconfig-minsize
	//
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// The desired state of warm pool instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-warmpoolconfig.html#cfn-eks-nodegroup-warmpoolconfig-poolstate
	//
	PoolState *string `field:"optional" json:"poolState" yaml:"poolState"`
	// Whether to return instances to the warm pool during scale-in instead of terminating them.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-warmpoolconfig.html#cfn-eks-nodegroup-warmpoolconfig-reuseonscalein
	//
	ReuseOnScaleIn interface{} `field:"optional" json:"reuseOnScaleIn" yaml:"reuseOnScaleIn"`
}

