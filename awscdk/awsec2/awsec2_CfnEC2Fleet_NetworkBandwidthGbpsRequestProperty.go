package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkBandwidthGbpsRequestProperty := &networkBandwidthGbpsRequestProperty{
//   	max: jsii.Number(123),
//   	min: jsii.Number(123),
//   }
//
type CfnEC2Fleet_NetworkBandwidthGbpsRequestProperty struct {
	// `CfnEC2Fleet.NetworkBandwidthGbpsRequestProperty.Max`.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// `CfnEC2Fleet.NetworkBandwidthGbpsRequestProperty.Min`.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

