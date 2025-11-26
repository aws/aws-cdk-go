package previewawsecsmixins


// The minimum and maximum network bandwidth in gigabits per second (Gbps) for instance type selection.
//
// This is important for network-intensive workloads.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkBandwidthGbpsRequestProperty := &NetworkBandwidthGbpsRequestProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-networkbandwidthgbpsrequest.html
//
type CfnCapacityProviderPropsMixin_NetworkBandwidthGbpsRequestProperty struct {
	// The maximum network bandwidth in Gbps.
	//
	// Instance types with higher network bandwidth are excluded from selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-networkbandwidthgbpsrequest.html#cfn-ecs-capacityprovider-networkbandwidthgbpsrequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum network bandwidth in Gbps.
	//
	// Instance types with lower network bandwidth are excluded from selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-networkbandwidthgbpsrequest.html#cfn-ecs-capacityprovider-networkbandwidthgbpsrequest-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

