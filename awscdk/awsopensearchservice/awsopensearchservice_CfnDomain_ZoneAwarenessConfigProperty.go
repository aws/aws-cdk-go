package awsopensearchservice


// Specifies zone awareness configuration options.
//
// Only use if `ZoneAwarenessEnabled` is `true` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   zoneAwarenessConfigProperty := &zoneAwarenessConfigProperty{
//   	availabilityZoneCount: jsii.Number(123),
//   }
//
type CfnDomain_ZoneAwarenessConfigProperty struct {
	// If you enabled multiple Availability Zones (AZs), the number of AZs that you want the domain to use.
	//
	// Valid values are `2` and `3` . Default is 2.
	AvailabilityZoneCount *float64 `field:"optional" json:"availabilityZoneCount" yaml:"availabilityZoneCount"`
}

