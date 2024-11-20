package awsgamelift


// Defines the range of ports on the instance that allow inbound traffic to connect with containers in a fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionPortRangeProperty := &ConnectionPortRangeProperty{
//   	FromPort: jsii.Number(123),
//   	ToPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-connectionportrange.html
//
type CfnContainerFleet_ConnectionPortRangeProperty struct {
	// A starting value for a range of allowed port numbers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-connectionportrange.html#cfn-gamelift-containerfleet-connectionportrange-fromport
	//
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// An ending value for a range of allowed port numbers.
	//
	// Port numbers are end-inclusive. This value must be higher than FromPort.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-connectionportrange.html#cfn-gamelift-containerfleet-connectionportrange-toport
	//
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
}

