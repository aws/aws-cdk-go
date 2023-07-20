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
//   zoneAwarenessConfigProperty := &ZoneAwarenessConfigProperty{
//   	AvailabilityZoneCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-zoneawarenessconfig.html
//
type CfnDomain_ZoneAwarenessConfigProperty struct {
	// If you enabled multiple Availability Zones (AZs), the number of AZs that you want the domain to use.
	//
	// Valid values are `2` and `3` . Default is 2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-zoneawarenessconfig.html#cfn-opensearchservice-domain-zoneawarenessconfig-availabilityzonecount
	//
	AvailabilityZoneCount *float64 `field:"optional" json:"availabilityZoneCount" yaml:"availabilityZoneCount"`
}

