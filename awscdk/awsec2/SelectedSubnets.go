package awsec2

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Result of selecting a subset of subnets from a VPC.
//
// Example:
//   vpc := ec2.NewVpc(this, jsii.String("VpcWithExplicitAzs"), &VpcProps{
//   	AvailabilityZones: []*string{
//   		jsii.String("us-east-1a"),
//   		jsii.String("us-east-1b"),
//   	},
//   	NatGateways: jsii.Number(0),
//   	SubnetConfiguration: []SubnetConfiguration{
//   		&SubnetConfiguration{
//   			Name: jsii.String("Application"),
//   			SubnetType: ec2.SubnetType_PRIVATE_ISOLATED,
//   		},
//   	},
//   })
//
//   selectedSubnets := vpc.selectSubnets(&SubnetSelection{
//   	SubnetType: ec2.SubnetType_PRIVATE_ISOLATED,
//   	AvailabilityZones: []*string{
//   		jsii.String("us-east-1a"),
//   	},
//   })
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
	// Default: false.
	//
	IsPendingLookup *bool `field:"optional" json:"isPendingLookup" yaml:"isPendingLookup"`
}

