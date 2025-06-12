package awsemr


// Managed scaling policy for an Amazon EMR cluster.
//
// The policy specifies the limits for resources that can be added or terminated from a cluster. The policy only applies to the core and task nodes. The master node cannot be scaled after initial configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedScalingPolicyProperty := &ManagedScalingPolicyProperty{
//   	ComputeLimits: &ComputeLimitsProperty{
//   		MaximumCapacityUnits: jsii.Number(123),
//   		MinimumCapacityUnits: jsii.Number(123),
//   		UnitType: jsii.String("unitType"),
//
//   		// the properties below are optional
//   		MaximumCoreCapacityUnits: jsii.Number(123),
//   		MaximumOnDemandCapacityUnits: jsii.Number(123),
//   	},
//   	ScalingStrategy: jsii.String("scalingStrategy"),
//   	UtilizationPerformanceIndex: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-managedscalingpolicy.html
//
type CfnCluster_ManagedScalingPolicyProperty struct {
	// The Amazon EC2 unit limits for a managed scaling policy.
	//
	// The managed scaling activity of a cluster is not allowed to go above or below these limits. The limit only applies to the core and task nodes. The master node cannot be scaled after initial configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-managedscalingpolicy.html#cfn-emr-cluster-managedscalingpolicy-computelimits
	//
	ComputeLimits interface{} `field:"optional" json:"computeLimits" yaml:"computeLimits"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-managedscalingpolicy.html#cfn-emr-cluster-managedscalingpolicy-scalingstrategy
	//
	ScalingStrategy *string `field:"optional" json:"scalingStrategy" yaml:"scalingStrategy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-managedscalingpolicy.html#cfn-emr-cluster-managedscalingpolicy-utilizationperformanceindex
	//
	UtilizationPerformanceIndex *float64 `field:"optional" json:"utilizationPerformanceIndex" yaml:"utilizationPerformanceIndex"`
}

