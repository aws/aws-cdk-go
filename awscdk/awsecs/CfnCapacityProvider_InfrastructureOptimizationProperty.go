package awsecs


// The configuration that controls how Amazon ECS optimizes your infrastructure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   infrastructureOptimizationProperty := &InfrastructureOptimizationProperty{
//   	ScaleInAfter: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-infrastructureoptimization.html
//
type CfnCapacityProvider_InfrastructureOptimizationProperty struct {
	// This parameter defines the number of seconds Amazon ECS Managed Instances waits before optimizing EC2 instances that have become idle or underutilized.
	//
	// A longer delay increases the likelihood of placing new tasks on idle or underutilized instances instances, reducing startup time. A shorter delay helps reduce infrastructure costs by optimizing idle or underutilized instances,instances more quickly.
	//
	// Valid values are:
	//
	// - `null` - Uses the default optimization behavior.
	// - `-1` - Disables automatic infrastructure optimization.
	// - A value between `0` and `3600` (inclusive) - Specifies the number of seconds to wait before optimizing instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-infrastructureoptimization.html#cfn-ecs-capacityprovider-infrastructureoptimization-scaleinafter
	//
	ScaleInAfter *float64 `field:"optional" json:"scaleInAfter" yaml:"scaleInAfter"`
}

