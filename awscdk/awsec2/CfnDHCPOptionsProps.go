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
type CfnDHCPOptionsProps struct {
	// This value is used to complete unqualified DNS hostnames.
	//
	// If you're using AmazonProvidedDNS in `us-east-1` , specify `ec2.internal` . If you're using AmazonProvidedDNS in another Region, specify *region* . `compute.internal` (for example, `ap-northeast-1.compute.internal` ). Otherwise, specify a domain name (for example, *MyCompany.com* ).
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The IPv4 addresses of up to four domain name servers, or `AmazonProvidedDNS` .
	//
	// The default is `AmazonProvidedDNS` . To have your instance receive a custom DNS hostname as specified in `DomainName` , you must set this property to a custom DNS server.
	DomainNameServers *[]*string `field:"optional" json:"domainNameServers" yaml:"domainNameServers"`
	// The IPv4 addresses of up to four NetBIOS name servers.
	NetbiosNameServers *[]*string `field:"optional" json:"netbiosNameServers" yaml:"netbiosNameServers"`
	// The NetBIOS node type (1, 2, 4, or 8).
	//
	// We recommend that you specify 2 (broadcast and multicast are not currently supported).
	NetbiosNodeType *float64 `field:"optional" json:"netbiosNodeType" yaml:"netbiosNodeType"`
	// The IPv4 addresses of up to four Network Time Protocol (NTP) servers.
	NtpServers *[]*string `field:"optional" json:"ntpServers" yaml:"ntpServers"`
	// Any tags assigned to the DHCP options set.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

