package awsautoscaling


// A structure that specifies an instance reuse policy for the `InstanceReusePolicy` property of the [AWS::AutoScaling::WarmPool](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-warmpool.html) resource.
//
// For more information, see [Warm pools for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-warm-pools.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceReusePolicyProperty := &InstanceReusePolicyProperty{
//   	ReuseOnScaleIn: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-warmpool-instancereusepolicy.html
//
type CfnWarmPool_InstanceReusePolicyProperty struct {
	// Specifies whether instances in the Auto Scaling group can be returned to the warm pool on scale in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-warmpool-instancereusepolicy.html#cfn-autoscaling-warmpool-instancereusepolicy-reuseonscalein
	//
	ReuseOnScaleIn interface{} `field:"optional" json:"reuseOnScaleIn" yaml:"reuseOnScaleIn"`
}

