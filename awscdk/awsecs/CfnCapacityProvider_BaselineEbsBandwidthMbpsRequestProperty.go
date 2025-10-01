package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baselineEbsBandwidthMbpsRequestProperty := &BaselineEbsBandwidthMbpsRequestProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-baselineebsbandwidthmbpsrequest.html
//
type CfnCapacityProvider_BaselineEbsBandwidthMbpsRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-baselineebsbandwidthmbpsrequest.html#cfn-ecs-capacityprovider-baselineebsbandwidthmbpsrequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-baselineebsbandwidthmbpsrequest.html#cfn-ecs-capacityprovider-baselineebsbandwidthmbpsrequest-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

