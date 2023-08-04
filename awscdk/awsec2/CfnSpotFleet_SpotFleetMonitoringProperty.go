package awsec2


// Describes whether monitoring is enabled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spotFleetMonitoringProperty := &SpotFleetMonitoringProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetmonitoring.html
//
type CfnSpotFleet_SpotFleetMonitoringProperty struct {
	// Enables monitoring for the instance.
	//
	// Default: `false`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetmonitoring.html#cfn-ec2-spotfleet-spotfleetmonitoring-enabled
	//
	// Default: - false.
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

