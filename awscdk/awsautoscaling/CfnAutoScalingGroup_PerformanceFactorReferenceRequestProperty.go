package awsautoscaling


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   performanceFactorReferenceRequestProperty := &PerformanceFactorReferenceRequestProperty{
//   	InstanceFamily: jsii.String("instanceFamily"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-performancefactorreferencerequest.html
//
type CfnAutoScalingGroup_PerformanceFactorReferenceRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-performancefactorreferencerequest.html#cfn-autoscaling-autoscalinggroup-performancefactorreferencerequest-instancefamily
	//
	InstanceFamily *string `field:"optional" json:"instanceFamily" yaml:"instanceFamily"`
}

