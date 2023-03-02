package awsec2


// Describes the options for instance hostnames.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateDnsNameOptionsProperty := &privateDnsNameOptionsProperty{
//   	enableResourceNameDnsAaaaRecord: jsii.Boolean(false),
//   	enableResourceNameDnsARecord: jsii.Boolean(false),
//   	hostnameType: jsii.String("hostnameType"),
//   }
//
type CfnLaunchTemplate_PrivateDnsNameOptionsProperty struct {
	// Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records.
	EnableResourceNameDnsAaaaRecord interface{} `field:"optional" json:"enableResourceNameDnsAaaaRecord" yaml:"enableResourceNameDnsAaaaRecord"`
	// Indicates whether to respond to DNS queries for instance hostnames with DNS A records.
	EnableResourceNameDnsARecord interface{} `field:"optional" json:"enableResourceNameDnsARecord" yaml:"enableResourceNameDnsARecord"`
	// The type of hostname for EC2 instances.
	//
	// For IPv4 only subnets, an instance DNS name must be based on the instance IPv4 address. For IPv6 only subnets, an instance DNS name must be based on the instance ID. For dual-stack subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. For more information, see [Amazon EC2 instance hostname types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *Amazon Elastic Compute Cloud User Guide* .
	HostnameType *string `field:"optional" json:"hostnameType" yaml:"hostnameType"`
}

