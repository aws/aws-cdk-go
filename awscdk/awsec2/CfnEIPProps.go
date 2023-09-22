package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEIP`.
//
// Example:
//   var instance instance
//
//   var myZone hostedZone
//
//
//   elasticIp := ec2.NewCfnEIP(this, jsii.String("EIP"), &CfnEIPProps{
//   	Domain: jsii.String("vpc"),
//   	InstanceId: instance.InstanceId,
//   })
//   route53.NewARecord(this, jsii.String("ARecord"), &ARecordProps{
//   	Zone: myZone,
//   	Target: route53.RecordTarget_FromIpAddresses(elasticIp.ref),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-eip.html
//
type CfnEIPProps struct {
	// The network ( `vpc` ).
	//
	// If you define an Elastic IP address and associate it with a VPC that is defined in the same template, you must declare a dependency on the VPC-gateway attachment by using the [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) on this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-eip.html#cfn-ec2-eip-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The ID of the instance.
	//
	// > Updates to the `InstanceId` property may require *some interruptions* . Updates on an EIP reassociates the address on its associated resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-eip.html#cfn-ec2-eip-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// A unique set of Availability Zones, Local Zones, or Wavelength Zones from which AWS advertises IP addresses.
	//
	// Use this parameter to limit the IP address to this location. IP addresses cannot move between network border groups.
	//
	// Use [DescribeAvailabilityZones](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAvailabilityZones.html) to view the network border groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-eip.html#cfn-ec2-eip-networkbordergroup
	//
	NetworkBorderGroup *string `field:"optional" json:"networkBorderGroup" yaml:"networkBorderGroup"`
	// The ID of an address pool that you own.
	//
	// Use this parameter to let Amazon EC2 select an address from the address pool.
	//
	// > Updates to the `PublicIpv4Pool` property may require *some interruptions* . Updates on an EIP reassociates the address on its associated resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-eip.html#cfn-ec2-eip-publicipv4pool
	//
	PublicIpv4Pool *string `field:"optional" json:"publicIpv4Pool" yaml:"publicIpv4Pool"`
	// Any tags assigned to the Elastic IP address.
	//
	// > Updates to the `Tags` property may require *some interruptions* . Updates on an EIP reassociates the address on its associated resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-eip.html#cfn-ec2-eip-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The Elastic IP address you are accepting for transfer.
	//
	// You can only accept one transferred address. For more information on Elastic IP address transfers, see [Transfer Elastic IP addresses](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#transfer-EIPs-intro) in the *Amazon Virtual Private Cloud User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-eip.html#cfn-ec2-eip-transferaddress
	//
	TransferAddress *string `field:"optional" json:"transferAddress" yaml:"transferAddress"`
}

