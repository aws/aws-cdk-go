package awsec2

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Result of selecting a subset of subnets from a VPC.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   vpc := ec2.NewVpc(this, jsii.String("TheVPC"), &vpcProps{
//   	ipAddresses: ipAddresses_Cidr(jsii.String("10.0.0.0/16")),
//   })
//
//   // Iterate the private subnets
//   selection := vpc.selectSubnets(&subnetSelection{
//   	subnetType: ec2.subnetType_PRIVATE_WITH_EGRESS,
//   })
//
//   for _, subnet := range selection.subnets {}
//
type SelectedSubnets struct {
	// The respective AZs of each subnet.
	AvailabilityZones *[]*string `field:"required" json:"availabilityZones" yaml:"availabilityZones"`
	// Whether any of the given subnets are from the VPC's public subnets.
	HasPublic *bool `field:"required" json:"hasPublic" yaml:"hasPublic"`
	// Dependency representing internet connectivity for these subnets.
	InternetConnectivityEstablished constructs.IDependable `field:"required" json:"internetConnectivityEstablished" yaml:"internetConnectivityEstablished"`
	// The subnet IDs.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Selected subnet objects.
	Subnets *[]ISubnet `field:"required" json:"subnets" yaml:"subnets"`
	// The subnet selection is not actually real yet.
	//
	// If this value is true, don't validate anything about the subnets. The count
	// or identities are not known yet, and the validation will most likely fail
	// which will prevent a successful lookup.
	IsPendingLookup *bool `field:"optional" json:"isPendingLookup" yaml:"isPendingLookup"`
}

