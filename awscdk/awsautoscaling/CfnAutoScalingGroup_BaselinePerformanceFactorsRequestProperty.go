package awsautoscaling


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baselinePerformanceFactorsRequestProperty := &BaselinePerformanceFactorsRequestProperty{
//   	Cpu: &CpuPerformanceFactorRequestProperty{
//   		References: []interface{}{
//   			&PerformanceFactorReferenceRequestProperty{
//   				InstanceFamily: jsii.String("instanceFamily"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-baselineperformancefactorsrequest.html
//
type CfnAutoScalingGroup_BaselinePerformanceFactorsRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-baselineperformancefactorsrequest.html#cfn-autoscaling-autoscalinggroup-baselineperformancefactorsrequest-cpu
	//
	Cpu interface{} `field:"optional" json:"cpu" yaml:"cpu"`
}

