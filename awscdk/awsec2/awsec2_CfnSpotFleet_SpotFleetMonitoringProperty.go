package awsec2


// Describes whether monitoring is enabled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spotFleetMonitoringProperty := &spotFleetMonitoringProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnSpotFleet_SpotFleetMonitoringProperty struct {
	// Enables monitoring for the instance.
	//
	// Default: `false`.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

