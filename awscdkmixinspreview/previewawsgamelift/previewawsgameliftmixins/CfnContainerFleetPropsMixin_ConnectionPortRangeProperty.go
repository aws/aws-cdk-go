package previewawsgameliftmixins


// The set of port numbers to open on each instance in a container fleet.
//
// Connection ports are used by inbound traffic to connect with processes that are running in containers on the fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   connectionPortRangeProperty := &ConnectionPortRangeProperty{
//   	FromPort: jsii.Number(123),
//   	ToPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-connectionportrange.html
//
type CfnContainerFleetPropsMixin_ConnectionPortRangeProperty struct {
	// Starting value for the port range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-connectionportrange.html#cfn-gamelift-containerfleet-connectionportrange-fromport
	//
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// Ending value for the port.
	//
	// Port numbers are end-inclusive. This value must be equal to or greater than `FromPort` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-connectionportrange.html#cfn-gamelift-containerfleet-connectionportrange-toport
	//
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

