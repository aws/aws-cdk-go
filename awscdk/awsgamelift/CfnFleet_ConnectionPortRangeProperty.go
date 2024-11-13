package awsgamelift


// *This data type is currently not available.
//
// It is under improvement as we respond to customer feedback from the Containers public preview.*
//
// The set of port numbers to open on each instance in a container fleet. Connection ports are used by inbound traffic to connect with processes that are running in containers on the fleet.
//
// *Part of:* `ContainerGroupsConfiguration` , `ContainerGroupsAttributes`.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-connectionportrange.html
//
type CfnFleet_ConnectionPortRangeProperty struct {
	// Starting value for the port range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-connectionportrange.html#cfn-gamelift-fleet-connectionportrange-fromport
	//
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// Ending value for the port.
	//
	// Port numbers are end-inclusive. This value must be equal to or greater than `FromPort` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-connectionportrange.html#cfn-gamelift-fleet-connectionportrange-toport
	//
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
}

