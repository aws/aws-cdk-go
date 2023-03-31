package awsec2


// Describes a range of ports.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portRangeProperty := &portRangeProperty{
//   	from: jsii.Number(123),
//   	to: jsii.Number(123),
//   }
//
type CfnNetworkInsightsAnalysis_PortRangeProperty struct {
	// The first port in the range.
	From *float64 `field:"optional" json:"from" yaml:"from"`
	// The last port in the range.
	To *float64 `field:"optional" json:"to" yaml:"to"`
}

