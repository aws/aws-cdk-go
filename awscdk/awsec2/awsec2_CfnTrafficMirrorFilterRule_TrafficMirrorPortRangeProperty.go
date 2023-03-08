package awsec2


// Describes the Traffic Mirror port range.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficMirrorPortRangeProperty := &trafficMirrorPortRangeProperty{
//   	fromPort: jsii.Number(123),
//   	toPort: jsii.Number(123),
//   }
//
type CfnTrafficMirrorFilterRule_TrafficMirrorPortRangeProperty struct {
	// The start of the Traffic Mirror port range.
	//
	// This applies to the TCP and UDP protocols.
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// The end of the Traffic Mirror port range.
	//
	// This applies to the TCP and UDP protocols.
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
}

