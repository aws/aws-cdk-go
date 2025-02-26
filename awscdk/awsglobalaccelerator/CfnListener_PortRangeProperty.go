package awsglobalaccelerator


// A complex type for a range of ports for a listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portRangeProperty := &PortRangeProperty{
//   	FromPort: jsii.Number(123),
//   	ToPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-listener-portrange.html
//
type CfnListener_PortRangeProperty struct {
	// The first port in the range of ports, inclusive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-listener-portrange.html#cfn-globalaccelerator-listener-portrange-fromport
	//
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// The last port in the range of ports, inclusive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-listener-portrange.html#cfn-globalaccelerator-listener-portrange-toport
	//
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
}

