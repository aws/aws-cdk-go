package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   networkInterfaceSpecificationRequestProperty := &NetworkInterfaceSpecificationRequestProperty{
//   	AssociatePublicIpAddress: jsii.Boolean(false),
//   	DeleteOnTermination: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	DeviceIndex: jsii.Number(123),
//   	Groups: []*string{
//   		jsii.String("groups"),
//   	},
//   	InterfaceType: jsii.String("interfaceType"),
//   	Ipv6AddressCount: jsii.Number(123),
//   	Ipv6Addresses: []interface{}{
//   		&Ipv6AddressRequestProperty{
//   			Ipv6Address: jsii.String("ipv6Address"),
//   		},
//   	},
//   	NetworkCardIndex: jsii.Number(123),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	PrivateIpAddress: jsii.String("privateIpAddress"),
//   	PrivateIpAddresses: []interface{}{
//   		&PrivateIpAddressSpecificationRequestProperty{
//   			Primary: jsii.Boolean(false),
//   			PrivateIpAddress: jsii.String("privateIpAddress"),
//   		},
//   	},
//   	SecondaryPrivateIpAddressCount: jsii.Number(123),
//   	SubnetId: jsii.String("subnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-networkinterfacespecificationrequest.html
//
type CfnEC2FleetPropsMixin_NetworkInterfaceSpecificationRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-networkinterfacespecificationrequest.html#cfn-ec2-ec2fleet-networkinterfacespecificationrequest-associatepublicipaddress
	//
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-networkinterfacespecificationrequest.html#cfn-ec2-ec2fleet-networkinterfacespecificationrequest-deleteontermination
	//
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-networkinterfacespecificationrequest.html#cfn-ec2-ec2fleet-networkinterfacespecificationrequest-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-networkinterfacespecificationrequest.html#cfn-ec2-ec2fleet-networkinterfacespecificationrequest-deviceindex
	//
	DeviceIndex *float64 `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-networkinterfacespecificationrequest.html#cfn-ec2-ec2fleet-networkinterfacespecificationrequest-groups
	//
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-networkinterfacespecificationrequest.html#cfn-ec2-ec2fleet-networkinterfacespecificationrequest-interfacetype
	//
	InterfaceType *string `field:"optional" json:"interfaceType" yaml:"interfaceType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-networkinterfacespecificationrequest.html#cfn-ec2-ec2fleet-networkinterfacespecificationrequest-ipv6addresscount
	//
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-networkinterfacespecificationrequest.html#cfn-ec2-ec2fleet-networkinterfacespecificationrequest-ipv6addresses
	//
	Ipv6Addresses interface{} `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-networkinterfacespecificationrequest.html#cfn-ec2-ec2fleet-networkinterfacespecificationrequest-networkcardindex
	//
	NetworkCardIndex *float64 `field:"optional" json:"networkCardIndex" yaml:"networkCardIndex"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-networkinterfacespecificationrequest.html#cfn-ec2-ec2fleet-networkinterfacespecificationrequest-networkinterfaceid
	//
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-networkinterfacespecificationrequest.html#cfn-ec2-ec2fleet-networkinterfacespecificationrequest-privateipaddress
	//
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-networkinterfacespecificationrequest.html#cfn-ec2-ec2fleet-networkinterfacespecificationrequest-privateipaddresses
	//
	PrivateIpAddresses interface{} `field:"optional" json:"privateIpAddresses" yaml:"privateIpAddresses"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-networkinterfacespecificationrequest.html#cfn-ec2-ec2fleet-networkinterfacespecificationrequest-secondaryprivateipaddresscount
	//
	SecondaryPrivateIpAddressCount *float64 `field:"optional" json:"secondaryPrivateIpAddressCount" yaml:"secondaryPrivateIpAddressCount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-networkinterfacespecificationrequest.html#cfn-ec2-ec2fleet-networkinterfacespecificationrequest-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

