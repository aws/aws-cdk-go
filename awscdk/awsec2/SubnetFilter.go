package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Contains logic which chooses a set of subnets from a larger list, in conjunction with SubnetSelection, to determine where to place AWS resources such as VPC endpoints, EC2 instances, etc.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetFilter := awscdk.Aws_ec2.SubnetFilter_AvailabilityZones([]*string{
//   	jsii.String("availabilityZones"),
//   })
//
type SubnetFilter interface {
	// Executes the subnet filtering logic, returning a filtered set of subnets.
	SelectSubnets(_subnets *[]ISubnet) *[]ISubnet
}

// The jsii proxy struct for SubnetFilter
type jsiiProxy_SubnetFilter struct {
	_ byte // padding
}

func NewSubnetFilter_Override(s SubnetFilter) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.SubnetFilter",
		nil, // no parameters
		s,
	)
}

// Chooses subnets which are in one of the given availability zones.
func SubnetFilter_AvailabilityZones(availabilityZones *[]*string) SubnetFilter {
	_init_.Initialize()

	if err := validateSubnetFilter_AvailabilityZonesParameters(availabilityZones); err != nil {
		panic(err)
	}
	var returns SubnetFilter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.SubnetFilter",
		"availabilityZones",
		[]interface{}{availabilityZones},
		&returns,
	)

	return returns
}

// Chooses subnets which have the provided CIDR netmask.
func SubnetFilter_ByCidrMask(mask *float64) SubnetFilter {
	_init_.Initialize()

	if err := validateSubnetFilter_ByCidrMaskParameters(mask); err != nil {
		panic(err)
	}
	var returns SubnetFilter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.SubnetFilter",
		"byCidrMask",
		[]interface{}{mask},
		&returns,
	)

	return returns
}

// Chooses subnets which are inside any of the specified CIDR range.
func SubnetFilter_ByCidrRanges(cidrs *[]*string) SubnetFilter {
	_init_.Initialize()

	if err := validateSubnetFilter_ByCidrRangesParameters(cidrs); err != nil {
		panic(err)
	}
	var returns SubnetFilter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.SubnetFilter",
		"byCidrRanges",
		[]interface{}{cidrs},
		&returns,
	)

	return returns
}

// Chooses subnets by id.
func SubnetFilter_ByIds(subnetIds *[]*string) SubnetFilter {
	_init_.Initialize()

	if err := validateSubnetFilter_ByIdsParameters(subnetIds); err != nil {
		panic(err)
	}
	var returns SubnetFilter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.SubnetFilter",
		"byIds",
		[]interface{}{subnetIds},
		&returns,
	)

	return returns
}

// Chooses subnets which contain any of the specified IP addresses.
func SubnetFilter_ContainsIpAddresses(ipv4addrs *[]*string) SubnetFilter {
	_init_.Initialize()

	if err := validateSubnetFilter_ContainsIpAddressesParameters(ipv4addrs); err != nil {
		panic(err)
	}
	var returns SubnetFilter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.SubnetFilter",
		"containsIpAddresses",
		[]interface{}{ipv4addrs},
		&returns,
	)

	return returns
}

// Chooses subnets such that there is at most one per availability zone.
func SubnetFilter_OnePerAz() SubnetFilter {
	_init_.Initialize()

	var returns SubnetFilter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.SubnetFilter",
		"onePerAz",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubnetFilter) SelectSubnets(_subnets *[]ISubnet) *[]ISubnet {
	if err := s.validateSelectSubnetsParameters(_subnets); err != nil {
		panic(err)
	}
	var returns *[]ISubnet

	_jsii_.Invoke(
		s,
		"selectSubnets",
		[]interface{}{_subnets},
		&returns,
	)

	return returns
}

