package mixinsawsecs


// Defines how Amazon ECS Managed Instances optimizes the infrastructure in your capacity provider.
//
// Configure it to turn on or off the infrastructure optimization in your capacity provider, and to control the idle EC2 instances optimization delay.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   infrastructureOptimizationProperty := &InfrastructureOptimizationProperty{
//   	ScaleInAfter: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-infrastructureoptimization.html
//
type CfnCapacityProviderPropsMixin_InfrastructureOptimizationProperty struct {
	// This parameter defines the number of seconds Amazon ECS Managed Instances waits before optimizing EC2 instances that have become idle or underutilized.
	//
	// A longer delay increases the likelihood of placing new tasks on idle instances, reducing startup time. A shorter delay helps reduce infrastructure costs by optimizing idle instances more quickly. Valid values are: Not set (null) - Uses the default optimization behavior, `-1` - Disables automatic infrastructure optimization, `0` to `3600` (inclusive) - Specifies the number of seconds to wait before optimizing instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-infrastructureoptimization.html#cfn-ecs-capacityprovider-infrastructureoptimization-scaleinafter
	//
	ScaleInAfter *float64 `field:"optional" json:"scaleInAfter" yaml:"scaleInAfter"`
}

