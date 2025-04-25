package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDHCPOptions`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDHCPOptionsProps := &CfnDHCPOptionsProps{
//   	DomainName: jsii.String("domainName"),
//   	DomainNameServers: []*string{
//   		jsii.String("domainNameServers"),
//   	},
//   	Ipv6AddressPreferredLeaseTime: jsii.Number(123),
//   	NetbiosNameServers: []*string{
//   		jsii.String("netbiosNameServers"),
//   	},
//   	NetbiosNodeType: jsii.Number(123),
//   	NtpServers: []*string{
//   		jsii.String("ntpServers"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcpoptions.html
//
type CfnDHCPOptionsProps struct {
	// This value is used to complete unqualified DNS hostnames.
	//
	// If you're using AmazonProvidedDNS in `us-east-1` , specify `ec2.internal` . If you're using AmazonProvidedDNS in another Region, specify *region* . `compute.internal` (for example, `ap-northeast-1.compute.internal` ). Otherwise, specify a domain name (for example, *MyCompany.com* ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcpoptions.html#cfn-ec2-dhcpoptions-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The IPv4 addresses of up to four domain name servers, or `AmazonProvidedDNS` .
	//
	// The default is `AmazonProvidedDNS` . To have your instance receive a custom DNS hostname as specified in `DomainName` , you must set this property to a custom DNS server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcpoptions.html#cfn-ec2-dhcpoptions-domainnameservers
	//
	DomainNameServers *[]*string `field:"optional" json:"domainNameServers" yaml:"domainNameServers"`
	// A value (in seconds, minutes, hours, or years) for how frequently a running instance with an IPv6 assigned to it goes through DHCPv6 lease renewal.
	//
	// Acceptable values are between 140 and 2147483647 seconds (approximately 68 years). If no value is entered, the default lease time is 140 seconds. If you use long-term addressing for EC2 instances, you can increase the lease time and avoid frequent lease renewal requests. Lease renewal typically occurs when half of the lease time has elapsed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcpoptions.html#cfn-ec2-dhcpoptions-ipv6addresspreferredleasetime
	//
	Ipv6AddressPreferredLeaseTime *float64 `field:"optional" json:"ipv6AddressPreferredLeaseTime" yaml:"ipv6AddressPreferredLeaseTime"`
	// The IPv4 addresses of up to four NetBIOS name servers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcpoptions.html#cfn-ec2-dhcpoptions-netbiosnameservers
	//
	NetbiosNameServers *[]*string `field:"optional" json:"netbiosNameServers" yaml:"netbiosNameServers"`
	// The NetBIOS node type (1, 2, 4, or 8).
	//
	// We recommend that you specify 2 (broadcast and multicast are not currently supported).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcpoptions.html#cfn-ec2-dhcpoptions-netbiosnodetype
	//
	NetbiosNodeType *float64 `field:"optional" json:"netbiosNodeType" yaml:"netbiosNodeType"`
	// The IPv4 addresses of up to four Network Time Protocol (NTP) servers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcpoptions.html#cfn-ec2-dhcpoptions-ntpservers
	//
	NtpServers *[]*string `field:"optional" json:"ntpServers" yaml:"ntpServers"`
	// Any tags assigned to the DHCP options set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcpoptions.html#cfn-ec2-dhcpoptions-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

