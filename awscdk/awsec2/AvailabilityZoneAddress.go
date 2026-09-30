package awsec2


// Configuration for a specific Availability Zone in a Regional NAT Gateway.
//
// This is used to manually specify which EIPs to use in each AZ.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   availabilityZoneAddress := &AvailabilityZoneAddress{
//   	AllocationIds: []*string{
//   		jsii.String("allocationIds"),
//   	},
//
//   	// the properties below are optional
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   }
//
type AvailabilityZoneAddress struct {
	// The allocation IDs of the Elastic IP addresses to be used for handling outbound NAT traffic in this specific Availability Zone.
	AllocationIds *[]*string `field:"required" json:"allocationIds" yaml:"allocationIds"`
	// The Availability Zone where this specific NAT gateway configuration will be active.
	//
	// Either `availabilityZone` or `availabilityZoneId` must be specified.
	// Default: - Use availabilityZoneId instead.
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The ID of the Availability Zone where this specific NAT gateway configuration will be active.
	//
	// Either `availabilityZone` or `availabilityZoneId` must be specified.
	// Default: - Use availabilityZone instead.
	//
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
}

