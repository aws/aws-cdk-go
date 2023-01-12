package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkBandwidthGbpsProperty := &networkBandwidthGbpsProperty{
//   	max: jsii.Number(123),
//   	min: jsii.Number(123),
//   }
//
type CfnLaunchTemplate_NetworkBandwidthGbpsProperty struct {
	// `CfnLaunchTemplate.NetworkBandwidthGbpsProperty.Max`.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// `CfnLaunchTemplate.NetworkBandwidthGbpsProperty.Min`.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

