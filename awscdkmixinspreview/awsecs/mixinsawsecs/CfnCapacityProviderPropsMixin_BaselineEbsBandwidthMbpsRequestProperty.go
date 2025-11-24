package mixinsawsecs


// The minimum and maximum baseline Amazon EBS bandwidth in megabits per second (Mbps) for instance type selection.
//
// This is important for workloads with high storage I/O requirements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   baselineEbsBandwidthMbpsRequestProperty := &BaselineEbsBandwidthMbpsRequestProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-baselineebsbandwidthmbpsrequest.html
//
type CfnCapacityProviderPropsMixin_BaselineEbsBandwidthMbpsRequestProperty struct {
	// The maximum baseline Amazon EBS bandwidth in Mbps.
	//
	// Instance types with higher Amazon EBS bandwidth are excluded from selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-baselineebsbandwidthmbpsrequest.html#cfn-ecs-capacityprovider-baselineebsbandwidthmbpsrequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum baseline Amazon EBS bandwidth in Mbps.
	//
	// Instance types with lower Amazon EBS bandwidth are excluded from selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-baselineebsbandwidthmbpsrequest.html#cfn-ecs-capacityprovider-baselineebsbandwidthmbpsrequest-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

