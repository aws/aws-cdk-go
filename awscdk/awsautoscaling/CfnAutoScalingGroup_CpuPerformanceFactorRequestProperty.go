package awsautoscaling


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cpuPerformanceFactorRequestProperty := &CpuPerformanceFactorRequestProperty{
//   	References: []interface{}{
//   		&PerformanceFactorReferenceRequestProperty{
//   			InstanceFamily: jsii.String("instanceFamily"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-cpuperformancefactorrequest.html
//
type CfnAutoScalingGroup_CpuPerformanceFactorRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-cpuperformancefactorrequest.html#cfn-autoscaling-autoscalinggroup-cpuperformancefactorrequest-references
	//
	References interface{} `field:"optional" json:"references" yaml:"references"`
}

